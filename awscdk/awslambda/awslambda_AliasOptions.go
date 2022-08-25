package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for `lambda.Alias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var destination iDestination
//   var version version
//
//   aliasOptions := &aliasOptions{
//   	additionalVersions: []versionWeight{
//   		&versionWeight{
//   			version: version,
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	maxEventAge: cdk.duration.minutes(jsii.Number(30)),
//   	onFailure: destination,
//   	onSuccess: destination,
//   	provisionedConcurrentExecutions: jsii.Number(123),
//   	retryAttempts: jsii.Number(123),
//   }
//
type AliasOptions struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The destination for failed invocations.
	OnFailure IDestination `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The destination for successful invocations.
	OnSuccess IDestination `field:"optional" json:"onSuccess" yaml:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2.
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
	AdditionalVersions *[]*VersionWeight `field:"optional" json:"additionalVersions" yaml:"additionalVersions"`
	// Description for the alias.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies a provisioned concurrency configuration for a function's alias.
	ProvisionedConcurrentExecutions *float64 `field:"optional" json:"provisionedConcurrentExecutions" yaml:"provisionedConcurrentExecutions"`
}

