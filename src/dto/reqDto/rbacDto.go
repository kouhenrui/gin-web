package reqDto

type UpdateRbac struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
type AddRbac struct {
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
}
type RbacList struct {
	Take int    `json:"take,omitempty" binding:"required"`
	Skip int    `json:"skip,omitempty" binding:""`
	Name string `json:"name,omitempty"`
}

type PermissionList struct {
	Take int    `json:"take" binding:"required"`
	Skip int    `json:"skip"binding:"" `
	Path string `json:"path,omitempty"`
}
type RuleList struct {
	Take int    `json:"take" binding:"required"`
	Skip int    `json:"skip"binding:"" `
	Name string `json:"name,omitempty"`
}
type GroupList struct {
	Take int    `json:"take" binding:"required"`
	Skip int    `json:"skip"binding:"" `
	Name string `json:"name,omitempty"`
}
type PermissionAdd struct {
	Host            string `json:"host"`
	Path            string `json:"path" binding:"required"`
	Method          string `json:"method" binding:"required"`
	AuthorizedRoles string `json:"authorized_roles" binding:"required"`
	ForbiddenRoles  string `json:"forbidden_roles" binding:"required"`
	AllowAnyone     bool   `json:"allow_anyone"`
}

type RuleAdd struct {
	Name string `json:"name"`
}
type GroupAdd struct {
	Name         string `json:"name"`
	RoleId       string `json:"role_id"`
	PermissionId string `json:"permission_id"`
}

type PermissionUpdate struct {
	ID              uint   `json:"id" binding:"required"`
	Host            string `json:"host" sql:"host"`
	Path            string `json:"path"`
	Method          string `json:"method"`
	AuthorizedRoles string `json:"authorized_roles"`
	ForbiddenRoles  string `json:"forbidden_roles"`
	AllowAnyone     bool   `json:"allow_anyone"`
}
type GroupUpdate struct {
	ID           uint   `json:"id" binding:"required"`
	Name         string `json:"name"`
	RoleId       string `json:"role_id"`
	PermissionId string `json:"permission_id"`
}

type RuleUpdate struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
