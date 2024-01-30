package application

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

type ConfigApplicationMigrate struct {
	Db               *mysql.Config
	FilePathCustomer string
	FilePathProduct  string
	FilePathInvoice  string
	FilePathSale     string
}

// NewApplicationMigrate returns a new ApplicationMigrate
func NewApplicationMigrate(config *ConfigApplicationMigrate) (a *ApplicationMigrate) {
	a = &ApplicationMigrate{
		config: config,
	}
	return
}

type ApplicationMigrate struct {
	// config is the configuration of the application
	config *ConfigApplicationMigrate
	// database is the database to load the data
	database *sql.DB
	// fileCustomer is the path to the file that contains the customers
	fileCustomer *os.File
	// fileProduct is the path to the file that contains the products
	fileProduct *os.File
	// fileInvoice is the path to the file that contains the invoices
	fileInvoice *os.File
	// fileSales is the path to the file that contains the sales
	fileSales *os.File
	// Migrators
	//migrators []internal.Migrator
}

func (a *ApplicationMigrate) TearDown() {
	// - close files
	if a.fileCustomer != nil {
		a.fileCustomer.Close()
	}
	if a.fileProduct != nil {
		a.fileProduct.Close()
	}
	if a.fileInvoice != nil {
		a.fileInvoice.Close()
	}
	if a.fileSales != nil {
		a.fileSales.Close()
	}
	// - close db
	if a.database != nil {
		a.database.Close()
	}
}
