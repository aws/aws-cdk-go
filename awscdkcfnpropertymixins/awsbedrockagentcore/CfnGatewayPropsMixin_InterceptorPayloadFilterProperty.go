package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   interceptorPayloadFilterProperty := &InterceptorPayloadFilterProperty{
//   	Exclude: []interface{}{
//   		&InterceptorPayloadExclusionSelectorProperty{
//   			Field: jsii.String("field"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-interceptorpayloadfilter.html
//
type CfnGatewayPropsMixin_InterceptorPayloadFilterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-interceptorpayloadfilter.html#cfn-bedrockagentcore-gateway-interceptorpayloadfilter-exclude
	//
	Exclude interface{} `field:"optional" json:"exclude" yaml:"exclude"`
}

