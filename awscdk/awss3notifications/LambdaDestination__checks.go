//go:build !no_runtime_type_checking

package awss3notifications

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

func (l *jsiiProxy_LambdaDestination) validateBindParameters(_scope constructs.Construct, bucket awss3.IBucket) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	return nil
}

func validateNewLambdaDestinationParameters(fn awslambda.IFunction) error {
	if fn == nil {
		return fmt.Errorf("parameter fn is required, but nil was provided")
	}

	return nil
}

