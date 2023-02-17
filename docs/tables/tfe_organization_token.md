# Table: tfe_organization_token

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ |  | 
| organization | string | X | √ |  | 
| token | string | X | √ |  | 
| force_regenerate | bool | X | √ |  | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


