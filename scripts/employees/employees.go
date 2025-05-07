package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	// Create a CSV file
	file, err := os.Create("synthetic_employee_data.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the CSV header
	header := []string{
		"EmployeeID",
		"Age",
		"Gender",
		"Ethnicity",
		"Department",
		"JobTitle",
		"PerformanceScore",
		"TenureYears",
		"DisabilityStatus",
		"VeteranStatus",
		"Location",
		"LGBTStatus",
		"Promotion",
		"InclusionProgramParticipant",
		"Salary",
	}
	err = writer.Write(header)
	if err != nil {
		fmt.Println("Error writing CSV header:", err)
		return
	}

	// Define the number of employees to generate
	numEmployees := 10500

	// Generate synthetic data and write to CSV
	for i := 0; i < numEmployees; i++ {
		employeeID := strconv.Itoa(10011 + i)
		age := strconv.Itoa(rand.Intn(43) + 22)
		gender := randGender()
		ethnicity := randEthnicity()
		department := randDepartment()
		jobTitle := randJobTitle()
		performanceScore := strconv.Itoa(rand.Intn(5) + 1)
		tenureYears := strconv.FormatFloat(rand.Float64()*19.5+0.5, 'f', 1, 64)
		disabilityStatus := randBool()
		veteranStatus := randBool()
		location := randLocation()
		lgbtStatus := randBool()
		promotion := randBool()
		inclusionProgramParticipant := randBool()
		salary := strconv.Itoa(rand.Intn(200000) + 88000)

		record := []string{
			employeeID,
			age,
			gender,
			ethnicity,
			department,
			jobTitle,
			performanceScore,
			tenureYears,
			disabilityStatus,
			veteranStatus,
			location,
			lgbtStatus,
			promotion,
			inclusionProgramParticipant,
			salary,
		}

		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error writing CSV record:", err)
			return
		}
	}
}

// Helper function to generate random gender
func randGender() string {
	genders := []string{"Male", "Female"}
	return genders[rand.Intn(len(genders))]
}

// Helper function to generate random ethnicity
func randEthnicity() string {
	ethnicities := []string{"White", "Black", "Hispanic", "Asian", "Indian"}
	return ethnicities[rand.Intn(len(ethnicities))]
}

// Helper function to generate random department
func randDepartment() string {
	departments := []string{"Sales", "Marketing", "Engineering", "Finance", "HR", "Legal"}
	return departments[rand.Intn(len(departments))]
}

// Helper function to generate random job title
func randJobTitle() string {
	jobTitles := []string{"Manager", "Engineer", "Analyst", "Specialist", "Leadership"}
	return jobTitles[rand.Intn(len(jobTitles))]
}

// Helper function to generate random location
func randLocation() string {
	locations := []string{"New York", "Los Angeles", "Chicago", "Houston", "Atlanta", "Austin"}
	return locations[rand.Intn(len(locations))]
}

// Helper function to generate random boolean value
func randBool() string {
	if rand.Intn(2) == 0 {
		return "Yes"
	}
	return "No"
}
