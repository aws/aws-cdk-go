package awslambda


// The auth types for a function url.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//
//   fnUrl := fn.AddFunctionUrl(&FunctionUrlOptions{
//   	AuthType: lambda.FunctionUrlAuthType_NONE,
//   })
//
//   cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewFunctionUrlOrigin(fnUrl),
//   	},
//   })
//
type FunctionUrlAuthType string

const (
	// Restrict access to authenticated IAM users only.
	FunctionUrlAuthType_AWS_IAM FunctionUrlAuthType = "AWS_IAM"
	// Bypass IAM authentication to create a public endpoint.
	FunctionUrlAuthType_NONE FunctionUrlAuthType = "NONE"
)

