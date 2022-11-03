package pluginutils

import (
	"strings"

	"github.com/grafana/grafana/pkg/plugins"
	ac "github.com/grafana/grafana/pkg/services/accesscontrol"
)

func ToRegistrations(pluginID, pluginName string, regs []plugins.RoleRegistration) []ac.RoleRegistration {
	res := make([]ac.RoleRegistration, 0, len(regs))
	for i := range regs {
		res = append(res, ac.RoleRegistration{
			Role: ac.RoleDTO{
				Version:     1,
				Name:        roleName(pluginID, regs[i].Role.Name),
				DisplayName: regs[i].Role.Name,
				Description: regs[i].Role.Description,
				Group:       pluginName,
				Permissions: toPermissions(regs[i].Role.Permissions),
				OrgID:       ac.GlobalOrgID,
			},
			Grants: regs[i].Grants,
		})
	}
	return res
}

func roleName(pluginID, roleName string) string {
	return ac.AppPluginRolePrefix + pluginID + ":" + strings.Replace(strings.ToLower(roleName), " ", "-", -1)
}

func toPermissions(perms []plugins.Permission) []ac.Permission {
	res := make([]ac.Permission, 0, len(perms))
	for i := range perms {
		res = append(res, ac.Permission{Action: perms[i].Action, Scope: perms[i].Scope})
	}
	return res
}
