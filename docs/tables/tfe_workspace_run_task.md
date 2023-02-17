# Table: tfe_workspace_run_task

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| enforcement_level | string | X | √ | The enforcement level of the task. Valid values are `advisory` and `mandatory`. | 
| id | string | X | √ |  | 
| stage | string | X | √ | The stage to run the task in. Valid values are `pre_plan`, `post_plan` and `pre_apply`. | 
| task_id | string | X | √ | The id of the Run task to associate to the Workspace. | 
| workspace_id | string | X | √ | The id of the workspace to associate the Run task to. | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


