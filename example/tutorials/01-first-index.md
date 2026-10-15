# Tutorial 1 — Your First Index

Goal: index a folder of markdown notes and query the result.

1. Create a folder with a few `.md` files.
2. Run:

   ```bash
   md-indexer --source ./my-notes --out index.db
   ```

3. Query it:

   ```bash
   sqlite3 index.db "SELECT path, title FROM docs LIMIT 10;"
   ```

Next: [Batch processing](02-batch-processing/README.md).
