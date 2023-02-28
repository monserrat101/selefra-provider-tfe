package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-tfe"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/terraform/selefra_terraform_schema"
)

// terraform resource: tfe_notification_configuration. S
func GetResource_tfe_notification_configuration() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "tfe_notification_configuration",
		TerraformResourceName: "tfe_notification_configuration",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			resp, err := client.tfeClient.NotificationConfigurations.List(ctx, client.WorkspaceId, &tfe.NotificationConfigurationListOptions{
				ListOptions: tfe.ListOptions{
					PageNumber: 0,
					PageSize:   10,
				},
			})
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			count := resp.TotalCount
			pages := count/10 + 1

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for i := 0; i < pages; i++ {
				resp, err := client.tfeClient.NotificationConfigurations.List(ctx, client.WorkspaceId, &tfe.NotificationConfigurationListOptions{
					ListOptions: tfe.ListOptions{
						PageNumber: i,
						PageSize:   5,
					},
				})
				if err != nil {
					return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
				}

				for _, item := range resp.Items {
					param := selefra_terraform_schema.ResourceRequestParam{
						ID: item.ID,
						ArgumentMap: map[string]interface{}{
							"name":             item.Name,
							"destination_type": item.DestinationType,
							"workspace_id":     client.WorkspaceId,
						},
					}

					// url is required if destination_type is generic, microsoft-teams, or slack
					if item.DestinationType == tfe.NotificationDestinationTypeGeneric || item.DestinationType == tfe.NotificationDestinationTypeMicrosoftTeams || item.DestinationType == tfe.NotificationDestinationTypeSlack {
						param.ArgumentMap["url"] = item.URL
					}
					resourceRequestParamSlice = append(resourceRequestParamSlice, &param)
				}
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: tfe_terraform_version
// func GetResource_tfe_terraform_version() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "tfe_terraform_version",
// 		TerraformResourceName: "tfe_terraform_version",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			//client := taskClient.(*Client)
// 			//
// 			//resp, err := client.tfeClient.Admin.TerraformVersions.List(ctx, &tfe.AdminTerraformVersionsListOptions{
// 			//	ListOptions: tfe.ListOptions{
// 			//		PageNumber: 0,
// 			//		PageSize:   10,
// 			//	},
// 			//})
// 			//if err != nil {
// 			//	return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//}
// 			//count := resp.TotalCount
// 			//pages := count/10 + 1
// 			//
// 			//resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			//for i := 0; i < pages; i++ {
// 			//	resp, err := client.tfeClient.Admin.TerraformVersions.List(ctx, &tfe.AdminTerraformVersionsListOptions{
// 			//		ListOptions: tfe.ListOptions{
// 			//			PageNumber: i,
// 			//			PageSize:   10,
// 			//		},
// 			//	})
// 			//	if err != nil {
// 			//		return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//	}
// 			//	for _, item := range resp.Items {
// 			//		resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
// 			//			ID: item.ID,
// 			//			ArgumentMap: map[string]interface{}{
// 			//				"version": item.Version,
// 			//				"url":     item.URL,
// 			//				"sha":     item.Sha,
// 			//			},
// 			//		})
// 			//	}
// 			//}
// 			//
// 			//return resourceRequestParamSlice, nil

// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: tfe_admin_organization_settings
// func GetResource_tfe_admin_organization_settings() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "tfe_admin_organization_settings",
// 		TerraformResourceName: "tfe_admin_organization_settings",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			//client := taskClient.(*Client)
// 			//
// 			//resp, err := client.tfeClient.Admin.Organizations.List(ctx, &tfe.AdminOrganizationListOptions{
// 			//	ListOptions: tfe.ListOptions{
// 			//		PageNumber: 0,
// 			//		PageSize:   10,
// 			//	},
// 			//})
// 			//if err != nil {
// 			//	return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//}
// 			//count := resp.TotalCount
// 			//pages := count/10 + 1
// 			//
// 			//resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			//for i := 0; i < pages; i++ {
// 			//	resp, err := client.tfeClient.Admin.Organizations.List(ctx, &tfe.AdminOrganizationListOptions{
// 			//		ListOptions: tfe.ListOptions{
// 			//			PageNumber: i,
// 			//			PageSize:   10,
// 			//		},
// 			//	})
// 			//	if err != nil {
// 			//		continue
// 			//	}
// 			//	if len(resp.Items) <= 0 {
// 			//		break
// 			//	}
// 			//}
// 			//
// 			//return resourceRequestParamSlice, nil

// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: tfe_workspace_variable_set
// func GetResource_tfe_workspace_variable_set() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "tfe_workspace_variable_set",
// 		TerraformResourceName: "tfe_workspace_variable_set",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			//client := taskClient.(*Client)
// 			//
// 			//resp, err := client.tfeClient.VariableSets.ListForWorkspace(ctx, client.WorkspaceId, &tfe.VariableSetListOptions{
// 			//	ListOptions: tfe.ListOptions{
// 			//		PageNumber: 0,
// 			//		PageSize:   10,
// 			//	},
// 			//})
// 			//if err != nil {
// 			//	return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//}
// 			//count := resp.TotalCount
// 			//pages := count/10 + 1
// 			//
// 			//resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			//for i := 0; i < pages; i++ {
// 			//	resp, err := client.tfeClient.VariableSets.ListForWorkspace(ctx, client.WorkspaceId, &tfe.VariableSetListOptions{
// 			//		ListOptions: tfe.ListOptions{
// 			//			PageNumber: i,
// 			//			PageSize:   10,
// 			//		},
// 			//	})
// 			//	if err != nil {
// 			//		return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//	}
// 			//	for _, item := range resp.Items {
// 			//		resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
// 			//			//ID: fmt.Sprintf("%s_%s", client.WorkspaceId, item.ID), // TODO: check
// 			//			ID: item.ID,
// 			//			ArgumentMap: map[string]interface{}{
// 			//				"variable_set_id": item.Name,
// 			//				"workspace_id":    client.WorkspaceId,
// 			//			},
// 			//		})
// 			//	}
// 			//}
// 			//
// 			//return resourceRequestParamSlice, nil

// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: tfe_team_members
// func GetResource_tfe_team_members() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "tfe_team_members",
// 		TerraformResourceName: "tfe_team_members",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			//client := taskClient.(*Client)
// 			//
// 			//resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			//
// 			//members, err := client.tfeClient.TeamMembers.List(ctx, client.TeamId)
// 			//if err != nil {
// 			//	return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//}
// 			//for _, user := range members {
// 			//	resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
// 			//		ID: client.TeamId,
// 			//		ArgumentMap: map[string]interface{}{
// 			//			"team_id":   client.TeamId,
// 			//			"usernames": user.Username,
// 			//		},
// 			//	})
// 			//}
// 			//
// 			//return resourceRequestParamSlice, nil

// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: tfe_team
// func GetResource_tfe_team() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "tfe_team",
// 		TerraformResourceName: "tfe_team",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			//client := taskClient.(*Client)
// 			//
// 			//resp, err := client.tfeClient.Teams.List(ctx, client.Organization, &tfe.TeamListOptions{
// 			//	ListOptions: tfe.ListOptions{
// 			//		PageNumber: 0,
// 			//		PageSize:   10,
// 			//	},
// 			//})
// 			//if err != nil {
// 			//	return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//}
// 			//count := resp.TotalCount
// 			//pages := count/10 + 1
// 			//
// 			//resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			//for i := 0; i < pages; i++ {
// 			//	resp, err := client.tfeClient.Teams.List(ctx, client.Organization, &tfe.TeamListOptions{
// 			//		ListOptions: tfe.ListOptions{
// 			//			PageNumber: i,
// 			//			PageSize:   10,
// 			//		},
// 			//	})
// 			//	if err != nil {
// 			//		return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//	}
// 			//	for _, item := range resp.Items {
// 			//		resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
// 			//			ID: item.ID,
// 			//			ArgumentMap: map[string]interface{}{
// 			//				"name": item.Name,
// 			//			},
// 			//		})
// 			//	}
// 			//}
// 			//
// 			//return resourceRequestParamSlice, nil

// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: tfe_team_access
// func GetResource_tfe_team_access() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "tfe_team_access",
// 		TerraformResourceName: "tfe_team_access",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			//client := taskClient.(*Client)
// 			//
// 			//resp, err := client.tfeClient.TeamAccess.List(ctx, &tfe.TeamAccessListOptions{
// 			//	ListOptions: tfe.ListOptions{
// 			//		PageNumber: 0,
// 			//		PageSize:   10,
// 			//	},
// 			//	WorkspaceID: client.WorkspaceId,
// 			//})
// 			//if err != nil {
// 			//	return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//}
// 			//count := resp.TotalCount
// 			//pages := count/10 + 1
// 			//
// 			//resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			//for i := 0; i < pages; i++ {
// 			//	resp, err := client.tfeClient.TeamAccess.List(ctx, &tfe.TeamAccessListOptions{
// 			//		ListOptions: tfe.ListOptions{
// 			//			PageNumber: i,
// 			//			PageSize:   10,
// 			//		},
// 			//		WorkspaceID: client.WorkspaceId,
// 			//	})
// 			//	if err != nil {
// 			//		return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//	}
// 			//	for _, item := range resp.Items {
// 			//		resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
// 			//			ID: fmt.Sprintf("%s/%s/%s", client.Organization, client.WorkspaceId, item.ID),
// 			//			ArgumentMap: map[string]interface{}{
// 			//				"team_id":      item.Team.ID,
// 			//				"workspace_id": item.Workspace.ID,
// 			//			},
// 			//		})
// 			//	}
// 			//}
// 			//
// 			//return resourceRequestParamSlice, nil

// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: tfe_organization_module_sharing
// func GetResource_tfe_organization_module_sharing() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "tfe_organization_module_sharing",
// 		TerraformResourceName: "tfe_organization_module_sharing",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: tfe_registry_module
// func GetResource_tfe_registry_module() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "tfe_registry_module",
// 		TerraformResourceName: "tfe_registry_module",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			//client := taskClient.(*Client)
// 			//
// 			//resp, err := client.tfeClient.RegistryModules.List(ctx, client.Organization, &tfe.RegistryModuleListOptions{
// 			//	ListOptions: tfe.ListOptions{
// 			//		PageNumber: 0,
// 			//		PageSize:   10,
// 			//	},
// 			//})
// 			//if err != nil {
// 			//	return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//}
// 			//count := resp.TotalCount
// 			//pages := count/10 + 1
// 			//
// 			//resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			//for i := 0; i < pages; i++ {
// 			//	resp, err := client.tfeClient.RegistryModules.List(ctx, client.Organization, &tfe.RegistryModuleListOptions{
// 			//		ListOptions: tfe.ListOptions{
// 			//			PageNumber: i,
// 			//			PageSize:   10,
// 			//		},
// 			//	})
// 			//	if err != nil {
// 			//		return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//	}
// 			//	for _, item := range resp.Items {
// 			//		resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
// 			//			ID:          fmt.Sprintf("%s/%s/%s/%s/%s/%s", client.Organization, item.Name, item.Namespace, item.Name, item.Provider, item.ID),
// 			//			ArgumentMap: map[string]interface{}{},
// 			//		})
// 			//	}
// 			//}
// 			//
// 			//return resourceRequestParamSlice, nil

// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: tfe_sentinel_policy
// func GetResource_tfe_sentinel_policy() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "tfe_sentinel_policy",
// 		TerraformResourceName: "tfe_sentinel_policy",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			// TODO
// 			return nil, nil
// 		},
// 	}
// }

// terraform resource: tfe_organization. S
func GetResource_tfe_organization() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "tfe_organization",
		TerraformResourceName: "tfe_organization",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			resp, err := client.tfeClient.Organizations.List(ctx, &tfe.OrganizationListOptions{
				ListOptions: tfe.ListOptions{
					PageNumber: 0,
					PageSize:   1,
				},
			})
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			count := resp.TotalCount
			pages := count/10 + 1

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for i := 0; i < pages; i++ {
				resp, err := client.tfeClient.Organizations.List(ctx, &tfe.OrganizationListOptions{
					ListOptions: tfe.ListOptions{
						PageNumber: i,
						PageSize:   10,
					},
				})
				if err != nil {
					return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
				}
				for _, item := range resp.Items {
					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						ID: client.Organization,
						ArgumentMap: map[string]interface{}{
							"name":  item.Name,
							"email": item.Email,
						},
					})
				}
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: tfe_project. S
func GetResource_tfe_project() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "tfe_project",
		TerraformResourceName: "tfe_project",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			resp, err := client.tfeClient.Projects.List(ctx, client.Organization, &tfe.ProjectListOptions{
				ListOptions: tfe.ListOptions{
					PageNumber: 1,
					PageSize:   10,
				},
			})
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			count := resp.TotalCount
			pages := count/10 + 1

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for i := 1; i <= pages; i++ {
				resp, err := client.tfeClient.Projects.List(ctx, client.Organization, &tfe.ProjectListOptions{
					ListOptions: tfe.ListOptions{
						PageNumber: i,
						PageSize:   10,
					},
				})
				if err != nil {
					return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
				}
				for _, item := range resp.Items {
					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						ID: item.ID,
						ArgumentMap: map[string]interface{}{
							"name": item.Name,
						},
					})
				}
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: tfe_ssh_key
// func GetResource_tfe_ssh_key() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "tfe_ssh_key",
// 		TerraformResourceName: "tfe_ssh_key",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			//client := taskClient.(*Client)
// 			//
// 			//// TODO: no data fetched
// 			//resp, err := client.tfeClient.SSHKeys.List(ctx, client.Organization, &tfe.SSHKeyListOptions{
// 			//	ListOptions: tfe.ListOptions{
// 			//		PageNumber: 0,
// 			//		PageSize:   10,
// 			//	},
// 			//})
// 			//if err != nil {
// 			//	return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//}
// 			//count := resp.TotalCount
// 			//pages := count/10 + 1
// 			//
// 			//resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			//for i := 0; i < pages; i++ {
// 			//	resp, err := client.tfeClient.SSHKeys.List(ctx, client.Organization, &tfe.SSHKeyListOptions{
// 			//		ListOptions: tfe.ListOptions{
// 			//			PageNumber: i,
// 			//			PageSize:   10,
// 			//		},
// 			//	})
// 			//	if err != nil {
// 			//		return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//	}
// 			//	for _, item := range resp.Items {
// 			//		resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
// 			//			ID: item.ID,
// 			//			ArgumentMap: map[string]interface{}{
// 			//				"name": item.Name,
// 			//			},
// 			//		})
// 			//	}
// 			//}
// 			//
// 			//return resourceRequestParamSlice, nil

// 			return nil, nil
// 		},
// 	}
// }

// terraform resource: tfe_organization_membership. S
func GetResource_tfe_organization_membership() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "tfe_organization_membership",
		TerraformResourceName: "tfe_organization_membership",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			resp, err := client.tfeClient.OrganizationMemberships.List(ctx, client.Organization, &tfe.OrganizationMembershipListOptions{
				ListOptions: tfe.ListOptions{
					PageNumber: 0,
					PageSize:   10,
				},
			})
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			count := resp.TotalCount
			pages := count/10 + 1

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for i := 0; i < pages; i++ {
				resp, err := client.tfeClient.OrganizationMemberships.List(ctx, client.Organization, &tfe.OrganizationMembershipListOptions{
					ListOptions: tfe.ListOptions{
						PageNumber: i,
						PageSize:   10,
					},
				})
				if err != nil {
					return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
				}
				for _, item := range resp.Items {
					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						ID: item.ID,
						ArgumentMap: map[string]interface{}{
							"email": item.Email,
						},
					})
				}
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: tfe_organization_run_task
// func GetResource_tfe_organization_run_task() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "tfe_organization_run_task",
// 		TerraformResourceName: "tfe_organization_run_task",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			//client := taskClient.(*Client)
// 			//
// 			//resp, err := client.tfeClient.RunTasks.List(ctx, client.Organization, &tfe.RunTaskListOptions{
// 			//	ListOptions: tfe.ListOptions{
// 			//		PageNumber: 0,
// 			//		PageSize:   10,
// 			//	},
// 			//})
// 			//if err != nil {
// 			//	return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//}
// 			//count := resp.TotalCount
// 			//pages := count/10 + 1
// 			//
// 			//resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			//for i := 0; i < pages; i++ {
// 			//	resp, err := client.tfeClient.RunTasks.List(ctx, client.Organization, &tfe.RunTaskListOptions{
// 			//		ListOptions: tfe.ListOptions{
// 			//			PageNumber: i,
// 			//			PageSize:   10,
// 			//		},
// 			//	})
// 			//	if err != nil {
// 			//		return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//	}
// 			//	for _, item := range resp.Items {
// 			//		resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
// 			//			ID: item.ID,
// 			//			ArgumentMap: map[string]interface{}{
// 			//				"url": item.URL,
// 			//			},
// 			//		})
// 			//	}
// 			//}
// 			//
// 			//return resourceRequestParamSlice, nil

// 			return nil, nil
// 		},
// 	}
// }

// // terraform resource: tfe_agent_pool
// func GetResource_tfe_agent_pool() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "tfe_agent_pool",
// 		TerraformResourceName: "tfe_agent_pool",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			//client := taskClient.(*Client)
// 			//
// 			//resp, err := client.tfeClient.AgentPools.List(ctx, client.Organization, &tfe.AgentPoolListOptions{
// 			//	ListOptions: tfe.ListOptions{
// 			//		PageNumber: 0,
// 			//		PageSize:   10,
// 			//	},
// 			//})
// 			//if err != nil {
// 			//	return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//}
// 			//count := resp.TotalCount
// 			//pages := count/10 + 1
// 			//
// 			//resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			//for i := 0; i < pages; i++ {
// 			//	resp, err := client.tfeClient.AgentPools.List(ctx, client.Organization, &tfe.AgentPoolListOptions{
// 			//		ListOptions: tfe.ListOptions{
// 			//			PageNumber: i,
// 			//			PageSize:   10,
// 			//		},
// 			//	})
// 			//	if err != nil {
// 			//		return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//	}
// 			//	for _, item := range resp.Items {
// 			//		resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
// 			//			ID: item.ID,
// 			//			ArgumentMap: map[string]interface{}{
// 			//				"name": item.Name,
// 			//			},
// 			//		})
// 			//	}
// 			//}
// 			//
// 			//return resourceRequestParamSlice, nil

// 			return nil, nil
// 		},
// 	}
// }

// terraform resource: tfe_policy. S
func GetResource_tfe_policy() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "tfe_policy",
		TerraformResourceName: "tfe_policy",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			resp, err := client.tfeClient.Policies.List(ctx, client.Organization, &tfe.PolicyListOptions{
				ListOptions: tfe.ListOptions{
					PageNumber: 0,
					PageSize:   10,
				},
			})
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			count := resp.TotalCount
			pages := count/10 + 1

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for i := 0; i < pages; i++ {
				resp, err := client.tfeClient.Policies.List(ctx, client.Organization, &tfe.PolicyListOptions{
					ListOptions: tfe.ListOptions{
						PageNumber: i,
						PageSize:   10,
					},
				})
				if err != nil {
					return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
				}
				for _, item := range resp.Items {
					policyBytes, err := json.Marshal(item)
					if err != nil {
						policyBytes = []byte{}
					}
					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						//ID: fmt.Sprintf("%s/%s", client.Organization, item.ID),
						ID: item.ID,
						ArgumentMap: map[string]interface{}{
							"name":   item.Name,
							"policy": string(policyBytes),
						},
					})
				}
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: tfe_workspace
// func GetResource_tfe_workspace() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "tfe_workspace",
// 		TerraformResourceName: "tfe_workspace",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			//client := taskClient.(*Client)
// 			//
// 			//resp, err := client.tfeClient.Workspaces.List(ctx, client.Organization, &tfe.WorkspaceListOptions{
// 			//	ListOptions: tfe.ListOptions{
// 			//		PageNumber: 0,
// 			//		PageSize:   10,
// 			//	},
// 			//	Include: []tfe.WSIncludeOpt{
// 			//		tfe.WSOrganization,
// 			//	},
// 			//})
// 			//if err != nil {
// 			//	return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//}
// 			//count := resp.TotalCount
// 			//pages := count/10 + 1
// 			//
// 			//resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			//for i := 0; i < pages; i++ {
// 			//	resp, err := client.tfeClient.Workspaces.List(ctx, client.Organization, &tfe.WorkspaceListOptions{
// 			//		ListOptions: tfe.ListOptions{
// 			//			PageNumber: i,
// 			//			PageSize:   10,
// 			//		},
// 			//		Include: []tfe.WSIncludeOpt{
// 			//			tfe.WSOrganization,
// 			//		},
// 			//	})
// 			//	if err != nil {
// 			//		return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//	}
// 			//	for _, item := range resp.Items {
// 			//		resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
// 			//			ID: item.ID,
// 			//			ArgumentMap: map[string]interface{}{
// 			//				"name": item.Name,
// 			//			},
// 			//		})
// 			//	}
// 			//}
// 			//
// 			//return resourceRequestParamSlice, nil

// 			return nil, nil
// 		},
// 	}
// }

// terraform resource: tfe_workspace_policy_set. S
func GetResource_tfe_workspace_policy_set() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "tfe_workspace_policy_set",
		TerraformResourceName: "tfe_workspace_policy_set",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			policySet, err := client.tfeClient.PolicySets.ReadWithOptions(ctx, client.PolicySetId, &tfe.PolicySetReadOptions{
				Include: []tfe.PolicySetIncludeOpt{tfe.PolicySetWorkspaces},
			})
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, workspace := range policySet.Workspaces {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: fmt.Sprintf("%s_%s", workspace.ID, policySet.ID),
					ArgumentMap: map[string]any{
						"policy_set_id": policySet.ID,
						"workspace_id":  workspace.ID,
					},
				})
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: tfe_variable_set. S
func GetResource_tfe_variable_set() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "tfe_variable_set",
		TerraformResourceName: "tfe_variable_set",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			resp, err := client.tfeClient.VariableSets.ListForWorkspace(ctx, client.WorkspaceId, &tfe.VariableSetListOptions{
				ListOptions: tfe.ListOptions{
					PageNumber: 0,
					PageSize:   10,
				},
			})
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			count := resp.TotalCount
			pages := count/10 + 1

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for i := 0; i < pages; i++ {
				resp, err := client.tfeClient.VariableSets.ListForWorkspace(ctx, client.WorkspaceId, &tfe.VariableSetListOptions{
					ListOptions: tfe.ListOptions{
						PageNumber: i,
						PageSize:   10,
					},
				})
				if err != nil {
					return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
				}
				for _, item := range resp.Items {
					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						ID: item.ID,
						ArgumentMap: map[string]interface{}{
							"name": item.Name,
						},
					})
				}
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: tfe_run_trigger. S
func GetResource_tfe_run_trigger() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "tfe_run_trigger",
		TerraformResourceName: "tfe_run_trigger",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			resp, err := client.tfeClient.RunTriggers.List(ctx, client.WorkspaceId, &tfe.RunTriggerListOptions{
				ListOptions: tfe.ListOptions{
					PageNumber: 0,
					PageSize:   10,
				},
				RunTriggerType: tfe.RunTriggerOutbound,
			})
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			count := resp.TotalCount
			pages := count/10 + 1

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for i := 0; i < pages; i++ {
				resp, err := client.tfeClient.RunTriggers.List(ctx, client.WorkspaceId, &tfe.RunTriggerListOptions{
					ListOptions: tfe.ListOptions{
						PageNumber: i,
						PageSize:   10,
					},
					RunTriggerType: tfe.RunTriggerOutbound,
				})
				if err != nil {
					return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
				}
				for _, item := range resp.Items {
					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						ID: item.ID,
						ArgumentMap: map[string]interface{}{
							"workspace_id":  client.WorkspaceId,
							"sourceable_id": item.Sourceable.ID,
						},
					})
				}
			}

			resp, err = client.tfeClient.RunTriggers.List(ctx, client.WorkspaceId, &tfe.RunTriggerListOptions{
				ListOptions: tfe.ListOptions{
					PageNumber: 0,
					PageSize:   10,
				},
				RunTriggerType: tfe.RunTriggerInbound,
			})
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			count = resp.TotalCount
			pages = count/10 + 1

			for i := 0; i < pages; i++ {
				resp, err := client.tfeClient.RunTriggers.List(ctx, client.WorkspaceId, &tfe.RunTriggerListOptions{
					ListOptions: tfe.ListOptions{
						PageNumber: i,
						PageSize:   10,
					},
					RunTriggerType: tfe.RunTriggerInbound,
				})
				if err != nil {
					return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
				}
				for _, item := range resp.Items {
					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						ID: item.ID,
						ArgumentMap: map[string]interface{}{
							"workspace_id":  client.WorkspaceId,
							"sourceable_id": item.Sourceable.ID,
						},
					})
				}
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: tfe_team_token. S
func GetResource_tfe_team_token() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "tfe_team_token",
		TerraformResourceName: "tfe_team_token",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			_, err := client.tfeClient.TeamTokens.Read(ctx, client.TeamId)
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
				ID: client.TeamId,
				ArgumentMap: map[string]any{
					"team_id": client.TeamId,
				},
			})

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: tfe_oauth_client. S
func GetResource_tfe_oauth_client() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "tfe_oauth_client",
		TerraformResourceName: "tfe_oauth_client",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			resp, err := client.tfeClient.OAuthClients.List(ctx, client.Organization, &tfe.OAuthClientListOptions{
				ListOptions: tfe.ListOptions{
					PageNumber: 0,
					PageSize:   10,
				},
			})
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			count := resp.TotalCount
			pages := count/10 + 1

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for i := 0; i < pages; i++ {
				resp, err := client.tfeClient.OAuthClients.List(ctx, client.Organization, &tfe.OAuthClientListOptions{
					ListOptions: tfe.ListOptions{
						PageNumber: i,
						PageSize:   10,
					},
				})
				if err != nil {
					return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
				}
				for _, item := range resp.Items {
					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						ID: item.ID,
						ArgumentMap: map[string]interface{}{
							"api_url":          item.APIURL,
							"http_url":         item.HTTPURL,
							"service_provider": item.ServiceProvider,
						},
					})
				}
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: tfe_variable FIXME: rpc error: code = Unavailable desc = error reading from server: EOF
// func GetResource_tfe_variable() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "tfe_variable",
// 		TerraformResourceName: "tfe_variable",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			//client := taskClient.(*Client)
// 			//
// 			//resp, err := client.tfeClient.Variables.List(ctx, client.WorkspaceId, &tfe.VariableListOptions{
// 			//	ListOptions: tfe.ListOptions{
// 			//		PageNumber: 0,
// 			//		PageSize:   10,
// 			//	},
// 			//})
// 			//if err != nil {
// 			//	return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//}
// 			//count := resp.TotalCount
// 			//pages := count/10 + 1
// 			//
// 			//resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			//for i := 0; i < pages; i++ {
// 			//	resp, err := client.tfeClient.Variables.List(ctx, client.WorkspaceId, &tfe.VariableListOptions{
// 			//		ListOptions: tfe.ListOptions{
// 			//			PageNumber: i,
// 			//			PageSize:   10,
// 			//		},
// 			//	})
// 			//	if err != nil {
// 			//		return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//	}
// 			//	for _, item := range resp.Items {
// 			//		resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
// 			//			ID: item.ID,
// 			//			ArgumentMap: map[string]interface{}{
// 			//				"key":      item.Key,
// 			//				"value":    item.Value,
// 			//				"category": item.Category,
// 			//			},
// 			//		})
// 			//	}
// 			//}
// 			//
// 			//return resourceRequestParamSlice, nil

// 			return nil, nil
// 		},
// 	}
// }

// terraform resource: tfe_team_organization_member. S
func GetResource_tfe_team_organization_member() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "tfe_team_organization_member",
		TerraformResourceName: "tfe_team_organization_member",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			resp, err := client.tfeClient.OrganizationMemberships.List(ctx, client.Organization, &tfe.OrganizationMembershipListOptions{
				ListOptions: tfe.ListOptions{
					PageNumber: 0,
					PageSize:   10,
				},
			})
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			count := resp.TotalCount
			pages := count/10 + 1

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for i := 0; i < pages; i++ {
				resp, err := client.tfeClient.OrganizationMemberships.List(ctx, client.Organization, &tfe.OrganizationMembershipListOptions{
					ListOptions: tfe.ListOptions{
						PageNumber: i,
						PageSize:   10,
					},
				})
				if err != nil {
					return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
				}
				for _, item := range resp.Items {
					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						ID: fmt.Sprintf("%s/%s", client.TeamId, item.ID),
						ArgumentMap: map[string]interface{}{
							"team_id ":                    client.TeamId,
							"organization_membership_id ": item.Teams[0].ID, // TODO: FIXME
						},
					})
				}
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: tfe_policy_set. S
func GetResource_tfe_policy_set() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "tfe_policy_set",
		TerraformResourceName: "tfe_policy_set",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			resp, err := client.tfeClient.PolicySets.List(ctx, client.Organization, &tfe.PolicySetListOptions{
				ListOptions: tfe.ListOptions{
					PageNumber: 0,
					PageSize:   10,
				},
			})
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			count := resp.TotalCount
			pages := count/10 + 1

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for i := 0; i < pages; i++ {
				resp, err := client.tfeClient.PolicySets.List(ctx, client.Organization, &tfe.PolicySetListOptions{
					ListOptions: tfe.ListOptions{
						PageNumber: i,
						PageSize:   10,
					},
				})
				if err != nil {
					return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
				}
				for _, item := range resp.Items {
					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						ID: item.ID,
						ArgumentMap: map[string]interface{}{
							"name": item.Name,
						},
					})
				}
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: tfe_workspace_run_task. TODO: need a url to test
// func GetResource_tfe_workspace_run_task() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "tfe_workspace_run_task",
// 		TerraformResourceName: "tfe_workspace_run_task",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			//client := taskClient.(*Client)
// 			//
// 			//resp, err := client.tfeClient.WorkspaceRunTasks.List(ctx, client.WorkspaceId, &tfe.WorkspaceRunTaskListOptions{
// 			//	ListOptions: tfe.ListOptions{
// 			//		PageNumber: 0,
// 			//		PageSize:   10,
// 			//	},
// 			//})
// 			//if err != nil {
// 			//	return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//}
// 			//count := resp.TotalCount
// 			//pages := count/10 + 1
// 			//
// 			//resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			//for i := 0; i < pages; i++ {
// 			//	resp, err := client.tfeClient.WorkspaceRunTasks.List(ctx, client.WorkspaceId, &tfe.WorkspaceRunTaskListOptions{
// 			//		ListOptions: tfe.ListOptions{
// 			//			PageNumber: i,
// 			//			PageSize:   10,
// 			//		},
// 			//	})
// 			//	if err != nil {
// 			//		return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//	}
// 			//	for _, item := range resp.Items {
// 			//		resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
// 			//			ID: item.ID,
// 			//			ArgumentMap: map[string]interface{}{
// 			//				"enforcement_level": item.EnforcementLevel,
// 			//				"task_id":           item.RunTask.ID,
// 			//				"workspace_id":      item.Workspace.ID,
// 			//			},
// 			//		})
// 			//	}
// 			//	i++
// 			//}
// 			//
// 			//return resourceRequestParamSlice, nil

// 			return nil, nil
// 		},
// 	}
// }

// terraform resource: tfe_agent_token. TODO: only Business account
// func GetResource_tfe_agent_token() *selefra_terraform_schema.SelefraTerraformResource {
// 	return &selefra_terraform_schema.SelefraTerraformResource{
// 		SelefraTableName:      "tfe_agent_token",
// 		TerraformResourceName: "tfe_agent_token",
// 		Description:           "",
// 		SubTables:             nil,
// 		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
// 			//client := taskClient.(*Client)
// 			//
// 			//resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
// 			//
// 			//resp, err := client.tfeClient.AgentTokens.List(ctx, client.AgentPoolId)
// 			//if err != nil {
// 			//	return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
// 			//}
// 			//for _, item := range resp.Items {
// 			//	resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
// 			//		ID: item.ID,
// 			//		ArgumentMap: map[string]interface{}{
// 			//			"agent_pool_id": item.ID,
// 			//			"description":   item.Description,
// 			//		},
// 			//	})
// 			//}
// 			//
// 			//return resourceRequestParamSlice, nil

// 			return nil, nil
// 		},
// 	}
// }

// terraform resource: tfe_policy_set_parameter. S
func GetResource_tfe_policy_set_parameter() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "tfe_policy_set_parameter",
		TerraformResourceName: "tfe_policy_set_parameter",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			resp, err := client.tfeClient.PolicySetParameters.List(ctx, client.PolicySetId, &tfe.PolicySetParameterListOptions{
				ListOptions: tfe.ListOptions{
					PageNumber: 0,
					PageSize:   10,
				},
			})
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			count := resp.TotalCount
			pages := count/10 + 1

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for i := 0; i < pages; i++ {
				resp, err := client.tfeClient.PolicySetParameters.List(ctx, client.PolicySetId, &tfe.PolicySetParameterListOptions{
					ListOptions: tfe.ListOptions{
						PageNumber: i,
						PageSize:   10,
					},
				})
				if err != nil {
					return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
				}
				for _, item := range resp.Items {
					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						ID: item.ID,
						ArgumentMap: map[string]interface{}{
							"key":           item.Key,
							"value":         item.Value,
							"policy_set_id": client.PolicySetId,
						},
					})
				}
				i++
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: tfe_team_project_access. S
func GetResource_tfe_team_project_access() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "tfe_team_project_access",
		TerraformResourceName: "tfe_team_project_access",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			resp, err := client.tfeClient.TeamProjectAccess.List(ctx, tfe.TeamProjectAccessListOptions{
				ListOptions: tfe.ListOptions{
					PageNumber: 0,
					PageSize:   10,
				},
				ProjectID: client.ProjectId,
			})
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			count := resp.TotalCount
			pages := count/10 + 1

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for i := 0; i < pages; i++ {
				resp, err := client.tfeClient.TeamProjectAccess.List(ctx, tfe.TeamProjectAccessListOptions{
					ListOptions: tfe.ListOptions{
						PageNumber: i,
						PageSize:   10,
					},
					ProjectID: client.ProjectId,
				})
				if err != nil {
					return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
				}
				for _, item := range resp.Items {
					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						ID: item.ID,
						ArgumentMap: map[string]interface{}{
							"team_id":    item.Team.ID,
							"project_id": item.Project.ID,
							"access":     item.Access,
						},
					})
				}
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: tfe_organization_token. S
func GetResource_tfe_organization_token() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "tfe_organization_token",
		TerraformResourceName: "tfe_organization_token",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)

			resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
				ID: client.Organization,
			})

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: tfe_team_member. S
func GetResource_tfe_team_member() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "tfe_team_member",
		TerraformResourceName: "tfe_team_member",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			members, err := client.tfeClient.TeamMembers.List(ctx, client.TeamId)
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			for _, user := range members {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: fmt.Sprintf("%s/%s", client.TeamId, user.Username),
					ArgumentMap: map[string]interface{}{
						"team_id":   client.TeamId,
						"usernames": user.Username,
					},
				})
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: tfe_team_organization_members. S
func GetResource_tfe_team_organization_members() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "tfe_team_organization_members",
		TerraformResourceName: "tfe_team_organization_members",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			resp, err := client.tfeClient.TeamMembers.ListOrganizationMemberships(ctx, client.TeamId)
			if err != nil {
				return nil, schema.NewDiagnosticsAddErrorMsg(err.Error())
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)

			ids := make([]string, 0)
			for _, item := range resp {
				ids = append(ids, item.ID)
			}

			resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
				ID: client.TeamId,
				ArgumentMap: map[string]any{
					"organization_membership_ids": ids,
				},
			})

			return resourceRequestParamSlice, nil
		},
	}
}
