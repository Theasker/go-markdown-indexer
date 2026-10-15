# Internals: Deep Dive Notes

> Working notes — this is the deepest doc level in the tree, used to test
> nested relative links: [back to performance](../performance.md) and
> [back to docs root](../../README.md).

## Pipeline

1. Walk source tree (`filepath.WalkDir`).
2. Parse markdown headings and links.
3. Upsert rows into SQLite (FTS5).
