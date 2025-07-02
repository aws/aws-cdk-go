package awslambda


// Environment variables options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentOptions := &EnvironmentOptions{
//   	RemoveInEdge: jsii.Boolean(false),
//   }
//
type EnvironmentOptions struct {
	// When used in Lambda@Edge via edgeArn() API, these environment variables will be removed.
	//
	// If not set, an error will be thrown.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-requirements-limits.html#lambda-requirements-lambda-function-configuration
	//
	// Default: false - using the function in Lambda@Edge will throw.
	//
	RemoveInEdge *bool `field:"optional" json:"removeInEdge" yaml:"removeInEdge"`
}

