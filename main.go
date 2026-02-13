package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// User struct with 4 fields
type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Age         string    `json:"age"`
	PhoneNumber string `json:"phone_number"`
}

// Application struct with 6 fields
type Application struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Status      string `json:"status"`
}

// UserApplicationRelation struct with 8 fields
type UserApplicationRelation struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	ApplicationID int       `json:"application_id"`
	Role          string    `json:"role"`
	Permissions   []string  `json:"permissions"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Status        string    `json:"status"`
}

// Hardcoded sample data
var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com", Age: "30"},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Age: "25"},
	{ID: 3, Name: "Bob Johnson", Email: "bob@example.com", Age: "35"},
}

var applications = []Application{
	{ID: 1, Name: "Task Manager", Version: "1.0.0", Description: "A simple task management app", Category: "Productivity", Status: "Active"},
	{ID: 2, Name: "Code Editor", Version: "2.1.0", Description: "Advanced code editing tool", Category: "Development", Status: "Active"},
	{ID: 3, Name: "Data Analyzer", Version: "1.5.0", Description: "Data analysis and visualization", Category: "Analytics", Status: "Beta"},
	{ID: 4, Name: "Chat App", Version: "3.0.0", Description: "Real-time messaging application", Category: "Communication", Status: "Active"},
	{ID: 5, Name: "File Storage", Version: "2.0.0", Description: "Cloud file storage service", Category: "Storage", Status: "Active"},
	{ID: 6, Name: "Project Tracker", Version: "1.2.0", Description: "Project management tool", Category: "Productivity", Status: "Active"},
}

var userApplicationRelations = []UserApplicationRelation{
	{
		ID:            1,
		UserID:        1,
		ApplicationID: 1,
		Role:          "Admin",
		Permissions:   []string{"read", "write", "delete", "admin"},
		StartDate:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:       time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		Status:        "Active",
	},
	{
		ID:            2,
		UserID:        1,
		ApplicationID: 2,
		Role:          "Developer",
		Permissions:   []string{"read", "write"},
		StartDate:     time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
		EndDate:       time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		Status:        "Active",
	},
	{
		ID:            3,
		UserID:        2,
		ApplicationID: 3,
		Role:          "User",
		Permissions:   []string{"read"},
		StartDate:     time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
		EndDate:       time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		Status:        "Active",
	},
	{
		ID:            4,
		UserID:        2,
		ApplicationID: 4,
		Role:          "Moderator",
		Permissions:   []string{"read", "write", "moderate"},
		StartDate:     time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
		EndDate:       time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		Status:        "Active",
	},
	{
		ID:            5,
		UserID:        3,
		ApplicationID: 5,
		Role:          "User",
		Permissions:   []string{"read", "write"},
		StartDate:     time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
		EndDate:       time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		Status:        "Active",
	},
	{
		ID:            6,
		UserID:        3,
		ApplicationID: 6,
		Role:          "Manager",
		Permissions:   []string{"read", "write", "manage"},
		StartDate:     time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
		EndDate:       time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		Status:        "Active",
	},
	{
		ID:            7,
		UserID:        1,
		ApplicationID: 4,
		Role:          "User",
		Permissions:   []string{"read"},
		StartDate:     time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC),
		EndDate:       time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		Status:        "Active",
	},
	{
		ID:            8,
		UserID:        2,
		ApplicationID: 6,
		Role:          "Viewer",
		Permissions:   []string{"read"},
		StartDate:     time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
		EndDate:       time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		Status:        "Active",
	},
}

// HTTP handlers
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getApplications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(applications)
}

func getUserApplicationRelations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userApplicationRelations)
}

func main() {
	// Set up routes
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/applications", getApplications)
	http.HandleFunc("/userApplicationRelations", getUserApplicationRelations)

	// Add a simple health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	fmt.Println("Server is listening on port 8080...")
	fmt.Println("Available Endpoints:")
	fmt.Println("  GET /users - Get all users")
	fmt.Println("  GET /applications - Get all applications")
	fmt.Println("  GET /userApplicationRelations - Get all user-application relations")
	fmt.Println("  GET /health - Health check")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
