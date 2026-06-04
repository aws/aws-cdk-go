package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimeTargetConfigurationProperty := &RuntimeTargetConfigurationProperty{
//   	Arn: jsii.String("arn"),
//
//   	// the properties below are optional
//   	Qualifier: jsii.String("qualifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-runtimetargetconfiguration.html
//
type CfnGatewayTarget_RuntimeTargetConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-runtimetargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-runtimetargetconfiguration-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-runtimetargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-runtimetargetconfiguration-qualifier
	//
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

