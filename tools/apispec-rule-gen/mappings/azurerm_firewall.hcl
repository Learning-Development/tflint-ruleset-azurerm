mapping "azurerm_firewall" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-07-01/azureFirewall.json"

  zones = AzureFirewall.zones
}
