package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for `lambda.Version`.
//
// Example:
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	CurrentVersionOptions: &VersionOptions{
//   		RemovalPolicy: awscdk.RemovalPolicy_RETAIN,
//   		 // retain old versions
//   		RetryAttempts: jsii.Number(1),
//   	},
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   })
//
//   fn.AddAlias(jsii.String("live"))
//
type VersionOptions struct {
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
	// SHA256 of the version of the Lambda source code.
	//
	// Specify to validate that you're deploying the right version.
	// Default: No validation is performed.
	//
	CodeSha256 *string `field:"optional" json:"codeSha256" yaml:"codeSha256"`
	// Description of the version.
	// Default: Description of the Lambda.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies a provisioned concurrency configuration for a function's version.
	// Default: No provisioned concurrency.
	//
	ProvisionedConcurrentExecutions *float64 `field:"optional" json:"provisionedConcurrentExecutions" yaml:"provisionedConcurrentExecutions"`
	// Whether to retain old versions of this function when a new version is created.
	// Default: RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

