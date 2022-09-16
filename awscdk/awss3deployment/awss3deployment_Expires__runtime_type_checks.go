//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awss3deployment

import (
	"fmt"
	"time"

	"github.com/aws/aws-cdk-go/awscdk"
)

func validateExpires_AfterParameters(t awscdk.Duration) error {
	if t == nil {
		return fmt.Errorf("parameter t is required, but nil was provided")
	}

	return nil
}

func validateExpires_AtDateParameters(d *time.Time) error {
	if d == nil {
		return fmt.Errorf("parameter d is required, but nil was provided")
	}

	return nil
}

func validateExpires_AtTimestampParameters(t *float64) error {
	if t == nil {
		return fmt.Errorf("parameter t is required, but nil was provided")
	}

	return nil
}

func validateExpires_FromStringParameters(s *string) error {
	if s == nil {
		return fmt.Errorf("parameter s is required, but nil was provided")
	}

	return nil
}

