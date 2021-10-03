package DBModels

type Organizations struct {
	OrganizationID   int         `json:"organization_id" db:"organization_id"`
	OrganizationName string      `json:"organization_name" db:"organization_name"`
	Departments      []*DepOfOrg `json:"departments"`
}

// DepOfOrg 组织下的社团表
type DepOfOrg struct {
	DepartmentID   int    `json:"department_id" db:"department_id"`
	DepartmentName string `json:"department_name" db:"department_name"`
}

type Organization struct {
	OrganizationName string `json:"organization_name" db:"organization_name"`
}

type Department struct {
	DepartmentName string `json:"department_name" db:"department_name"`
	Account        string `json:"account" db:"account"`
	Password       string `json:"password" db:"password"`
	OrganizationID int    `json:"organization_id" db:"organization_id"`
}
