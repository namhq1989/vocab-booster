package domain

type Action string

const (
	ActionUnknown Action = ""
	ActionCreate  Action = "create"
	ActionRead    Action = "read"
	ActionUpdate  Action = "update"
	ActionDelete  Action = "delete"
)

func (r Action) String() string {
	switch r {
	case ActionCreate, ActionRead, ActionUpdate, ActionDelete:
		return string(r)
	default:
		return ""
	}
}

func (r Action) IsValid() bool {
	return r != ActionUnknown
}

func ToAction(value string) Action {
	switch value {
	case ActionCreate.String():
		return ActionCreate
	case ActionRead.String():
		return ActionRead
	case ActionUpdate.String():
		return ActionUpdate
	case ActionDelete.String():
		return ActionDelete
	default:
		return ActionUnknown
	}
}
