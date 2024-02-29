# LogFile Reader

Includes Functions for reading GigaBytes of log messages

## Approach

There are three approaches included:

## 1. Buffered Reading with bufio:

Uses the 'bufio' package in Golang for buffered scanning.
Reads the log file line by line, buffering the content for more efficient I/O operations.

## 2. Concurrent Processing:

This approach assumes the log lines are independent of each other. 
Uses goroutines to achieve parallelism and reads log-lines concurrently.
Requires a sync mechanism like wait groups to ensure all goroutines complete before proceeding.

## 3. Batch Processing:

Reads log entries in fixed-size batches.
Reduces the number of I/O operations by reading multiple lines at once.
Each batch of log entries is then processed collectively.

## Best Approach:

For gigabytes of log messages, a combination of concurrent processing and batch processing might be effective.

## Usage

To use this project, follow these steps:

1. Replace the "example.log" with your desired log file

2. Run the corresponding go file based on your approach, for example

   ```bash
   go run buffered.go

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.
