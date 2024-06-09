package domain

type StaffStatus string

const (
	StaffStatusUnknown  StaffStatus = ""
	StaffStatusActive   StaffStatus = "active"
	StaffStatusInactive StaffStatus = "inactive"
)

func (s StaffStatus) String() string {
	switch s {
	case StaffStatusActive, StaffStatusInactive:
		return string(s)
	default:
		return ""
	}
}

func (s StaffStatus) IsValid() bool {
	return s != StaffStatusUnknown
}

func ToStaffStatus(value string) StaffStatus {
	switch value {
	case StaffStatusActive.String():
		return StaffStatusActive
	case StaffStatusInactive.String():
		return StaffStatusInactive
	default:
		return StaffStatusUnknown
	}
}
