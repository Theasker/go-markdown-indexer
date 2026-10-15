# Performance Tuning

- Use `--workers N` to parallelize file reads (default: number of CPUs).
- Exclude large directories with `--exclude`.
- The SQLite writer batches transactions; lower `--batch-size` reduces memory but hurts throughput.
