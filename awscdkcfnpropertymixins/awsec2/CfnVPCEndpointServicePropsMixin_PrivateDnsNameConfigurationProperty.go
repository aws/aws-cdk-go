package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   privateDnsNameConfigurationProperty := &PrivateDnsNameConfigurationProperty{
//   	Name: jsii.String("name"),
//   	State: jsii.String("state"),
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcendpointservice-privatednsnameconfiguration.html
//
type CfnVPCEndpointServicePropsMixin_PrivateDnsNameConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcendpointservice-privatednsnameconfiguration.html#cfn-ec2-vpcendpointservice-privatednsnameconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcendpointservice-privatednsnameconfiguration.html#cfn-ec2-vpcendpointservice-privatednsnameconfiguration-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcendpointservice-privatednsnameconfiguration.html#cfn-ec2-vpcendpointservice-privatednsnameconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcendpointservice-privatednsnameconfiguration.html#cfn-ec2-vpcendpointservice-privatednsnameconfiguration-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

