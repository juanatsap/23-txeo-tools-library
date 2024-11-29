package db

import (
	"database/sql"
	"log"
	"time"
	"txeo-tools-library/models"

	_ "github.com/mattn/go-sqlite3"
)

// InitSQLite initializes the SQLite database and retrieves distinct categories
func InitSQLite() (*sql.DB, models.Clients, models.Incomes, models.Categories, error) {
	db, err := sql.Open("sqlite3", "./mony.db")

	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Prepare statement to query all income records
	incomeStmt, err := db.Prepare("SELECT * FROM income")
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Prepare statement to query all client records
	clientStmt, err := db.Prepare("SELECT * FROM clients")
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Retrieve categories dynamically from income table
	categories, err := GetCategories(db)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	log.Println("SQLite initialized with income statement and categories retrieved.")

	incomes := getIncomes(incomeStmt)
	clients := getClients(clientStmt)

	// Associate clients with incomes using index-based modification
	for i := range incomes {
		for _, client := range clients {
			if client.LongName == incomes[i].Description {
				incomes[i].Client = client
				break
			}
		}
	}

	return db, clients, incomes, categories, nil
}

// GetCategories retrieves all categories from the database
func GetCategories(db *sql.DB) (models.Categories, error) {
	rows, err := db.Query("SELECT DISTINCT category FROM income")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}

		// Create a Category instance for each distinct category
		cat := models.Category{
			Name:    name,
			Count:   0,     // Placeholder for count, if not required, this can be removed
			Deleted: false, // Default value for Deleted
		}
		categories = append(categories, cat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
func getIncomes(incomeStmt *sql.Stmt) models.Incomes {
	rows, err := incomeStmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var incomes []models.Income
	for rows.Next() {
		var id int
		var num string
		var title string
		var description string
		var date string
		var amountReceived float64
		var subtotal float64
		var iva float64
		var retention float64
		var category string
		var paymentMethod string
		var isRecurring bool
		var incomeSource string

		if err := rows.Scan(&id, &num, &title, &description, &date, &amountReceived, &subtotal, &iva, &retention, &category, &paymentMethod, &isRecurring, &incomeSource); err != nil {
			log.Fatal(err)
		}

		// Parse the date using RFC3339 format
		dateToDate, err := time.Parse(time.RFC3339, date)
		if err != nil {
			log.Printf("Error parsing date for record %v: %v\n", date, err)
			continue
		}

		incomes = append(incomes, models.Income{
			ID:             id,
			Num:            num,
			Title:          title,
			Description:    description,
			Date:           dateToDate,
			AmountReceived: amountReceived,
			Subtotal:       subtotal,
			IVA:            iva,
			Retention:      retention,
			Category:       category,
			PaymentMethod:  paymentMethod,
			IsRecurring:    isRecurring,
			IncomeSource:   incomeSource,
		})
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Sort the incomes by Num
	for i := 0; i < len(incomes)-1; i++ {
		for j := i + 1; j < len(incomes); j++ {
			if incomes[i].Num > incomes[j].Num {
				incomes[i], incomes[j] = incomes[j], incomes[i]
			}
		}
	}
	return incomes
}
func getClients(clientStmt *sql.Stmt) models.Clients {
	rows, err := clientStmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var clients models.Clients
	for rows.Next() {
		var id int
		var longName string
		var shortName string
		if err := rows.Scan(&id, &longName, &shortName); err != nil {
			log.Fatal(err)
		}
		clients = append(clients, models.Client{
			ID:        id,
			LongName:  longName,
			ShortName: shortName,
		})
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return clients
}
