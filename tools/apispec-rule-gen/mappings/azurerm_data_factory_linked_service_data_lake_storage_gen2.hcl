mapping "azurerm_data_factory_linked_service_data_lake_storage_gen2" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/datafactory.json"

  name                = linkedServiceName
  resource_group_name = resourceGroupName
  data_factory_name   = factoryName
}