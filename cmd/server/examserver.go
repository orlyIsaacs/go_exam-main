package main

import (
	"context"
	"errors"

	exam "github.com/priority-infra/go_exam/internal/protos"
)

type ExamServer struct {
	exam.UnimplementedExamServer
	employees   map[int64]*exam.Employee
	departments map[int64]*exam.Department
	projects    map[int64]*exam.Project
}

var ErrItemNotFound = errors.New("item does not exist")

func NewExamServer() *ExamServer {
	e := &ExamServer{}
	e.buildData()
	return e
}

func (e *ExamServer) GetEmployeeByID(_ context.Context, in *exam.ID) (*exam.Employee, error) {
	if emp, found := e.employees[in.GetValue()]; !found {
		return nil, ErrItemNotFound
	} else {
		return emp, nil
	}
}
func (e *ExamServer) GetEmployeeList(context.Context, *exam.Empty) (*exam.Employees, error) {
	res := &exam.Employees{}
	for _, emp := range e.employees {
		res.Employees = append(res.Employees, emp)
	}
	return res, nil
}
func (e *ExamServer) GetDepartmentByID(_ context.Context, in *exam.ID) (*exam.Department, error) {
	if dept, found := e.departments[in.GetValue()]; !found {
		return nil, ErrItemNotFound
	} else {
		return dept, nil
	}
}
func (e *ExamServer) GetDepartmentList(context.Context, *exam.Empty) (*exam.Departments, error) {
	res := &exam.Departments{}
	for _, depts := range e.departments {
		res.Departments = append(res.Departments, depts)
	}
	return res, nil
}
func (e *ExamServer) GetProjectByID(_ context.Context, in *exam.ID) (*exam.Project, error) {
	if proj, found := e.projects[in.GetValue()]; !found {
		return nil, ErrItemNotFound
	} else {
		return proj, nil
	}
}
func (e *ExamServer) GetProjectList(context.Context, *exam.Empty) (*exam.Projects, error) {
	res := &exam.Projects{}
	for _, proj := range e.projects {
		res.Projects = append(res.Projects, proj)
	}
	return res, nil
}
