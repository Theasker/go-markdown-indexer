# Troubleshooting

## Build fails with `go: modules disabled`

Run `go mod tidy` and retry.

## Index is empty

Check that the source path contains `.md` files and is readable:

```bash
find ./my-notes -name '*.md' | head
```

## Slow indexing

See [Advanced: Performance](advanced/performance.md).
