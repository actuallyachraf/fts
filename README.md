# fts

`fts` is a full-text search package in pure Go.

## Design (WIP)

While the goal is to build a search engine this project isn't designed for production uses.

- Support XML, JSON and Text corpus.
- BoltDB for storing compressed posting lists and inverted indexes.
- Support index building at runtime and index compression in the future.
- Support tokenization, stemming out of the box.
- Support byte-encoding pairs.

## Datasets

While implementing and testing use the [enwiki dataset](https://dumps.wikimedia.org/enwiki/latest/).
