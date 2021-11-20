# fts

`fts` is a full-text search package in pure Go.

## Design (WIP)

While the goal is to build a search engine this project isn't designed for production uses.

Currently the following features are supported :

- Parsing XML Documents to a Document collection
- Index building at runtime
- Tokenization, Normalization and Stemming (Using Porter's Algorithm)

In the future the following features will be added :

- Support JSON and Text corpus.
- BoltDB for on-disk storage.
- Support index building concurrently and index compression.
- Support skip-pointer construction for posting lists.
- Support more tokenizers.
- Support boolean queries.
- Use Roaring Bitmaps for posting lists.
- Support ranking and term weighting (TFIDF).

## Datasets

While implementing and testing use the [enwiki dataset](https://dumps.wikimedia.org/enwiki/latest/).
