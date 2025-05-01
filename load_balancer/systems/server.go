package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
)

// ! run_backends.ps1 file automatically opens 3 powershell windows and spins up the 3 backend servers for you

// quadraticFormula generates 3 random numbers and calculates the quadratic formula
func quadraticFormula(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// generate random values for a, b, and c
	a := rand.Intn(10) + 1
	b := rand.Intn(10) + 1
	c := rand.Intn(10) + 1

	log.Printf("Calculating quadratic formula for a=%d, b=%d, c=%d", a, b, c)

	// quadratic formula: x = (-b ± sqrt(b^2 - 4ac)) / 2a
	discriminant := float64(b*b - 4*a*c)
	if discriminant < 0 {
		_, err := fmt.Fprintf(w, "Backend 1 on port %s: No real solutions for a=%d, b=%d, c=%d\n", port, a, b, c)
		if err != nil {
			return
		}
		return
	}

	sqrtDisc := math.Sqrt(discriminant)
	x1 := (-float64(b) + sqrtDisc) / (2 * float64(a))
	x2 := (-float64(b) - sqrtDisc) / (2 * float64(a))

	_, err := fmt.Fprintf(w, "Backend 1 on port %s: Solutions for a=%d, b=%d, c=%d are x1 = %.2f, x2 = %.2f\n", port, a, b, c, x1, x2)
	if err != nil {
		return
	}
}

// pythagoreanTheorem generates 2 random numbers and calculates the pythagorean theorem
func pythagoreanTheorem(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	// Generate random values for sides a and b
	a := rand.Intn(10) + 1
	b := rand.Intn(10) + 1

	log.Printf("Calculating pythagorean formula for a=%d, b=%d", a, b)

	// pythagorean theorem: c = sqrt(a^2 + b^2)
	c := math.Sqrt(float64(a*a + b*b))

	_, err := fmt.Fprintf(w, "Backend 2 on port %s: For a=%d and b=%d, the hypotenuse c = %.2f\n", port, a, b, c)
	if err != nil {
		return
	}
}

// compoundInterest generates 3 random numbers and calculates the compound interest for it
func compoundInterest(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	// generate random values for P (principal), R (rate), and n (years)
	P := rand.Intn(1000) + 100 // principal amount between 100 and 1000
	R := rand.Intn(10) + 1     // interest rate between 1% and 10%
	n := rand.Intn(20) + 1     // number of years between 1 and 20

	log.Printf("Calculating compound interest for P=%d, R=%d, n=%d", P, R, n)

	// Compound Interest formula: C.I = P * (1 + R/100)^n - P
	CI := float64(P)*math.Pow(1+float64(R)/100, float64(n)) - float64(P)

	_, err := fmt.Fprintf(w, "Backend 3 on port %s: For principal P=%d, rate R=%d, and years n=%d, the compound interest = %.2f\n", port, P, R, n, CI)
	if err != nil {
		return
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch port {
		case "8081":
			quadraticFormula(w, r) // backend server handling quadratic formula
		case "8082":
			pythagoreanTheorem(w, r) // backend server handling Pythagorean theorem
		case "8083":
			compoundInterest(w, r) // backend server handling compound interest
		default:
			http.Error(w, "Port not supported", http.StatusBadRequest)
		}
	})

	log.Printf("Backend server listening on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
