# Architecture

## Overview
The client retrieves employees, departments, and projects from the gRPC server.
It calculates how many projects each department is responsible for and associates each department with its manager.

## Approach
All required data is fetched using existing gRPC list endpoints.
Maps are used to join employees, departments, and projects efficiently, avoiding nested loops.

Only managers responsible for more than one project are included in the output.
The final result is sorted by project count in descending order.

## Efficiency
Each dataset (employees, departments, projects) is processed once, resulting in linear time complexity.
Map lookups are constant time and keep the aggregation fast and scalable.

Sorting is performed only on the final list of managers, which is small compared to the total data size, so its impact is minimal.

## Future Improvements
If I had more time, I would add unit tests for the aggregation logic and improve the output formatting.
