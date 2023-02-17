# Table: tfe_organization_run_task

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| organization | string | X | √ |  | 
| url | string | X | √ |  | 
| category | string | X | √ |  | 
| description | string | X | √ |  | 
| enabled | bool | X | √ |  | 
| hmac_key | string | X | √ |  | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


