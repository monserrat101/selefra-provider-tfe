# Table: tfe_organization

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| session_timeout_minutes | float | X | √ |  | 
| allow_force_delete_workspaces | bool | X | √ |  | 
| assessments_enforced | bool | X | √ |  | 
| email | string | X | √ |  | 
| name | string | X | √ |  | 
| owners_team_saml_role_id | string | X | √ |  | 
| session_remember_minutes | float | X | √ |  | 
| collaborator_auth_policy | string | X | √ |  | 
| cost_estimation_enabled | bool | X | √ |  | 
| default_project_id | string | X | √ |  | 
| id | string | X | √ |  | 
| send_passing_statuses_for_untriggered_speculative_plans | bool | X | √ |  | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


