package awslambda


// Per-region endpoint information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   regionalEndpointProperty := &RegionalEndpointProperty{
//   	AuthType: jsii.String("authType"),
//   	DomainName: jsii.String("domainName"),
//   	RevisionWeights: []interface{}{
//   		&RevisionWeightProperty{
//   			RevisionId: jsii.String("revisionId"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	ScalingConfig: &ScalingConfigProperty{
//   		MaxEnvironments: jsii.Number(123),
//   	},
//   	State: jsii.String("state"),
//   	StateReason: jsii.String("stateReason"),
//   	ThrottleConfig: &ThrottleConfigProperty{
//   		RateLimit: jsii.Number(123),
//   	},
//   	UpdateStatus: jsii.String("updateStatus"),
//   	UpdateStatusReason: jsii.String("updateStatusReason"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionendpoint-regionalendpoint.html
//
type CfnWebFunctionEndpointPropsMixin_RegionalEndpointProperty struct {
	// The authentication type for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionendpoint-regionalendpoint.html#cfn-lambda-webfunctionendpoint-regionalendpoint-authtype
	//
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The domain name of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionendpoint-regionalendpoint.html#cfn-lambda-webfunctionendpoint-regionalendpoint-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The revision weights for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionendpoint-regionalendpoint.html#cfn-lambda-webfunctionendpoint-regionalendpoint-revisionweights
	//
	RevisionWeights interface{} `field:"optional" json:"revisionWeights" yaml:"revisionWeights"`
	// The scaling configuration for the endpoint.
	//
	// Optionally constrains how many concurrent execution environments the endpoint can use, in addition to your account's vCPU quota.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionendpoint-regionalendpoint.html#cfn-lambda-webfunctionendpoint-regionalendpoint-scalingconfig
	//
	ScalingConfig interface{} `field:"optional" json:"scalingConfig" yaml:"scalingConfig"`
	// The current state of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionendpoint-regionalendpoint.html#cfn-lambda-webfunctionendpoint-regionalendpoint-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// The reason for the current state of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionendpoint-regionalendpoint.html#cfn-lambda-webfunctionendpoint-regionalendpoint-statereason
	//
	StateReason *string `field:"optional" json:"stateReason" yaml:"stateReason"`
	// The throttling configuration for the endpoint.
	//
	// Optionally constrains the request rate that the endpoint accepts, in addition to your account's rate limit quota.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionendpoint-regionalendpoint.html#cfn-lambda-webfunctionendpoint-regionalendpoint-throttleconfig
	//
	ThrottleConfig interface{} `field:"optional" json:"throttleConfig" yaml:"throttleConfig"`
	// The status of the most recent update to the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionendpoint-regionalendpoint.html#cfn-lambda-webfunctionendpoint-regionalendpoint-updatestatus
	//
	UpdateStatus *string `field:"optional" json:"updateStatus" yaml:"updateStatus"`
	// The reason for the current update status of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionendpoint-regionalendpoint.html#cfn-lambda-webfunctionendpoint-regionalendpoint-updatestatusreason
	//
	UpdateStatusReason *string `field:"optional" json:"updateStatusReason" yaml:"updateStatusReason"`
}

