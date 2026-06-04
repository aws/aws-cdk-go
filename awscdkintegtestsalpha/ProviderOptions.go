package awscdkintegtestsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Shared options for configuring the assertion provider lambda function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   providerOptions := &ProviderOptions{
//   	ProviderLogLevel: awscdk.Aws_lambda.ApplicationLogLevel_INFO,
//   }
//
// Experimental.
type ProviderOptions struct {
	// The log level of the provider lambda function.
	// Default: ApplicationLogLevel.FATAL
	//
	// Experimental.
	ProviderLogLevel awslambda.ApplicationLogLevel `field:"optional" json:"providerLogLevel" yaml:"providerLogLevel"`
}

