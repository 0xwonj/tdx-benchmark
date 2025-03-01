package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/0xwonj/tdx-benchmark/attest"
	benchmark "github.com/0xwonj/tdx-benchmark/benchmark"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

// BenchmarkType represents different types of benchmarks
type BenchmarkType int

const (
	NetworkBenchmark BenchmarkType = iota
	CPUBenchmark
	MemoryBenchmark
	DiskIOBenchmark
	AttestationBenchmark
)

// BenchmarkConfig represents the YAML configuration for all benchmarks
type BenchmarkConfig struct {
	Network struct {
		Enabled  bool `yaml:"enabled"`
		Configurations []struct {
			Clients           int `yaml:"clients"`
			RequestsPerClient int `yaml:"requests_per_client"`
		} `yaml:"configurations"`
	} `yaml:"network"`
	
	CPU struct {
		Enabled    bool    `yaml:"enabled"`
		Iterations []int64 `yaml:"iterations"`
	} `yaml:"cpu"`
	
	Memory struct {
		Enabled bool    `yaml:"enabled"`
		SizesMB []int32 `yaml:"sizes_mb"`
	} `yaml:"memory"`
	
	DiskIO struct {
		Enabled  bool `yaml:"enabled"`
		Configurations []struct {
			FileSizeMB int32 `yaml:"file_size_mb"`
			NumFiles   int32 `yaml:"num_files"`
		} `yaml:"configurations"`
	} `yaml:"disk_io"`
	
	Attestation struct {
		Enabled  bool `yaml:"enabled"`
		Configurations []struct {
			Clients           int `yaml:"clients"`
			RequestsPerClient int `yaml:"requests_per_client"`
		} `yaml:"configurations"`
	} `yaml:"attestation"`
	
	OutputDir string `yaml:"output_dir"`
}

// BenchmarkResult represents a benchmark result
type BenchmarkResult struct {
	Name       string
	Parameters map[string]any
	Metrics    map[string]any
	StartTime  time.Time
	Duration   time.Duration
}

// Logger wraps multiple log.Logger instances for different outputs
type Logger struct {
	Console *log.Logger
	File    *log.Logger
}

// LogResult logs a benchmark result to both console and file
func (l *Logger) LogResult(result BenchmarkResult) {
	for _, logger := range []*log.Logger{l.Console, l.File} {
		logger.Printf("\n%s Results:", result.Name)
		logger.Printf("-----------------------------------")
		
		// Log parameters
		for key, value := range result.Parameters {
			logger.Printf("%s: %v", key, value)
		}
		
		// Log metrics
		for key, value := range result.Metrics {
			logger.Printf("%s: %v", key, value)
		}
		
		logger.Printf("Total Duration: %v", result.Duration)
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// run performs the main program execution
func run() error {
	// Record total start time
	totalStartTime := time.Now()
	
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("loading .env file: %w", err)
	}

	// Check command line args
	if len(os.Args) < 2 {
		return errors.New("usage: client <config.yaml>")
	}
	
	// Load benchmark configuration from YAML
	configPath := os.Args[1]
	config, err := loadBenchmarkConfig(configPath)
	if err != nil {
		return fmt.Errorf("loading benchmark configuration: %w", err)
	}
	
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(config.OutputDir, 0755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}
	
	// Initialize client
	serverAddr := fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDRESS"), os.Getenv("PORT"))
	client, err := benchmark.NewClient(serverAddr)
	if err != nil {
		return fmt.Errorf("creating client: %w", err)
	}
	defer client.Close()
	
	// Setup logging
	logger, logFile, err := setupLogging(config.OutputDir)
	if err != nil {
		return fmt.Errorf("setting up logging: %w", err)
	}
	defer logFile.Close()
	
	// Run benchmarks and collect results
	results, err := runBenchmarks(client, config, logger)
	if err != nil {
		return fmt.Errorf("running benchmarks: %w", err)
	}
	
	// Calculate total execution time
	totalDuration := time.Since(totalStartTime)
	
	// Output summary information
	logger.Console.Printf("\nTotal Benchmark Execution Time: %v", totalDuration)
	logger.File.Printf("\nTotal Benchmark Execution Time: %v", totalDuration)
	
	// Create separate CSV files for each benchmark type
	timestamp := time.Now().Format("20060102_150405")
	if err := writeResultsToSeparateCSVs(results, config.OutputDir, timestamp); err != nil {
		return fmt.Errorf("writing results to CSV files: %w", err)
	}
	
	logger.Console.Printf("All benchmarks completed. Results saved to %s directory", config.OutputDir)
	return nil
}

// setupLogging creates and configures console and file loggers
func setupLogging(outputDir string) (*Logger, *os.File, error) {
	timestamp := time.Now().Format("20060102_150405")
	logFilePath := filepath.Join(outputDir, fmt.Sprintf("benchmark_%s.log", timestamp))
	logFile, err := os.Create(logFilePath)
	if err != nil {
		return nil, nil, fmt.Errorf("creating log file: %w", err)
	}
	
	logger := &Logger{
		Console: log.New(os.Stdout, "", log.LstdFlags),
		File:    log.New(logFile, "", log.LstdFlags),
	}
	
	return logger, logFile, nil
}

// loadBenchmarkConfig loads the benchmark configuration from a YAML file
func loadBenchmarkConfig(path string) (*BenchmarkConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}
	
	var config BenchmarkConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("parsing YAML: %w", err)
	}
	
	return &config, nil
}

// runBenchmarks executes all configured benchmarks
func runBenchmarks(client *benchmark.Client, config *BenchmarkConfig, logger *Logger) ([]BenchmarkResult, error) {
	var results []BenchmarkResult
	
	// Run network benchmarks (both sequential and concurrent)
	if config.Network.Enabled {
		for _, config := range config.Network.Configurations {
			result := runNetworkBenchmark(client, config.Clients, config.RequestsPerClient)
			results = append(results, result)
			logger.LogResult(result)
		}
	}
	
	// Run CPU benchmarks
	if config.CPU.Enabled {
		for _, iterations := range config.CPU.Iterations {
			result := runCPUBenchmark(client, iterations)
			results = append(results, result)
			logger.LogResult(result)
		}
	}
	
	// Run memory benchmarks
	if config.Memory.Enabled {
		for _, sizeMB := range config.Memory.SizesMB {
			result := runMemoryBenchmark(client, sizeMB)
			results = append(results, result)
			logger.LogResult(result)
		}
	}
	
	// Run disk I/O benchmarks
	if config.DiskIO.Enabled {
		for _, config := range config.DiskIO.Configurations {
			result := runDiskIOBenchmark(client, config.FileSizeMB, config.NumFiles)
			results = append(results, result)
			logger.LogResult(result)
		}
	}
	
	// Run attestation benchmarks
	if config.Attestation.Enabled {
		// Initialize attestation client
		serverAddr := fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDRESS"), os.Getenv("PORT"))
		attestClient, err := attest.NewClient(serverAddr)
		if err != nil {
			return results, fmt.Errorf("creating attestation client: %w", err)
		}
		defer attestClient.Close()
		
		for _, config := range config.Attestation.Configurations {
			result := runAttestationBenchmark(attestClient, config.Clients, config.RequestsPerClient)
			results = append(results, result)
			logger.LogResult(result)
		}
	}
	
	return results, nil
}

// writeResultsToSeparateCSVs writes benchmark results to separate CSV files by benchmark type
func writeResultsToSeparateCSVs(results []BenchmarkResult, outputDir, timestamp string) error {
	// Group results by benchmark type
	benchmarkGroups := groupResultsByBenchmarkType(results)
	
	// Create a CSV file for each benchmark type
	for benchmarkType, groupResults := range benchmarkGroups {
		if err := writeBenchmarkCSV(benchmarkType, groupResults, outputDir, timestamp); err != nil {
			return fmt.Errorf("writing CSV for %s: %w", benchmarkType, err)
		}
	}
	
	return nil
}

// groupResultsByBenchmarkType groups benchmark results by their type
func groupResultsByBenchmarkType(results []BenchmarkResult) map[string][]BenchmarkResult {
	groups := make(map[string][]BenchmarkResult)
	
	for _, result := range results {
		groups[result.Name] = append(groups[result.Name], result)
	}
	
	return groups
}

// writeBenchmarkCSV writes results for a single benchmark type to a CSV file with a simplified format
func writeBenchmarkCSV(benchmarkType string, results []BenchmarkResult, outputDir, timestamp string) error {
	// Create sanitized filename from benchmark type
	sanitizedType := strings.ToLower(benchmarkType)
	sanitizedType = strings.ReplaceAll(sanitizedType, " ", "_")
	sanitizedType = strings.ReplaceAll(sanitizedType, "/", "_")
	sanitizedType = strings.ReplaceAll(sanitizedType, "\\", "_")
	filename := fmt.Sprintf("%s_%s.csv", sanitizedType, timestamp)
	filepath := filepath.Join(outputDir, filename)
	
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("creating file %s: %w", filepath, err)
	}
	defer file.Close()
	
	// Extract parameter names from the first result
	var paramNames []string
	if len(results) > 0 {
		for paramName := range results[0].Parameters {
			paramNames = append(paramNames, paramName)
		}
		sort.Strings(paramNames) // Sort parameter names for consistent order
	}
	
	// Determine the throughput metric name based on benchmark type
	var throughputMetric string
	switch {
	case strings.Contains(benchmarkType, "Network"):
		throughputMetric = "Throughput (req/s)"
	case strings.Contains(benchmarkType, "CPU"):
		throughputMetric = "Iterations Per Second"
	case strings.Contains(benchmarkType, "Memory"):
		throughputMetric = "Average Time Per Page (ns)" // Not exactly throughput, but a performance metric
	case strings.Contains(benchmarkType, "Disk"):
		throughputMetric = "Throughput (MB/s)"
	case strings.Contains(benchmarkType, "Attestation"):
		throughputMetric = "Throughput (quotes/s)"
	default:
		throughputMetric = "Throughput"
	}
	
	// Write CSV header
	header := append(paramNames, "Duration (seconds)", throughputMetric)
	if _, err := file.WriteString(strings.Join(header, ",") + "\n"); err != nil {
		return fmt.Errorf("writing header: %w", err)
	}
	
	// Write each benchmark result as a row
	for _, result := range results {
		var row []string
		
		// Add parameter values
		for _, paramName := range paramNames {
			if value, exists := result.Parameters[paramName]; exists {
				row = append(row, fmt.Sprintf("%v", value))
			} else {
				row = append(row, "")
			}
		}
		
		// Add duration in seconds
		row = append(row, fmt.Sprintf("%.6f", result.Duration.Seconds()))
		
		// Add throughput value
		if throughputValue, exists := result.Metrics[throughputMetric]; exists {
			row = append(row, fmt.Sprintf("%v", throughputValue))
		} else {
			row = append(row, "N/A")
		}
		
		// Write the row
		if _, err := file.WriteString(strings.Join(row, ",") + "\n"); err != nil {
			return fmt.Errorf("writing result row: %w", err)
		}
	}
	
	return nil
}

// runNetworkBenchmark tests network performance with both sequential and concurrent requests
func runNetworkBenchmark(client *benchmark.Client, numClients, requestsPerClient int) BenchmarkResult {
	totalRequests := numClients * requestsPerClient
	
	benchmarkType := "Sequential"
	if numClients > 1 {
		benchmarkType = "Concurrent"
	}
	
	result := BenchmarkResult{
		Name: fmt.Sprintf("Network Benchmark (%s)", benchmarkType),
		Parameters: map[string]any{
			"Clients":           numClients,
			"RequestsPerClient": requestsPerClient,
			"TotalRequests":     totalRequests,
		},
		Metrics:   make(map[string]any),
		StartTime: time.Now(),
	}
	
	// For sequential case (single client), we can use a simpler approach
	if numClients == 1 {
		start := time.Now()
		errorCount := 0
		
		for i := 0; i < requestsPerClient; i++ {
			if _, err := client.Hello(); err != nil {
				errorCount++
				log.Printf("Request %d failed: %v", i, err)
			}
		}
		
		result.Duration = time.Since(start)
		successfulRequests := requestsPerClient - errorCount
		
		// Record metrics
		result.Metrics["Successful Requests"] = successfulRequests
		result.Metrics["Success Rate"] = float64(successfulRequests) * 100 / float64(requestsPerClient)
		result.Metrics["Failed Requests"] = errorCount
		result.Metrics["Throughput (req/s)"] = float64(successfulRequests) / result.Duration.Seconds()
		
		return result
	}
	
	// For concurrent case (multiple clients)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	
	start := time.Now()
	clientResults, errorCount := executeNetworkConcurrentRequests(ctx, client, numClients, requestsPerClient)
	result.Duration = time.Since(start)
	
	// Calculate successful requests
	successfulRequests := totalRequests - errorCount
	
	// Record metrics
	recordNetworkMetrics(&result, clientResults, successfulRequests, errorCount, totalRequests)
	
	return result
}

// executeNetworkConcurrentRequests executes concurrent requests to the server
func executeNetworkConcurrentRequests(
	ctx context.Context, 
	client *benchmark.Client, 
	numClients, 
	requestsPerClient int,
) ([]time.Duration, int) {
	// Channel to collect results from each goroutine
	resultCh := make(chan time.Duration, numClients)
	errorCh := make(chan error, numClients*requestsPerClient) // Channel to collect errors
	
	// Create a wait group to wait for all goroutines to complete
	var wg sync.WaitGroup
	wg.Add(numClients)
	
	// Launch client goroutines
	for i := range numClients {
		go func(clientID int) {
			defer wg.Done()
			
			clientStart := time.Now()
			for j := 0; j < requestsPerClient; j++ {
				select {
				case <-ctx.Done():
					errorCh <- ctx.Err()
					return
				default:
					if _, err := client.Hello(); err != nil {
						errorCh <- fmt.Errorf("client %d: request %d: %w", clientID, j, err)
						continue
					}
				}
			}
			clientDuration := time.Since(clientStart)
			resultCh <- clientDuration
		}(i)
	}
	
	// Close the result channel once all goroutines are done
	go func() {
		wg.Wait()
		close(resultCh)
		close(errorCh)
	}()
	
	// Collect results
	clientDurations := make([]time.Duration, 0, numClients)
	for duration := range resultCh {
		clientDurations = append(clientDurations, duration)
	}
	
	// Process any errors
	errorCount := 0
	for err := range errorCh {
		errorCount++
		log.Printf("Error: %v", err)
	}
	
	return clientDurations, errorCount
}

// recordNetworkMetrics calculates and records metrics for network benchmarks
func recordNetworkMetrics(
	result *BenchmarkResult, 
	clientDurations []time.Duration, 
	successfulRequests, 
	errorCount, 
	totalRequests int,
) {
	result.Metrics["Successful Requests"] = successfulRequests
	result.Metrics["Success Rate"] = float64(successfulRequests) * 100 / float64(totalRequests)
	result.Metrics["Failed Requests"] = errorCount
	
	if len(clientDurations) > 0 {
		// Calculate duration statistics
		var totalClientDuration time.Duration
		minDuration := time.Hour
		maxDuration := time.Duration(0)
		
		for _, duration := range clientDurations {
			totalClientDuration += duration
			if duration < minDuration {
				minDuration = duration
			}
			if duration > maxDuration {
				maxDuration = duration
			}
		}
		
		result.Metrics["Average Client Duration"] = totalClientDuration / time.Duration(len(clientDurations))
		result.Metrics["Fastest Client"] = minDuration
		result.Metrics["Slowest Client"] = maxDuration
	}
	
	result.Metrics["Throughput (req/s)"] = float64(successfulRequests) / result.Duration.Seconds()
}

// runCPUBenchmark tests CPU performance
func runCPUBenchmark(client *benchmark.Client, iterations int64) BenchmarkResult {
	result := BenchmarkResult{
		Name: "CPU Benchmark",
		Parameters: map[string]any{
			"Iterations": iterations,
		},
		Metrics:   make(map[string]any),
		StartTime: time.Now(),
	}
	
	// Execute benchmark
	start := time.Now()
	if _, err := client.CPUIntensive(iterations); err != nil {
		log.Printf("Failed to execute CPU benchmark: %v", err)
		result.Metrics["Error"] = err.Error()
		result.Duration = time.Since(start)
		return result
	}
	result.Duration = time.Since(start)
	
	// Record metrics
	result.Metrics["Iterations"] = iterations
	result.Metrics["Iterations Per Second"] = float64(iterations) / result.Duration.Seconds()
	
	return result
}

// runMemoryBenchmark tests memory performance
func runMemoryBenchmark(client *benchmark.Client, sizeMB int32) BenchmarkResult {
	result := BenchmarkResult{
		Name: "Memory Benchmark",
		Parameters: map[string]any{
			"Size (MB)": sizeMB,
		},
		Metrics:   make(map[string]any),
		StartTime: time.Now(),
	}
	
	// Execute benchmark
	start := time.Now()
	resp, err := client.MemoryIntensive(sizeMB)
	if err != nil {
		log.Printf("Failed to execute Memory benchmark: %v", err)
		result.Metrics["Error"] = err.Error()
		result.Duration = time.Since(start)
		return result
	}
	result.Duration = time.Since(start)
	
	// Record metrics
	result.Metrics["Pages Accessed"] = resp.PagesAccessed
	result.Metrics["Random Access Time (ns)"] = resp.AccessTimeNs
	
	if resp.PagesAccessed > 0 {
		result.Metrics["Average Time Per Page (ns)"] = resp.AccessTimeNs / int64(resp.PagesAccessed)
	}
	
	return result
}

// runDiskIOBenchmark tests disk I/O performance
func runDiskIOBenchmark(client *benchmark.Client, fileSizeMB, numFiles int32) BenchmarkResult {
	result := BenchmarkResult{
		Name: "Disk I/O Benchmark",
		Parameters: map[string]any{
			"File Size (MB)": fileSizeMB,
			"Number of Files": numFiles,
			"Total Size (MB)": fileSizeMB * numFiles,
		},
		Metrics:   make(map[string]any),
		StartTime: time.Now(),
	}
	
	// Execute benchmark
	start := time.Now()
	if _, err := client.DiskIO(fileSizeMB, numFiles); err != nil {
		log.Printf("Failed to execute I/O benchmark: %v", err)
		result.Metrics["Error"] = err.Error()
		result.Duration = time.Since(start)
		return result
	}
	result.Duration = time.Since(start)
	
	// Record metrics
	result.Metrics["Throughput (MB/s)"] = float64(fileSizeMB*numFiles) / result.Duration.Seconds()
	
	return result
}

// runAttestationBenchmark tests attestation (quote generation) performance with concurrent clients
func runAttestationBenchmark(client *attest.Client, numClients, requestsPerClient int) BenchmarkResult {
	totalRequests := numClients * requestsPerClient
	
	result := BenchmarkResult{
		Name: "Attestation Benchmark",
		Parameters: map[string]any{
			"Clients":           numClients,
			"RequestsPerClient": requestsPerClient,
			"TotalRequests":     totalRequests,
		},
		Metrics:   make(map[string]any),
		StartTime: time.Now(),
	}
	
	// Generate fixed report data (64 bytes is standard for TDX)
	reportData := make([]byte, 64)
	for i := range reportData {
		reportData[i] = byte(i % 256) // Simple pattern for reproducibility
	}
	
	// Execute benchmark with a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	
	start := time.Now()
	clientResults, errorCount := executeAttestationConcurrentRequests(ctx, client, numClients, requestsPerClient, reportData)
	result.Duration = time.Since(start)
	
	// Calculate successful requests
	successfulRequests := totalRequests - errorCount
	
	// Record metrics
	recordAttestationMetrics(&result, clientResults, successfulRequests, errorCount, totalRequests)
	
	return result
}

// executeAttestationConcurrentRequests executes concurrent quote requests to the attestation server
func executeAttestationConcurrentRequests(
	ctx context.Context, 
	client *attest.Client, 
	numClients, 
	requestsPerClient int,
	reportData []byte,
) ([]time.Duration, int) {
	// Channel to collect results from each goroutine
	resultCh := make(chan time.Duration, numClients)
	errorCh := make(chan error, numClients*requestsPerClient) // Channel to collect errors
	
	// Create a wait group to wait for all goroutines to complete
	var wg sync.WaitGroup
	wg.Add(numClients)
	
	// Launch client goroutines
	for i := range  numClients {
		go func(clientID int) {
			defer wg.Done()
			
			clientStart := time.Now()
			for j := range requestsPerClient {
				select {
				case <-ctx.Done():
					errorCh <- ctx.Err()
					return
				default:
					// Get a quote from the server
					requestCtx, requestCancel := context.WithTimeout(ctx, 30*time.Second)
					_, err := client.GetQuote(requestCtx, reportData)
					requestCancel()
					
					if err != nil {
						errorCh <- fmt.Errorf("client %d: request %d: %w", clientID, j, err)
						continue
					}
				}
			}
			clientDuration := time.Since(clientStart)
			resultCh <- clientDuration
		}(i)
	}
	
	// Close the result channel once all goroutines are done
	go func() {
		wg.Wait()
		close(resultCh)
		close(errorCh)
	}()
	
	// Collect results
	clientDurations := make([]time.Duration, 0, numClients)
	for duration := range resultCh {
		clientDurations = append(clientDurations, duration)
	}
	
	// Process any errors
	errorCount := 0
	for err := range errorCh {
		errorCount++
		log.Printf("Error: %v", err)
	}
	
	return clientDurations, errorCount
}

// recordAttestationMetrics calculates and records metrics for attestation benchmarks
func recordAttestationMetrics(
	result *BenchmarkResult, 
	clientDurations []time.Duration, 
	successfulRequests, 
	errorCount, 
	totalRequests int,
) {
	result.Metrics["Successful Requests"] = successfulRequests
	result.Metrics["Success Rate"] = float64(successfulRequests) * 100 / float64(totalRequests)
	result.Metrics["Failed Requests"] = errorCount
	
	if len(clientDurations) > 0 {
		// Calculate duration statistics
		var totalClientDuration time.Duration
		minDuration := time.Hour
		maxDuration := time.Duration(0)
		
		for _, duration := range clientDurations {
			totalClientDuration += duration
			if duration < minDuration {
				minDuration = duration
			}
			if duration > maxDuration {
				maxDuration = duration
			}
		}
		
		result.Metrics["Average Client Duration"] = totalClientDuration / time.Duration(len(clientDurations))
		result.Metrics["Fastest Client"] = minDuration
		result.Metrics["Slowest Client"] = maxDuration
		
		if successfulRequests > 0 {
			result.Metrics["Average Quote Generation Time"] = totalClientDuration / time.Duration(successfulRequests)
		}
	}
	
	result.Metrics["Throughput (quotes/s)"] = float64(successfulRequests) / result.Duration.Seconds()
}