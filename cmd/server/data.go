package main

import exam "github.com/priority-infra/go_exam/internal/protos"

func (e *ExamServer) buildData() {
	e.buildDepartments()
	e.buildEmployees()
	e.buildProjects()
}

func (e *ExamServer) buildEmployees() {
	e.employees = map[int64]*exam.Employee{
		1: {
			ID:           1,
			DepartmentID: 1,
			Name:         "Andy Serkis",
		},
		2: {
			ID:           2,
			DepartmentID: 1,
			Name:         "Elijah Wood",
		},
		3: {
			ID:           3,
			DepartmentID: 1,
			Name:         "Ian McKellen",
		},
		4: {
			ID:           4,
			DepartmentID: 1,
			Name:         "Sean Astin",
		},
		5: {
			ID:           5,
			DepartmentID: 2,
			Name:         "Billy Boyd",
		},
		6: {
			ID:           6,
			DepartmentID: 2,
			Name:         "Dominic Monaghan",
		},
		7: {
			ID:           7,
			DepartmentID: 4,
			Name:         "Viggo Mortensen",
		},
		8: {
			ID:           8,
			DepartmentID: 2,
			Name:         "Christopher Lee",
		},
		9: {
			ID:           9,
			DepartmentID: 1,
			Name:         "Orlando Bloom",
		},
		10: {
			ID:           10,
			DepartmentID: 1,
			Name:         "John Rhys-Davies",
		},
		11: {
			ID:           11,
			DepartmentID: 3,
			Name:         "Sean Bean",
		},
		12: {
			ID:           12,
			DepartmentID: 3,
			Name:         "Bruce Spence",
		},
		13: {
			ID:           13,
			DepartmentID: 3,
			Name:         "John Leigh",
		},
		14: {
			ID:           14,
			DepartmentID: 3,
			Name:         "Sala Baker",
		},
		15: {
			ID:           15,
			DepartmentID: 3,
			Name:         "Liv Tyler",
		},
		16: {
			ID:           16,
			DepartmentID: 1,
			Name:         "Cate Blanchett",
		},
		17: {
			ID:           17,
			DepartmentID: 1,
			Name:         "Bernard Hill",
		},
		18: {
			ID:           18,
			DepartmentID: 4,
			Name:         "Miranda Otto",
		},
		19: {
			ID:           19,
			DepartmentID: 4,
			Name:         "Karl Urban",
		},
		20: {
			ID:           20,
			DepartmentID: 4,
			Name:         "David Wenham",
		},
		21: {
			ID:           21,
			DepartmentID: 4,
			Name:         "Brad Dourif",
		},
		22: {
			ID:           22,
			DepartmentID: 1,
			Name:         "Ian Holm",
		},
		23: {
			ID:           23,
			DepartmentID: 4,
			Name:         "John Noble",
		},
		24: {
			ID:           24,
			DepartmentID: 1,
			Name:         "Marton Csokas",
		},
		25: {
			ID:           25,
			DepartmentID: 4,
			Name:         "Craig Parker",
		},
		26: {
			ID:           26,
			DepartmentID: 1,
			Name:         "Lawrence Makoare",
		},
		27: {
			ID:           27,
			DepartmentID: 3,
			Name:         "Bruce Hopkins",
		},
		28: {
			ID:           28,
			DepartmentID: 2,
			Name:         "Hugo Weaving",
		},
		29: {
			ID:           29,
			DepartmentID: 1,
			Name:         "Sarah McLeod",
		},
		30: {
			ID:           30,
			DepartmentID: 1,
			Name:         "Harry Sinclair",
		},
	}
}

func (e *ExamServer) buildDepartments() {
	e.departments = map[int64]*exam.Department{
		1: {
			ID:        1,
			ManagerID: 1,
			Name:      "Marketing",
		},
		2: {
			ID:        2,
			ManagerID: 28,
			Name:      "R&D",
		},
		3: {
			ID:        3,
			ManagerID: 15,
			Name:      "Product",
		},
		4: {
			ID:        4,
			ManagerID: 7,
			Name:      "Support",
		},
	}
}

func (e *ExamServer) buildProjects() {
	e.projects = map[int64]*exam.Project{
		1: {
			ID:           1,
			DepartmentID: 1,
			Name:         "BrandAscend: Amplifying Market Presence",
		},
		2: {
			ID:           2,
			DepartmentID: 1,
			Name:         "StratEdge: Crafting Strategic Campaigns",
		},
		3: {
			ID:           3,
			DepartmentID: 1,
			Name:         "BuzzSprint: Accelerating Product Visibility",
		},
		4: {
			ID:           4,
			DepartmentID: 2,
			Name:         "VisionaryX: Breakthroughs in Emerging Technologies",
		},
		5: {
			ID:           5,
			DepartmentID: 2,
			Name:         "InnoLab Quest: Trailblazing R&D Endeavors",
		},
		6: {
			ID:           6,
			DepartmentID: 3,
			Name:         "BetaBoost: Optimizing Features for Peak Performance",
		},
		7: {
			ID:           7,
			DepartmentID: 3,
			Name:         "Project Nexus: Seamless Integration & Scalability",
		},
		8: {
			ID:           8,
			DepartmentID: 3,
			Name:         "CodeCraft: Revolutionizing User Experience",
		},
		9: {
			ID:           9,
			DepartmentID: 3,
			Name:         "InnoVision: Evolving User-Centric Solutions",
		},
		10: {
			ID:           10,
			DepartmentID: 4,
			Name:         "Operation Support Harmony: Enhancing Customer Experience & Service Excellence",
		},
		11: {
			ID:           11,
			DepartmentID: 2,
			Name:         "AlgoSphere: Pioneering Next-Gen Algorithms",
		},
		12: {
			ID:           12,
			DepartmentID: 2,
			Name:         "TechHorizon: Exploring Future Innovations",
		},
	}
}
