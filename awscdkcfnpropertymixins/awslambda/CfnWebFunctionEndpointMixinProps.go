package awslambda


// Properties for CfnWebFunctionEndpointPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnWebFunctionEndpointMixinProps := &CfnWebFunctionEndpointMixinProps{
//   	AuthType: jsii.String("authType"),
//   	Description: jsii.String("description"),
//   	EndpointName: jsii.String("endpointName"),
//   	EndpointType: jsii.String("endpointType"),
//   	FunctionName: jsii.String("functionName"),
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   	RevisionWeights: []interface{}{
//   		&RevisionWeightProperty{
//   			RevisionId: jsii.String("revisionId"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	ScalingConfig: &ScalingConfigProperty{
//   		MaxEnvironments: jsii.Number(123),
//   	},
//   	ThrottleConfig: &ThrottleConfigProperty{
//   		RateLimit: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionendpoint.html
//
type CfnWebFunctionEndpointMixinProps struct {
	// The authentication type for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionendpoint.html#cfn-lambda-webfunctionendpoint-authtype
	//
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// A description of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionendpoint.html#cfn-lambda-webfunctionendpoint-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionendpoint.html#cfn-lambda-webfunctionendpoint-endpointname
	//
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// The type of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionendpoint.html#cfn-lambda-webfunctionendpoint-endpointtype
	//
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
	// The name of the web function this endpoint belongs to.
	//
	// The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionendpoint.html#cfn-lambda-webfunctionendpoint-functionname
	//
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// The list of AWS Regions for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionendpoint.html#cfn-lambda-webfunctionendpoint-regions
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// List of revision routing entries.
	//
	// 1 or 2 entries. With 1 entry, weight must be 100. With 2 entries, weights must sum to 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionendpoint.html#cfn-lambda-webfunctionendpoint-revisionweights
	//
	RevisionWeights interface{} `field:"optional" json:"revisionWeights" yaml:"revisionWeights"`
	// The scaling configuration for the endpoint.
	//
	// Optionally constrains how many concurrent execution environments the endpoint can use, in addition to your account's vCPU quota.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionendpoint.html#cfn-lambda-webfunctionendpoint-scalingconfig
	//
	ScalingConfig interface{} `field:"optional" json:"scalingConfig" yaml:"scalingConfig"`
	// The throttling configuration for the endpoint.
	//
	// Optionally constrains the request rate that the endpoint accepts, in addition to your account's rate limit quota.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionendpoint.html#cfn-lambda-webfunctionendpoint-throttleconfig
	//
	ThrottleConfig interface{} `field:"optional" json:"throttleConfig" yaml:"throttleConfig"`
}

