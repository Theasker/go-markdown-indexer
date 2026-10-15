# Walkthrough: Batch Processing

Index several repositories in one go using a file list:

```bash
md-indexer --list tutorials/assets/sample-input.csv --out bulk.db
```

The CSV format is: `name,path` per line (see [sample-input.csv](../assets/sample-input.csv)).
