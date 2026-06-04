package awscdkintegtestsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for a lambda function provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaFunctionProviderProps := &LambdaFunctionProviderProps{
//   	Handler: jsii.String("handler"),
//   	LogRetention: awscdk.Aws_logs.RetentionDays_ONE_DAY,
//   	ProviderLogLevel: awscdk.Aws_lambda.ApplicationLogLevel_INFO,
//   }
//
// Experimental.
type LambdaFunctionProviderProps struct {
	// The log level of the provider lambda function.
	// Default: ApplicationLogLevel.FATAL
	//
	// Experimental.
	ProviderLogLevel awslambda.ApplicationLogLevel `field:"optional" json:"providerLogLevel" yaml:"providerLogLevel"`
	// The handler to use for the lambda function.
	// Default: index.handler
	//
	// Experimental.
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
	// How long, in days, the log contents will be retained.
	// Default: - no retention days specified.
	//
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
}

