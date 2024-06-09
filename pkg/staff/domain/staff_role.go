package domain

type StaffRole string

const (
	StaffRoleUnknown StaffRole = ""
	StaffRoleAdmin   StaffRole = "admin"
	StaffRoleEditor  StaffRole = "editor"
)

func (s StaffRole) String() string {
	switch s {
	case StaffRoleAdmin, StaffRoleEditor:
		return string(s)
	default:
		return ""
	}
}

func (s StaffRole) IsValid() bool {
	return s != StaffRoleUnknown
}

func ToStaffRole(value string) StaffRole {
	switch value {
	case StaffRoleAdmin.String():
		return StaffRoleAdmin
	case StaffRoleEditor.String():
		return StaffRoleEditor
	default:
		return StaffRoleUnknown
	}
}
