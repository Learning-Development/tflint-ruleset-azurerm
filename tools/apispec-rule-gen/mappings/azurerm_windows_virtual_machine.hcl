mapping "azurerm_windows_virtual_machine" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2020-06-01/compute.json"

  admin_password             = OSProfile.adminPassword
  admin_username             = OSProfile.adminUsername
  size                       = any //HardwareProfile.vmSize
  allow_extension_operations = OSProfile.allowExtensionOperations
  computer_name              = OSProfile.computerName
  custom_data                = OSProfile.customData
  enable_automatic_updates   = WindowsConfiguration.enableAutomaticUpdates
  eviction_policy            = evictionPolicy
  license_type               = VirtualMachineProperties.licenseType
  max_bid_price              = BillingProfile.maxPrice
  priority                   = priority
  provision_vm_agent         = WindowsConfiguration.provisionVMAgent
}