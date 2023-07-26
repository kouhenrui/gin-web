package rbacService

import (
	"HelloGin/src/dto/reqDto"
	"HelloGin/src/dto/resDto"
	"HelloGin/src/pojo"
	"fmt"
)

/**
 * @ClassName rbacService
 * @Description 权限表操作层
 * @Author khr
 * @Date 2023/4/3 16:05
 * @Version 1.0
 */
var (
	ruleServiceImpl       = pojo.RbacRule()
	groupServiceImpl      = pojo.RbacGroup()
	permissionServiceImpl = pojo.RbacPermission()
	permission            = &pojo.Permission{}
	permissionInfo        = &resDto.PermissionInfo{}
	resList               = &resDto.CommonList{}
	err                   error
)

/*
 * @MethodName FindPermissionByPath
 * @Description  根据路径模糊查询权限
 * @Author Acer
 * @Date 2023/4/4 9:14
 */

func FindPermissionByPath(path string) (error, *resDto.PermissionInfo) {
	err, permissionInfo = permissionServiceImpl.FindPermissionByPath(path)
	if err != nil {
		return err, nil
	}
	return nil, permissionInfo
}

/*
 * @MethodName FindPermissionById
 * @Description 根据id精准查询权限
 * @Author Acer
 * @Date 2023/4/4 9:27
 */

func FindPermissionById(id int) (error, *resDto.PermissionInfo) {
	err, permissionInfo = permissionServiceImpl.FindPermissionById(id)
	if err != nil {
		return err, nil
	}
	return nil, permissionInfo
}

/*
 * @MethodName AddPermission
 * @Description 增加权限
 * @Author Acer
 * @Date 2023/4/4 9:28
 */

func AddPermission(permission pojo.Permission) error {
	return permissionServiceImpl.AddPermission(permission)
}

/*
 * @MethodName FindPermissionList
 * @Description 权限列表 可模糊查询路径
 * @Author Acer
 * @Date 2023/4/4 9:28
 */

func FindPermissionList(list reqDto.PermissionList) (error, *resDto.CommonList) {
	err, resList = permissionServiceImpl.FindPermissionList(list)
	if err != nil {
		return err, nil
	}
	return nil, resList
}

/*
 * @MethodName UpdatePermission
 * @Description  更新权限 根据id 选择性更新 host path method AuthorizedRoles ForbiddenRoles AllowAnyone 字段
 * @Author Acer
 * @Date 2023/4/4 9:28
 */

func UpdatePermission(permission reqDto.PermissionUpdate) error {
	per := pojo.Permission{
		Host:            permission.Host,
		Path:            permission.Path,
		AuthorizedRoles: permission.AuthorizedRoles,
		ForbiddenRoles:  permission.ForbiddenRoles,
		Method:          permission.Method,
		AllowAnyone:     permission.AllowAnyone,
	}
	fmt.Println(":到达转换点")
	//util.DtoToStruct(permission, pojo.Permission{})
	per.ID = permission.ID
	return permissionServiceImpl.SavePermission(per)
}

/*
 * @MethodName DelPermmmission
 * @Description 根据id删除权限
 * @Author Acer
 * @Date 2023/4/4 9:32
 */

func DelPermission(id int) error {
	return permissionServiceImpl.DelPermission(id)
}

/*
 * @MethodName AddRole
 * @Description 增加角色
 * @Author
 * @Date 2023/4/4 9:35
 */

func AddRule(add reqDto.RuleAdd) error {
	rule := &pojo.Rule{
		Name: add.Name,
	}
	return ruleServiceImpl.AddRole(rule)
}

/*
 * @MethodName FindRuleById
 * @Description  根据id查询rule
 * @Author khr
 * @Date 2023/4/4 9:44
 */

func FindRuleById(id int) (error, *pojo.Rule) {
	var rule = &pojo.Rule{}
	err, rule = ruleServiceImpl.FindRoleName(uint(id))
	if err != nil {
		return err, nil
	}
	return nil, rule
}

/*
 * @MethodName UpdateRule
 * @Description 根据id修改角色
 * @Author khr
 * @Date 2023/4/4 9:48
 */

func UpdateRule(updateRule reqDto.RuleUpdate) error {
	rule := &pojo.Rule{
		Name: updateRule.Name,
	}
	rule.ID = updateRule.Id
	return ruleServiceImpl.UpdateRole(rule)
}

/*
 * @MethodName DelRule
 * @Description 根据id删除角色
 * @Author khr
 * @Date 2023/4/4 9:51
 */

func DelRule(id int) error {
	return ruleServiceImpl.DelRole(id)
}

/*
 * @MethodName FindRuleList
 * @Description 查询角色列表 可模糊查询角色名称
 * @Author khr
 * @Date 2023/4/4 9:54
 */

func FindRuleList(list reqDto.RuleList) (error, *resDto.CommonList) {
	err, resList = ruleServiceImpl.FindRoleList(list)
	if err != nil {
		return err, nil
	}
	return nil, resList
}

/*
 * @MethodName AddGroup
 * @Description 增加组
 * @Author khr
 * @Date 2023/4/4 10:42
 */

func AddGroup(addGroup reqDto.GroupAdd) error {
	group := &pojo.Group{
		Name:         addGroup.Name,
		RoleId:       addGroup.RoleId,
		PermissionId: addGroup.PermissionId,
	}
	err = groupServiceImpl.AddGroup(group)
	return err
}

/*
 * @MethodName UpdateGroup
 * @Description 修改组
 * @Author khr
 * @Date 2023/4/4 10:50
 */

func UpdateGroup(updateGroup reqDto.GroupUpdate) error {
	group := &pojo.Group{
		Name:         updateGroup.Name,
		RoleId:       updateGroup.RoleId,
		PermissionId: updateGroup.PermissionId,
	}
	group.ID = updateGroup.ID
	return groupServiceImpl.SaveGroup(group)

}

/*
 * @MethodName DelGroup
 * @Description 删除group
 * @Author khr
 * @Date 2023/4/4 10:56
 */

func DelGroup(id int) error {
	return groupServiceImpl.DelGroup(id)
}

/*
 * @MethodName FindGroupList
 * @Description
 * @Author khr
 * @Date 2023/4/4 13:40
 */

func FindGroupList(list reqDto.GroupList) (error, *resDto.CommonList) {
	err, resList = groupServiceImpl.FindGroupList(&list)
	if err != nil {
		return err, nil
	}
	return nil, resList
}

/*
 * @MethodName FindGroupById
 * @Description
 * @Author khr
 * @Date 2023/4/4 13:48
 */

func FindGroupById(id int) (error, *resDto.GroupInfo) {
	var groupInfo = &resDto.GroupInfo{}
	err, groupInfo = groupServiceImpl.FindGroupName(uint(id))
	if err != nil {
		return err, nil
	}
	return nil, groupInfo
}
