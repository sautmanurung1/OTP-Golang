package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	// OTPExpiry is the duration after which the OTP will expire
	OTPExpiry = 3 * time.Minute
	// OTPLength is the length of the OTP
	OTPLength = 6
)

// OTPError is the error type for OTP related errors
type OTPError struct {
	Message string
}

func (e OTPError) Error() string {
	return e.Message
}

// GenerateOTP generates a new OTP and returns it as a string
func GenerateOTP(t time.Time) (string, error) {
	rand.Seed(time.Now().UnixNano())
	otp := fmt.Sprintf("%06d", rand.Intn(1000000))
	return otp, nil
}

// ValidateOTP checks if the given OTP is valid
func ValidateOTP(otp string, now time.Time) error {
	// Check if the OTP is the correct length
	if len(otp) != OTPLength {
		return &OTPError{"Invalid length"}
	}

	// Check if the OTP has expired
	expiry := now.Add(-OTPExpiry)
	otps, _ := GenerateOTP(expiry)
	if otp == otps {
		return nil
	} else if now.Before(expiry) {
		return &OTPError{"Expired"}
	} else {
		return &OTPError{"Invalid the code OTP"}
	}
}

func main() {
	// Generate a new OTP
	now := time.Now()
	otp, err := GenerateOTP(now)
	if err != nil {
		fmt.Println("Error generating OTP:", err)
		return
	}
	fmt.Println("Generated OTP:", otp)

	fmt.Print("Enter OTP: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputOTP := scanner.Text()

		// Validate the OTP
		err = ValidateOTP(inputOTP, now)
		if err == nil {
			fmt.Println("OTP is invalid")
		} else {
			fmt.Println("OTP is valid")
		}
	}

}
