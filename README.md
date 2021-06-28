# kesl-benchmark
Benchmark tool written on GoLang for Benchmarkin Kaspersky Endpoint Security for Linux installation

With this tool you can simply check what number of files can check KESL in seconds

Run this command to get help
```bash
go run main.go -h 
```

Help command output
```bash
root@kesl#: go run main.go -h

-fileSize int
        File size in bytes that need to be generate for KESL scan-file task (default 1000000)
-scansCount int
        Number of tasks of scanning files in KESL (each task generate file for itself be ware with fileSize*scansCount formula) (default 10)
-threads int
        Number of gorutines that need to be running concurrently (default 4)
```
