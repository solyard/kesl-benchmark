# kesl-benchmark
Benchmark tool written in GoLang for Benchmarkin Kaspersky Endpoint Security for Linux installation

With this tool you can simply check what number of files can check KESL in seconds

Run this command to get help
```bash
go run main.go -h 
```
If u used a binary
```bash
./main_deb64 -h
```


Help command output
```bash
root@kesl#: go run main.go -h

-fileSize int
        File size in bytes that need to be generate for KESL scan-file task (default 1000000)
  -keslCommand string
        kesl-control command for file scanning (default "/opt/kaspersky/kesl/bin/kesl-control --scan-file %s --action Skip")
  -scansCount int
        Number of tasks of scanning files in KESL (default 10)
  -threads int
        Number of gorutines that need to be running (default 4)
```
