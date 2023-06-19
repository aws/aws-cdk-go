package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for an EventInvokeConfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var destination iDestination
//   var duration duration
//   var function_ function
//
//   eventInvokeConfigProps := &EventInvokeConfigProps{
//   	Function: function_,
//
//   	// the properties below are optional
//   	MaxEventAge: duration,
//   	OnFailure: destination,
//   	OnSuccess: destination,
//   	Qualifier: jsii.String("qualifier"),
//   	RetryAttempts: jsii.Number(123),
//   }
//
// Experimental.
type EventInvokeConfigProps struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours.
	// Experimental.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The destination for failed invocations.
	// Experimental.
	OnFailure IDestination `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The destination for successful invocations.
	// Experimental.
	OnSuccess IDestination `field:"optional" json:"onSuccess" yaml:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2.
	// Experimental.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The Lambda function.
	// Experimental.
	Function IFunction `field:"required" json:"function" yaml:"function"`
	// The qualifier.
	// Experimental.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

