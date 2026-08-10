package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelMappingProperty := &ModelMappingProperty{
//   	ProviderPrefix: &ProviderPrefixProperty{
//   		Separator: jsii.String("separator"),
//   		Strip: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-modelmapping.html
//
type CfnGatewayTarget_ModelMappingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-modelmapping.html#cfn-bedrockagentcore-gatewaytarget-modelmapping-providerprefix
	//
	ProviderPrefix interface{} `field:"optional" json:"providerPrefix" yaml:"providerPrefix"`
}

