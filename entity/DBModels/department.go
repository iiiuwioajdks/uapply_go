package DBModels

// DepartmentInfo 部门信息
type DepartmentInfo struct {
	DepartmentID   int64  `json:"department_id" db:"department_id"`
	OrganizationID int64  `json:"organization_id" db:"organization_id"`
	DepartmentName string `json:"department_name" db:"department_name"`
}
