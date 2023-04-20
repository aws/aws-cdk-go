package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for a new Lambda version.
//
// Example:
//   var fn function
//
//   version := lambda.NewVersion(this, jsii.String("MyVersion"), &VersionProps{
//   	Lambda: fn,
//   })
//
// Experimental.
type VersionProps struct {
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
	// SHA256 of the version of the Lambda source code.
	//
	// Specify to validate that you're deploying the right version.
	// Experimental.
	CodeSha256 *string `field:"optional" json:"codeSha256" yaml:"codeSha256"`
	// Description of the version.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies a provisioned concurrency configuration for a function's version.
	// Experimental.
	ProvisionedConcurrentExecutions *float64 `field:"optional" json:"provisionedConcurrentExecutions" yaml:"provisionedConcurrentExecutions"`
	// Whether to retain old versions of this function when a new version is created.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Function to get the value of.
	// Experimental.
	Lambda IFunction `field:"required" json:"lambda" yaml:"lambda"`
}

