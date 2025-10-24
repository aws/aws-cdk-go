package awslambda


// The auth types for a function url.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   var fn Function
//
//
//   fnUrl := fn.AddFunctionUrl(&FunctionUrlOptions{
//   	AuthType: lambda.FunctionUrlAuthType_AWS_IAM,
//   })
//
//   cloudfront.NewDistribution(this, jsii.String("MyDistribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.FunctionUrlOrigin_WithOriginAccessControl(fnUrl),
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

