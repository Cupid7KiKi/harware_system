// This file is generated by GoAdmin CLI adm.
package tables

import (
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"hardware_system/pages"
)

// Generators is a map of table models.
//
// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// example end
var Generators = map[string]table.Generator{
	"product_categories": pages.GetProductcategoriesTable,
	"product":            pages.GetProductsTable,
	"warehouses":         pages.GetWarehousesTable,
	"inventory":          pages.GetInventoryTable,
	"orders":             pages.GetOrdersTable,
	"order_items":        pages.GetOrderitemsTable,
	"customers":          pages.GetCustomersTable,
	"customer_contacts":  pages.GetCustomercontactsTable,
	"financial_records":  pages.GetFinancialrecordsTable,
	// generators end
}
