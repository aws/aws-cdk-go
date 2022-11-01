//go:build !no_runtime_type_checking

package awscodepipelineactions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

func validateCacheControl_FromStringParameters(s *string) error {
	if s == nil {
		return fmt.Errorf("parameter s is required, but nil was provided")
	}

	return nil
}

func validateCacheControl_MaxAgeParameters(t awscdk.Duration) error {
	if t == nil {
		return fmt.Errorf("parameter t is required, but nil was provided")
	}

	return nil
}

func validateCacheControl_SMaxAgeParameters(t awscdk.Duration) error {
	if t == nil {
		return fmt.Errorf("parameter t is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CacheControl) validateSetValueParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

