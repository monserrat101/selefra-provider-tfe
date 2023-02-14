package resources

import (
	"context"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/terraform/bridge"
)

func GetSelefraProvider() *provider.Provider {
	diagnostics := schema.NewDiagnostics()
	selefraProvider, d := GetSelefraTerraformProvider().ToSelefraProvider(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) *bridge.TerraformBridge {
		return taskClient.(*Client).TerraformBridge
	})
	if diagnostics.AddDiagnostics(d).HasError() {
		panic(diagnostics.ToString())
	}

    selefraProvider.TableList = GetSelefraTables()

	return selefraProvider
}

func GetSelefraTables() []*schema.Table {

    diagnostics := schema.NewDiagnostics()
    tables := make([]*schema.Table, 0)
    var table *schema.Table
    var d *schema.Diagnostics

    
    table, d = TableSchemaGenerator_tfe_notification_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_terraform_version()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_admin_organization_settings()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_workspace_variable_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_team_members()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_team()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_team_access()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_organization_module_sharing()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_registry_module()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_sentinel_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_organization()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_project()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_ssh_key()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_organization_membership()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_organization_run_task()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_agent_pool()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_workspace()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_workspace_policy_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_variable_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_run_trigger()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_team_token()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_oauth_client()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_variable()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_team_organization_member()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_policy_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_workspace_run_task()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_agent_token()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_policy_set_parameter()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_team_project_access()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_organization_token()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_team_member()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_tfe_team_organization_members()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    

    if diagnostics.HasError() {
        panic(diagnostics.ToString())
    }

	return tables
}
