package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   runtimeTargetConfigurationProperty := &RuntimeTargetConfigurationProperty{
//   	Arn: jsii.String("arn"),
//   	Qualifier: jsii.String("qualifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-runtimetargetconfiguration.html
//
type CfnGatewayTargetPropsMixin_RuntimeTargetConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-runtimetargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-runtimetargetconfiguration-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-runtimetargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-runtimetargetconfiguration-qualifier
	//
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

