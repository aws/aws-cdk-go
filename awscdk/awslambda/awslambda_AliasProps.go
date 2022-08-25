package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a new Lambda alias.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   lambdaCode := lambda.code.fromCfnParameters()
//   func := lambda.NewFunction(this, jsii.String("Lambda"), &functionProps{
//   	code: lambdaCode,
//   	handler: jsii.String("index.handler"),
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   })
//   // used to make sure each CDK synthesis produces a different Version
//   version := func.currentVersion
//   alias := lambda.NewAlias(this, jsii.String("LambdaAlias"), &aliasProps{
//   	aliasName: jsii.String("Prod"),
//   	version: version,
//   })
//
//   codedeploy.NewLambdaDeploymentGroup(this, jsii.String("DeploymentGroup"), &lambdaDeploymentGroupProps{
//   	alias: alias,
//   	deploymentConfig: codedeploy.lambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE(),
//   })
//
type AliasProps struct {
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
	// Name of this alias.
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
	// Function version this alias refers to.
	//
	// Use lambda.currentVersion to reference a version with your latest changes.
	Version IVersion `field:"required" json:"version" yaml:"version"`
}

