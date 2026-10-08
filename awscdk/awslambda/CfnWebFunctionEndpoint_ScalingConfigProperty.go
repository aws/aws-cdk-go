package awslambda


// The scaling configuration for the endpoint.
//
// Optionally constrains how many concurrent execution environments the endpoint can use, in addition to your account's vCPU quota.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingConfigProperty := &ScalingConfigProperty{
//   	MaxEnvironments: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionendpoint-scalingconfig.html
//
type CfnWebFunctionEndpoint_ScalingConfigProperty struct {
	// The maximum number of concurrent execution environments for the endpoint.
	//
	// This optional limit further constrains the endpoint's scaling. When omitted, the endpoint's scaling is limited only by your account's vCPU quota.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionendpoint-scalingconfig.html#cfn-lambda-webfunctionendpoint-scalingconfig-maxenvironments
	//
	MaxEnvironments *float64 `field:"optional" json:"maxEnvironments" yaml:"maxEnvironments"`
}

