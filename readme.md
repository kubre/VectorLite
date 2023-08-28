# VectorLite: Simple Vector Database

Database for storing any querying the embeddings and metadata alongside it.

## WIP: Just toy project as of now not even in working state

**To Run:** `go run main.go record.go`

There is also python file to generate and test with embeddings as of now

## Development Log (26 Aug 2023)

**Requirements**

Indexing

-   Flat Index
    -   No modification to vectors store directly as series of floats
-   How to build and index for embeddings
    -   Random Projection
        -   Using matrix of MxM size to convert N size vector to M size
    -   HNSW (Hierarchical Navigable Small World)- Querying

Querying

-   Cosine similarity
    -   Formula: dot(v, w) / (v.norm() \* w.norm())
        -   dot v w: sum(vi \* wi)
        -   norm v: sqrt(sum(vi \* vi))
-   Euclidean Distance
-   Dot Product

Storing

-   How to store vector format
    -   Store test data into csv files from python with index and paragraph
