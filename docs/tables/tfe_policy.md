# Table: tfe_policy

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | Text describing the policy's purpose | 
| enforce_mode | string | X | √ | The enforcement configuration of the policy. For Sentinel, valid values are `hard-mandatory`, `soft-mandatory` and `advisory`. For OPA, Valid values are ``mandatory` and `advisory`` | 
| id | string | X | √ |  | 
| kind | string | X | √ | The policy-as-code framework for the policy. Valid values are sentinel and opa | 
| name | string | X | √ | The name of the policy | 
| organization | string | X | √ | Name of the organization that this policy belongs to | 
| policy | string | X | √ | Text of a valid Sentinel or OPA policy | 
| query | string | X | √ | The OPA query to run. Required for OPA policies | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


