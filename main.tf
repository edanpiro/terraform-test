resource "azurerm_resource_group" "test" {
  name     = "${var.resource-group}"
  location = "${var.location}"
}

resource "azurerm_storage_account" "tessa" {
  name                     = "storagedemopacifico20"
  resource_group_name      = "${azurerm_resource_group.test.name}"
  location                 = "${var.location}"
  account_tier             = "Standard"
  account_replication_type = "GRS"
  
  tags {
    environment = "production"
  }
}
