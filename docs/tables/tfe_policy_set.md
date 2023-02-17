# Table: tfe_policy_set

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| policies_path | string | X | √ |  | 
| slug | json | X | √ |  | 
| description | string | X | √ |  | 
| global | bool | X | √ |  | 
| organization | string | X | √ |  | 
| overridable | bool | X | √ |  | 
| policy_ids | json | X | √ |  | 
| workspace_ids | json | X | √ |  | 
| vcs_repo | json | X | √ |  | 
| id | string | X | √ |  | 
| kind | string | X | √ |  | 
| name | string | X | √ |  | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


