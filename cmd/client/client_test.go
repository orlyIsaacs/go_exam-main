package main

import (
	exam "github.com/priority-infra/go_exam/internal/protos"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("buildRows", func() {
	It("filters out managers with 1 or 0 projects", func() {
		employees := []*exam.Employee{
			{ID: 1, Name: "Manager A"},
			{ID: 2, Name: "Manager B"},
		}
		departments := []*exam.Department{
			{ID: 10, Name: "Dept A", ManagerID: 1},
			{ID: 20, Name: "Dept B", ManagerID: 2},
		}
		projects := []*exam.Project{
			{ID: 100, DepartmentID: 10, Name: "P1"},
			{ID: 101, DepartmentID: 10, Name: "P2"}, // Dept A has 2 projects
			{ID: 200, DepartmentID: 20, Name: "P3"}, // Dept B has 1 project
		}

		rows := buildRows(employees, departments, projects)

		Expect(rows).To(HaveLen(1))
		Expect(rows[0].ManagerName).To(Equal("Manager A"))
		Expect(rows[0].ProjectCount).To(Equal(2))
		Expect(rows[0].Department).To(Equal("Dept A"))
	})

	It("sorts by project count descending and tie-breaks by manager name", func() {
		employees := []*exam.Employee{
			{ID: 1, Name: "Bob"},
			{ID: 2, Name: "Alice"},
			{ID: 3, Name: "Charlie"},
		}
		departments := []*exam.Department{
			{ID: 10, Name: "Dept 10", ManagerID: 1}, // Bob
			{ID: 20, Name: "Dept 20", ManagerID: 2}, // Alice
			{ID: 30, Name: "Dept 30", ManagerID: 3}, // Charlie
		}
		projects := []*exam.Project{
			// Dept 10 -> 3 projects
			{ID: 100, DepartmentID: 10},
			{ID: 101, DepartmentID: 10},
			{ID: 102, DepartmentID: 10},

			// Dept 20 -> 3 projects (tie with Dept 10)
			{ID: 200, DepartmentID: 20},
			{ID: 201, DepartmentID: 20},
			{ID: 202, DepartmentID: 20},

			// Dept 30 -> 2 projects
			{ID: 300, DepartmentID: 30},
			{ID: 301, DepartmentID: 30},
		}

		rows := buildRows(employees, departments, projects)

		Expect(rows).To(HaveLen(3))
		// First two have 3 projects; tie-break by name => Alice then Bob
		Expect(rows[0].ManagerName).To(Equal("Alice"))
		Expect(rows[0].ProjectCount).To(Equal(3))

		Expect(rows[1].ManagerName).To(Equal("Bob"))
		Expect(rows[1].ProjectCount).To(Equal(3))

		Expect(rows[2].ManagerName).To(Equal("Charlie"))
		Expect(rows[2].ProjectCount).To(Equal(2))
	})

	It("skips departments if manager is missing from employees list", func() {
		employees := []*exam.Employee{
			{ID: 1, Name: "Existing Manager"},
		}
		departments := []*exam.Department{
			{ID: 10, Name: "Dept A", ManagerID: 1},
			{ID: 20, Name: "Dept Missing Manager", ManagerID: 999}, // Missing manager
		}
		projects := []*exam.Project{
			{ID: 100, DepartmentID: 10},
			{ID: 101, DepartmentID: 10},
			{ID: 200, DepartmentID: 20},
			{ID: 201, DepartmentID: 20},
		}

		rows := buildRows(employees, departments, projects)

		Expect(rows).To(HaveLen(1))
		Expect(rows[0].Department).To(Equal("Dept A"))
	})
})
