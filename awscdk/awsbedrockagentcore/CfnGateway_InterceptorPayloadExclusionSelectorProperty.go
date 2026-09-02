package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   interceptorPayloadExclusionSelectorProperty := &InterceptorPayloadExclusionSelectorProperty{
//   	Field: jsii.String("field"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-interceptorpayloadexclusionselector.html
//
type CfnGateway_InterceptorPayloadExclusionSelectorProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-interceptorpayloadexclusionselector.html#cfn-bedrockagentcore-gateway-interceptorpayloadexclusionselector-field
	//
	Field *string `field:"required" json:"field" yaml:"field"`
}

