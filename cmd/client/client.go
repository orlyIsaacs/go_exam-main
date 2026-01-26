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
	ManagerID    int64
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

	// Timeout so the client won't hang forever
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 2) Fetch all the data we need
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

	employees := employeesResp.GetEmployees()
	departments := departmentsResp.GetDepartments()
	projects := projectsResp.GetProjects()

	// 3) Build lookup: employeeId -> employee
	empByID := make(map[int64]*exam.Employee, len(employees))
	for _, e := range employees {
		empByID[e.GetID()] = e
	}

	// 4) Count projects per department
	projectsByDept := make(map[int64]int)
	for _, p := range projects {
		projectsByDept[p.GetDepartmentID()]++
	}

	// 5) Build rows: manager + number of projects in their department
	var rows []row
	for _, d := range departments {
		count := projectsByDept[d.GetID()]
		if count <= 1 {
			continue // only managers with > 1 project
		}

		mgr, ok := empByID[d.GetManagerID()]
		if !ok {
			continue // defensive: skip if manager missing
		}

		rows = append(rows, row{
			ManagerName:  mgr.GetName(),
			ManagerID:    mgr.GetID(),
			ProjectCount: count,
			Department:   d.GetName(),
		})
	}

	// 6) Sort by project count descending (tie-break by name)
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].ProjectCount == rows[j].ProjectCount {
			return rows[i].ManagerName < rows[j].ManagerName
		}
		return rows[i].ProjectCount > rows[j].ProjectCount
	})

	// 7) Print
	if len(rows) == 0 {
		fmt.Println("No managers found with more than 1 project.")
		return
	}

	fmt.Println("Manager\tProjects\tDepartment")
	for _, r := range rows {
		fmt.Printf("%s\t%d\t%s\n", r.ManagerName, r.ProjectCount, r.Department)
	}
}
