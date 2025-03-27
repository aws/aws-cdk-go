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
//   aliasOptions := &AliasOptions{
//   	AdditionalVersions: []versionWeight{
//   		&versionWeight{
//   			Version: version,
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	MaxEventAge: cdk.Duration_Minutes(jsii.Number(30)),
//   	OnFailure: destination,
//   	OnSuccess: destination,
//   	ProvisionedConcurrentExecutions: jsii.Number(123),
//   	RetryAttempts: jsii.Number(123),
//   }
//
type AliasOptions struct {
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
	// Additional versions with individual weights this alias points to.
	//
	// Individual additional version weights specified here should add up to
	// (less than) one. All remaining weight is routed to the default
	// version.
	//
	// For example, the config is
	//
	//    version: "1"
	//    additionalVersions: [{ version: "2", weight: 0.05 }]
	//
	// Then 5% of traffic will be routed to function version 2, while
	// the remaining 95% of traffic will be routed to function version 1.
	// Default: No additional versions.
	//
	AdditionalVersions *[]*VersionWeight `field:"optional" json:"additionalVersions" yaml:"additionalVersions"`
	// Description for the alias.
	// Default: No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies a provisioned concurrency configuration for a function's alias.
	// Default: No provisioned concurrency.
	//
	ProvisionedConcurrentExecutions *float64 `field:"optional" json:"provisionedConcurrentExecutions" yaml:"provisionedConcurrentExecutions"`
}

