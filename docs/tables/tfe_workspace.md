# Table: tfe_workspace

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| allow_destroy_plan | bool | X | √ |  | 
| description | string | X | √ |  | 
| structured_run_output_enabled | bool | X | √ |  | 
| tag_names | json | X | √ |  | 
| trigger_prefixes | json | X | √ |  | 
| vcs_repo | json | X | √ |  | 
| auto_apply | bool | X | √ |  | 
| execution_mode | string | X | √ |  | 
| global_remote_state | bool | X | √ |  | 
| name | string | X | √ |  | 
| organization | string | X | √ |  | 
| speculative_enabled | bool | X | √ |  | 
| ssh_key_id | string | X | √ |  | 
| file_triggers_enabled | bool | X | √ |  | 
| remote_state_consumer_ids | json | X | √ |  | 
| working_directory | string | X | √ |  | 
| agent_pool_id | string | X | √ |  | 
| assessments_enabled | bool | X | √ |  | 
| force_delete | bool | X | √ |  | 
| id | string | X | √ |  | 
| operations | bool | X | √ |  | 
| project_id | string | X | √ |  | 
| queue_all_runs | bool | X | √ |  | 
| resource_count | float | X | √ |  | 
| terraform_version | string | X | √ |  | 
| trigger_patterns | json | X | √ |  | 
| selefra_terraform_original_result | json | X | √ | save terraform original result for compatibility | 


