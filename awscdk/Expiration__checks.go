//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"
	"time"
)

func (e *jsiiProxy_Expiration) validateIsAfterParameters(t Duration) error {
	if t == nil {
		return fmt.Errorf("parameter t is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Expiration) validateIsBeforeParameters(t Duration) error {
	if t == nil {
		return fmt.Errorf("parameter t is required, but nil was provided")
	}

	return nil
}

func validateExpiration_AfterParameters(t Duration) error {
	if t == nil {
		return fmt.Errorf("parameter t is required, but nil was provided")
	}

	return nil
}

func validateExpiration_AtDateParameters(d *time.Time) error {
	if d == nil {
		return fmt.Errorf("parameter d is required, but nil was provided")
	}

	return nil
}

func validateExpiration_AtTimestampParameters(t *float64) error {
	if t == nil {
		return fmt.Errorf("parameter t is required, but nil was provided")
	}

	return nil
}

func validateExpiration_FromStringParameters(s *string) error {
	if s == nil {
		return fmt.Errorf("parameter s is required, but nil was provided")
	}

	return nil
}

