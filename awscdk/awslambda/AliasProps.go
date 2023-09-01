package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a new Lambda alias.
//
// Example:
//   lambdaCode := lambda.Code_FromCfnParameters()
//   func := lambda.NewFunction(this, jsii.String("Lambda"), &FunctionProps{
//   	Code: lambdaCode,
//   	Handler: jsii.String("index.handler"),
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   })
//   // used to make sure each CDK synthesis produces a different Version
//   version := func.currentVersion
//   alias := lambda.NewAlias(this, jsii.String("LambdaAlias"), &AliasProps{
//   	AliasName: jsii.String("Prod"),
//   	Version: Version,
//   })
//
//   codedeploy.NewLambdaDeploymentGroup(this, jsii.String("DeploymentGroup"), &LambdaDeploymentGroupProps{
//   	Alias: Alias,
//   	DeploymentConfig: codedeploy.LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE(),
//   })
//
type AliasProps struct {
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
	// Name of this alias.
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
	// Function version this alias refers to.
	//
	// Use lambda.currentVersion to reference a version with your latest changes.
	Version IVersion `field:"required" json:"version" yaml:"version"`
}

