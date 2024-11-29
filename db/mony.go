package db

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./mony.db")
	if err != nil {
		return nil, err
	}

	// SQL for creating the income table
	createIncomeTable := `
    CREATE TABLE IF NOT EXISTS income (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        num TEXT,
        title TEXT NOT NULL,
        description TEXT,
        date DATE NOT NULL,
        amount_received REAL NOT NULL,
        subtotal REAL,
        iva REAL,
        retention REAL,
        category TEXT,
        payment_method TEXT,
        is_recurring BOOLEAN,
        income_source TEXT
    );`
	_, err = db.Exec(createIncomeTable)
	if err != nil {
		return nil, err
	}

	// SQL for creating the clients table
	createClientsTable := `
    CREATE TABLE IF NOT EXISTS clients (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        long_name TEXT NOT NULL,
        short_name TEXT NOT NULL
    );`
	_, err = db.Exec(createClientsTable)
	if err != nil {
		return nil, err
	}

	log.Println("Database and tables 'income' and 'clients' created successfully.")
	return db, nil
}

func LoadIncomesFromCSV(filename string, db *sql.DB) error {
	log.Println("Opening CSV file:", filename) // Log file opening
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file %v: %v", filename, err)
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';' // Set the separator to semicolon

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV data: %v", err)
		return err
	}

	log.Println("CSV file read successfully. Processing records...")

	for i, record := range records {
		if i == 0 {
			log.Println("Skipping header row.")
			continue // Skip header row
		}

		// Adjusted indexes based on the updated CSV structure
		dateStr := record[0]
		num := strings.TrimSpace(record[1])
		client := strings.TrimSpace(record[3])
		description := strings.TrimSpace(record[4])
		subtotalStr := record[9]
		ivaStr := record[10]
		retentionStr := record[11]
		totalStr := record[14]

		// Parse date
		date, err := time.Parse("02/01/06", dateStr)
		if err != nil {
			log.Printf("Error parsing date for record %v: %v\n", dateStr, err)
			continue
		}

		title := num + " - " + client // Combine Num and Client for title

		// Parse financial amounts
		subtotal, err := parseAmount(subtotalStr)
		if err != nil {
			log.Printf("Error parsing subtotal for record %v: %v\n", subtotalStr, err)
			continue
		}

		iva, err := parseAmount(ivaStr)
		if err != nil {
			log.Printf("Error parsing IVA for record %v: %v\n", ivaStr, err)
			iva = 0.0 // Default to 0 if missing
		}

		retention, err := parseAmount(retentionStr)
		if err != nil {
			log.Printf("Error parsing retention for record %v: %v\n", retentionStr, err)
			retention = 0.0 // Default to 0 if missing
		}

		total, err := parseAmount(totalStr)
		if err != nil {
			log.Printf("Error parsing total amount for record %v: %v\n", totalStr, err)
			continue
		}

		log.Printf("Inserting record %v into database.", num)

		// Insert into the database, including Num and other updated fields
		_, err = db.Exec(
			`INSERT INTO income (num, title, description, date, amount_received, subtotal, iva, retention, category, payment_method, is_recurring, income_source) 
             VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			num, title, description, date, total, subtotal, iva, retention, "Client Payment", "Bank Transfer", false, client,
		)
		if err != nil {
			log.Printf("Error inserting record %v: %v", num, err)
			return err
		}
	}

	log.Println("CSV data loaded into the 'income' table with updated fields.")
	return nil
}
func LoadClientsFromCSV(filename string, db *sql.DB) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file %v: %v", filename, err)
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV data: %v", err)
		return err
	}

	// Map to track distinct clients by long name
	clientsMap := make(map[string]string)

	// Process each record
	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}

		description := strings.TrimSpace(record[4]) // Extract the description field

		// Determine the short name based on keywords in the description
		var shortName string
		if strings.Contains(strings.ToLower(description), "ioc") {
			shortName = "ioc"
		} else if strings.Contains(strings.ToLower(description), "aena") {
			shortName = "aena"
		} else if strings.Contains(strings.ToLower(description), "cazatucasa") {
			shortName = "cazatucasa"
		} else if strings.Contains(strings.ToLower(description), "bedfiles") {
			shortName = "bedfiles"
		} else if strings.Contains(strings.ToLower(description), "livgolf") {
			shortName = "livgolf"
		} else if strings.Contains(strings.ToLower(description), "banco") {
			shortName = "banco"
		} else {
			shortName = strings.ToLower(description)
		}

		// Only add unique descriptions to clientsMap
		if _, exists := clientsMap[description]; !exists {
			clientsMap[description] = shortName
		}
	}

	// Insert each unique client into the database
	for longName, shortName := range clientsMap {
		_, err := db.Exec(
			`INSERT INTO clients (long_name, short_name) VALUES (?, ?)`,
			longName, shortName,
		)
		if err != nil {
			log.Printf("Error inserting client %v: %v", longName, err)
			return err
		}
	}

	log.Println("Distinct clients loaded into the 'clients' table.")
	return nil
}

// Helper function to parse amounts
func parseAmount(value string) (float64, error) {
	cleanedValue := strings.ReplaceAll(value, ".", "")                          // Remove thousands separator
	cleanedValue = strings.ReplaceAll(cleanedValue, ",", ".")                   // Replace comma with dot for decimals
	cleanedValue = strings.TrimSpace(strings.ReplaceAll(cleanedValue, "€", "")) // Remove currency symbol
	if cleanedValue == "-" || cleanedValue == "" {                              // Handle empty or "-" values
		return 0.0, nil
	}
	return strconv.ParseFloat(cleanedValue, 64)
}
