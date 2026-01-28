
## Overview
The client gets employees, departments, and projects from the gRPC server.
It counts how many projects each department has and connects each department to its manager.

Only managers with more than one project are printed.

## Approach
All the data is fetched from the server using the existing gRPC endpoints.
Maps are used to connect employees, departments, and projects and to avoid nested loops.

The main calculation is done in a separate function (`buildRows`), which makes the code easier to read and test.
The results are sorted by project count in descending order.

## Efficiency
Each list is processed once.
Map lookups are fast and keep the logic simple and scalable.

Sorting is done only on the final list of managers, which is small, so it does not have a big impact.

## Testing
Unit tests were added for the main calculation logic using Ginkgo and Gomega.
The tests check filtering, sorting, and edge cases without depending on the gRPC server.

Tests can be run from the project root using:
`go test ./...`

## Future Improvements
If I had more time, I would add tests for the output formatting.
For larger systems, I would also consider moving the aggregation logic to the server side.
