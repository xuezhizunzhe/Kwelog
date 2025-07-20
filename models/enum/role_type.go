package enum

type RoleType string

const (
	AdminRole  RoleType = "admin"
	NormalRole RoleType = "normal"
	VisitRole  RoleType = "visit"
)
