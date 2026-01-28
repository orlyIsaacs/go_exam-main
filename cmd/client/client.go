package main

import (
	"context"
	"fmt"
	"log"
	"sort"
	"time"

	exam "github.com/priority-infra/go_exam/internal/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type row struct {
	ManagerName  string
	ProjectCount int
	Department   string
}

func main() {
	// 1) Connect to the gRPC server
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can't dial server: %v", err)
	}
	defer conn.Close()

	c := exam.NewExamClient(conn)

	// 2) Set a timeout so the client won't hang forever
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 3) Fetch all the data we need
	employeesResp, err := c.GetEmployeeList(ctx, &exam.Empty{})
	if err != nil {
		log.Fatalf("can't get employees: %v", err)
	}
	departmentsResp, err := c.GetDepartmentList(ctx, &exam.Empty{})
	if err != nil {
		log.Fatalf("can't get departments: %v", err)
	}
	projectsResp, err := c.GetProjectList(ctx, &exam.Empty{})
	if err != nil {
		log.Fatalf("can't get projects: %v", err)
	}

	rows := buildRows(employeesResp.GetEmployees(), departmentsResp.GetDepartments(), projectsResp.GetProjects())
	printRows(rows)
}

// buildRows contains the core logic of the assignment.
// It is pure (no side effects), which makes it easy to unit test.
func buildRows(employees []*exam.Employee, departments []*exam.Department, projects []*exam.Project) []row {
	// employeeId -> employee
	employeeByID := make(map[int64]*exam.Employee, len(employees))
	for _, e := range employees {
		employeeByID[e.GetID()] = e
	}

	// departmentId -> number of projects
	projectCountByDept := make(map[int64]int)
	for _, p := range projects {
		projectCountByDept[p.GetDepartmentID()]++
	}

	rows := make([]row, 0, len(departments))
	for _, d := range departments {
		count := projectCountByDept[d.GetID()]
		if count <= 1 {
			continue
		}

		manager, ok := employeeByID[d.GetManagerID()]
		if !ok {
			// defensive: if data is inconsistent, skip this department
			continue
		}

		rows = append(rows, row{
			ManagerName:  manager.GetName(),
			ProjectCount: count,
			Department:   d.GetName(),
		})
	}

	// Sort by project count descending (tie-break by name)
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].ProjectCount == rows[j].ProjectCount {
			return rows[i].ManagerName < rows[j].ManagerName
		}
		return rows[i].ProjectCount > rows[j].ProjectCount
	})

	return rows
}

func printRows(rows []row) {
	if len(rows) == 0 {
		fmt.Println("No managers found with more than 1 project.")
		return
	}

	fmt.Println("Manager\tProjects\tDepartment")
	for _, r := range rows {
		fmt.Printf("%s\t%d\t%s\n", r.ManagerName, r.ProjectCount, r.Department)
	}
}
