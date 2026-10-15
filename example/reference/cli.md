# CLI Reference

```bash
md-indexer [flags]
```

| Flag          | Type   | Default       | Description                |
|---------------|--------|---------------|----------------------------|
| `--source`    | string | required      | Path to scan               |
| `--out`       | string | `index.db`    | SQLite output file         |
| `--workers`   | int    | NumCPU        | Parallel workers           |
| `--exclude`   | list   | —             | Directories to skip        |
| `--batch-size`| int    | `500`         | Rows per transaction       |
