package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Options for `lambda.Alias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var destination iDestination
//   var duration duration
//   var version version
//
//   aliasOptions := &AliasOptions{
//   	AdditionalVersions: []versionWeight{
//   		&versionWeight{
//   			Version: version,
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	MaxEventAge: duration,
//   	OnFailure: destination,
//   	OnSuccess: destination,
//   	ProvisionedConcurrentExecutions: jsii.Number(123),
//   	RetryAttempts: jsii.Number(123),
//   }
//
// Experimental.
type AliasOptions struct {
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
	// Additional versions with individual weights this alias points to.
	//
	// Individual additional version weights specified here should add up to
	// (less than) one. All remaining weight is routed to the default
	// version.
	//
	// For example, the config is
	//
	//     version: "1"
	//     additionalVersions: [{ version: "2", weight: 0.05 }]
	//
	// Then 5% of traffic will be routed to function version 2, while
	// the remaining 95% of traffic will be routed to function version 1.
	// Experimental.
	AdditionalVersions *[]*VersionWeight `field:"optional" json:"additionalVersions" yaml:"additionalVersions"`
	// Description for the alias.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies a provisioned concurrency configuration for a function's alias.
	// Experimental.
	ProvisionedConcurrentExecutions *float64 `field:"optional" json:"provisionedConcurrentExecutions" yaml:"provisionedConcurrentExecutions"`
}

