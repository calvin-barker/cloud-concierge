package terraformerCLI

import terraformValueObjects "github.com/dragondrop-cloud/cloud-concierge/main/internal/implementations/terraform_value_objects"

var azureResourceGroups = map[terraformValueObjects.ResourceName]string{
	"azurerm_analysis_services_server":                           "analysis",
	"azurerm_app_service":                                        "app_service",
	"azurerm_application_gateway":                                "application_gateway",
	"azurerm_container_group":                                    "container",
	"azurerm_container_registry":                                 "container",
	"azurerm_container_registry_webhook":                         "container",
	"azurerm_cosmosdb_account":                                   "cosmosdb",
	"azurerm_cosmosdb_sql_container":                             "cosmosdb",
	"azurerm_cosmosdb_sql_database":                              "cosmosdb",
	"azurerm_cosmosdb_table":                                     "cosmosdb",
	"azurerm_mariadb_configuration":                              "database",
	"azurerm_mariadb_database":                                   "database",
	"azurerm_mariadb_firewall_rule":                              "database",
	"azurerm_mariadb_server":                                     "database",
	"azurerm_mariadb_virtual_network_rule":                       "database",
	"azurerm_mysql_configuration":                                "database",
	"azurerm_mysql_database":                                     "database",
	"azurerm_mysql_firewall_rule":                                "database",
	"azurerm_mysql_server":                                       "database",
	"azurerm_mysql_virtual_network_rule":                         "database",
	"azurerm_postgresql_configuration":                           "database",
	"azurerm_postgresql_database":                                "database",
	"azurerm_postgresql_firewall_rule":                           "database",
	"azurerm_postgresql_server":                                  "database",
	"azurerm_postgresql_virtual_network_rule":                    "database",
	"azurerm_sql_database":                                       "database",
	"azurerm_sql_active_directory_administrator":                 "database",
	"azurerm_sql_elasticpool":                                    "database",
	"azurerm_sql_failover_group":                                 "database",
	"azurerm_sql_firewall_rule":                                  "database",
	"azurerm_sql_server":                                         "database",
	"azurerm_sql_virtual_network_rule":                           "database",
	"azurerm_databricks_workspace":                               "databricks",
	"azurerm_data_factory":                                       "data_factory",
	"azurerm_data_factory_pipeline":                              "data_factory",
	"azurerm_data_factory_data_flow":                             "data_factory",
	"azurerm_data_factory_dataset_azure_blob":                    "data_factory",
	"azurerm_data_factory_dataset_binary":                        "data_factory",
	"azurerm_data_factory_dataset_cosmosdb_sqlapi":               "data_factory",
	"azurerm_data_factory_custom_dataset":                        "data_factory",
	"azurerm_data_factory_dataset_delimited_text":                "data_factory",
	"azurerm_data_factory_dataset_http":                          "data_factory",
	"azurerm_data_factory_dataset_json":                          "data_factory",
	"azurerm_data_factory_dataset_mysql":                         "data_factory",
	"azurerm_data_factory_dataset_parquet":                       "data_factory",
	"azurerm_data_factory_dataset_postgresql":                    "data_factory",
	"azurerm_data_factory_dataset_snowflake":                     "data_factory",
	"azurerm_data_factory_dataset_sql_server_table":              "data_factory",
	"azurerm_data_factory_integration_runtime_azure":             "data_factory",
	"azurerm_data_factory_integration_runtime_managed":           "data_factory",
	"azurerm_data_factory_integration_runtime_azure_ssis":        "data_factory",
	"azurerm_data_factory_integration_runtime_self_hosted":       "data_factory",
	"azurerm_data_factory_linked_service_azure_blob_storage":     "data_factory",
	"azurerm_data_factory_linked_service_azure_databricks":       "data_factory",
	"azurerm_data_factory_linked_service_azure_file_storage":     "data_factory",
	"azurerm_data_factory_linked_service_azure_function":         "data_factory",
	"azurerm_data_factory_linked_service_azure_search":           "data_factory",
	"azurerm_data_factory_linked_service_azure_sql_database":     "data_factory",
	"azurerm_data_factory_linked_service_azure_table_storage":    "data_factory",
	"azurerm_data_factory_linked_service_cosmosdb":               "data_factory",
	"azurerm_data_factory_linked_custom_service":                 "data_factory",
	"azurerm_data_factory_linked_service_data_lake_storage_gen2": "data_factory",
	"azurerm_data_factory_linked_service_key_vault":              "data_factory",
	"azurerm_data_factory_linked_service_kusto":                  "data_factory",
	"azurerm_data_factory_linked_service_mysql":                  "data_factory",
	"azurerm_data_factory_linked_service_odata":                  "data_factory",
	"azurerm_data_factory_linked_service_postgresql":             "data_factory",
	"azurerm_data_factory_linked_service_sftp":                   "data_factory",
	"azurerm_data_factory_linked_service_snowflake":              "data_factory",
	"azurerm_data_factory_linked_service_sql_server":             "data_factory",
	"azurerm_data_factory_linked_service_synapse":                "data_factory",
	"azurerm_data_factory_linked_service_web":                    "data_factory",
	"azurerm_data_factory_trigger_blob_event":                    "data_factory",
	"azurerm_data_factory_trigger_schedule":                      "data_factory",
	"azurerm_data_factory_trigger_tumbling_window":               "data_factory",
	"azurerm_managed_disk":                                       "disk",
	"azurerm_dns_a_record":                                       "dns",
	"azurerm_dns_aaaa_record":                                    "dns",
	"azurerm_dns_caa_record":                                     "dns",
	"azurerm_dns_cname_record":                                   "dns",
	"azurerm_dns_mx_record":                                      "dns",
	"azurerm_dns_ns_record":                                      "dns",
	"azurerm_dns_ptr_record":                                     "dns",
	"azurerm_dns_srv_record":                                     "dns",
	"azurerm_dns_txt_record":                                     "dns",
	"azurerm_dns_zone":                                           "dns",
	"azurerm_lb":                                                 "load_balancer",
	"azurerm_lb_backend_address_pool":                            "load_balancer",
	"azurerm_lb_nat_rule":                                        "load_balancer",
	"azurerm_lb_probe":                                           "load_balancer",
	"azurerm_eventhub_namespace":                                 "eventhub",
	"azurerm_eventhub":                                           "eventhub",
	"azurerm_eventhub_consumer_group":                            "eventhub",
	"azurerm_eventhub_namespace_authorization_rule":              "eventhub",
	"azurerm_network_interface":                                  "network_interface",
	"azurerm_network_security_group":                             "network_security_group",
	"azurerm_network_security_rule":                              "network_security_group",
	"azurerm_network_watcher":                                    "network_watcher",
	"azurerm_network_watcher_flow_log":                           "network_watcher",
	"azurerm_network_packet_capture":                             "network_watcher",
	"azurerm_private_dns_a_record":                               "private_dns",
	"azurerm_private_dns_aaaa_record":                            "private_dns",
	"azurerm_private_dns_cname_record":                           "private_dns",
	"azurerm_private_dns_mx_record":                              "private_dns",
	"azurerm_private_dns_ptr_record":                             "private_dns",
	"azurerm_private_dns_srv_record":                             "private_dns",
	"azurerm_private_dns_txt_record":                             "private_dns",
	"azurerm_private_dns_zone":                                   "private_dns",
	"azurerm_private_dns_zone_virtual_network_link":              "private_dns",
	"azurerm_private_endpoint":                                   "private_endpoint",
	"azurerm_private_link_service":                               "private_endpoint",
	"azurerm_public_ip":                                          "public_ip",
	"azurerm_public_ip_prefix":                                   "public_ip",
	"azurerm_redis_cache":                                        "redis",
	"azurerm_purview_account":                                    "purview",
	"azurerm_resource_group":                                     "resource_group",
	"azurerm_management_lock":                                    "resource_group",
	"azurerm_route_table":                                        "route_table",
	"azurerm_route":                                              "route_table",
	"azurerm_route_filter":                                       "route_table",
	"azurerm_virtual_machine_scale_set":                          "scaleset",
	"azurerm_security_center_contact":                            "security_center",
	"azurerm_security_center_subscription_pricing":               "security_center",
	"azurerm_storage_account":                                    "storage_account",
	"azurerm_storage_blob":                                       "storage_account",
	"azurerm_storage_container":                                  "storage_account",
	"azurerm_synapse_workspace":                                  "synapse",
	"azurerm_synapse_sql_pool":                                   "synapse",
	"azurerm_synapse_spark_pool":                                 "synapse",
	"azurerm_synapse_firewall_rule":                              "synapse",
	"azurerm_synapse_managed_private_endpoint":                   "synapse",
	"azurerm_synapse_private_link_hub":                           "synapse",
	"azurerm_ssh_public_key":                                     "virtual_machine",
	"azurerm_virtual_machine":                                    "virtual_machine",
	"azurerm_virtual_network":                                    "virtual_network",
	"azurerm_subnet":                                             "subnet",
	"azurerm_subnet_service_endpoint_storage_policy":             "subnet",
	"azurerm_subnet_nat_gateway_association":                     "subnet",
	"azurerm_subnet_route_table_association":                     "subnet",
	"azurerm_subnet_network_security_group_association":          "subnet",
}
