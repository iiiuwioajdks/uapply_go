package DBModels

type Organizations struct {
	OrganizationID   int          `json:"organization_id" db:"organization_id"`
	OrganizationName string       `json:"organization_name" db:"organization_name"`
	Departments      []*DepOfOrg2 `json:"departments"`
}

type DepOfOrg2 struct {
	DepartmentID   int    `json:"department_id" db:"department_id"`
	DepartmentName string `json:"department_name" db:"department_name"`
}

// DepOfOrg 组织下的社团表
type DepOfOrg struct {
	OrganizationID int    `json:"organization_id" db:"organization_id"`
	DepartmentID   int    `json:"department_id" db:"department_id"`
	DepartmentName string `json:"department_name" db:"department_name"`
}

type Organization struct {
	OrganizationName string `json:"organization_name" db:"organization_name"`
}

// Department 社团注册模型
type Department struct {
	DepartmentName string `json:"department_name" db:"department_name" binding:"required"`
	Account        string `json:"account" db:"account" binding:"required"`
	Password       string `json:"password" db:"password" binding:"required"`
	OrganizationID int    `json:"organization_id" db:"organization_id" binding:"required"`
}
