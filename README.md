# Welcome :-)

This project includes both client and server implementations for a simple microservice.

## System Requirements

Make sure you have Go installed with version 1.20 or later.

## Build and Execute Instructions for Server

Navigate to the root project directory and follow these steps:

```bash
cd cmd/server
go build -o server
./server
```

## Types of Entities

This microservice involves three types of entities: Employees, Departments, and Projects. Refer to the `exam.proto` file for detailed documentation on these entities.

## Mission: Definition of Done

Your mission is to write a program that accomplishes the following:

- Print all department managers and the number of projects each manager is responsible for (i.e., the projects his or her department is working on).
- Display only managers who are responsible for more than 1 project.
- Sort the list by the number of projects in descending order.

If you have any questions or need assistance, please don't hesitate to reach out.
