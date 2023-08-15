package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for an EventInvokeConfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var destination iDestination
//   var function_ function
//
//   eventInvokeConfigProps := &EventInvokeConfigProps{
//   	Function: function_,
//
//   	// the properties below are optional
//   	MaxEventAge: cdk.Duration_Minutes(jsii.Number(30)),
//   	OnFailure: destination,
//   	OnSuccess: destination,
//   	Qualifier: jsii.String("qualifier"),
//   	RetryAttempts: jsii.Number(123),
//   }
//
type EventInvokeConfigProps struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours.
	// Default: Duration.hours(6)
	//
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The destination for failed invocations.
	// Default: - no destination.
	//
	OnFailure IDestination `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The destination for successful invocations.
	// Default: - no destination.
	//
	OnSuccess IDestination `field:"optional" json:"onSuccess" yaml:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2.
	// Default: 2.
	//
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The Lambda function.
	Function IFunction `field:"required" json:"function" yaml:"function"`
	// The qualifier.
	// Default: - latest version.
	//
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

