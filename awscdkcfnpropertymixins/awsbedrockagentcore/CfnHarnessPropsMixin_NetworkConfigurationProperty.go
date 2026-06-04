package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   networkConfigurationProperty := &NetworkConfigurationProperty{
//   	NetworkMode: jsii.String("networkMode"),
//   	NetworkModeConfig: &VpcConfigProperty{
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-networkconfiguration.html
//
type CfnHarnessPropsMixin_NetworkConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-networkconfiguration.html#cfn-bedrockagentcore-harness-networkconfiguration-networkmode
	//
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-networkconfiguration.html#cfn-bedrockagentcore-harness-networkconfiguration-networkmodeconfig
	//
	NetworkModeConfig interface{} `field:"optional" json:"networkModeConfig" yaml:"networkModeConfig"`
}

