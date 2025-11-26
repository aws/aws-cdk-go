package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVPCEncryptionControl`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPCEncryptionControlProps := &CfnVPCEncryptionControlProps{
//   	EgressOnlyInternetGatewayExclusionInput: jsii.String("egressOnlyInternetGatewayExclusionInput"),
//   	ElasticFileSystemExclusionInput: jsii.String("elasticFileSystemExclusionInput"),
//   	InternetGatewayExclusionInput: jsii.String("internetGatewayExclusionInput"),
//   	LambdaExclusionInput: jsii.String("lambdaExclusionInput"),
//   	Mode: jsii.String("mode"),
//   	NatGatewayExclusionInput: jsii.String("natGatewayExclusionInput"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VirtualPrivateGatewayExclusionInput: jsii.String("virtualPrivateGatewayExclusionInput"),
//   	VpcId: jsii.String("vpcId"),
//   	VpcLatticeExclusionInput: jsii.String("vpcLatticeExclusionInput"),
//   	VpcPeeringExclusionInput: jsii.String("vpcPeeringExclusionInput"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html
//
type CfnVPCEncryptionControlProps struct {
	// Used to enable or disable EIGW exclusion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-egressonlyinternetgatewayexclusioninput
	//
	EgressOnlyInternetGatewayExclusionInput *string `field:"optional" json:"egressOnlyInternetGatewayExclusionInput" yaml:"egressOnlyInternetGatewayExclusionInput"`
	// Used to enable or disable EFS exclusion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-elasticfilesystemexclusioninput
	//
	ElasticFileSystemExclusionInput *string `field:"optional" json:"elasticFileSystemExclusionInput" yaml:"elasticFileSystemExclusionInput"`
	// Used to enable or disable IGW exclusion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-internetgatewayexclusioninput
	//
	InternetGatewayExclusionInput *string `field:"optional" json:"internetGatewayExclusionInput" yaml:"internetGatewayExclusionInput"`
	// Used to enable or disable Lambda exclusion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-lambdaexclusioninput
	//
	LambdaExclusionInput *string `field:"optional" json:"lambdaExclusionInput" yaml:"lambdaExclusionInput"`
	// The VPC encryption control mode, either monitor or enforce.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Used to enable or disable Nat gateway exclusion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-natgatewayexclusioninput
	//
	NatGatewayExclusionInput *string `field:"optional" json:"natGatewayExclusionInput" yaml:"natGatewayExclusionInput"`
	// The tags to assign to the VPC encryption control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Used to enable or disable VGW exclusion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-virtualprivategatewayexclusioninput
	//
	VirtualPrivateGatewayExclusionInput *string `field:"optional" json:"virtualPrivateGatewayExclusionInput" yaml:"virtualPrivateGatewayExclusionInput"`
	// The VPC on which this VPC encryption control is applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// Used to enable or disable Vpc Lattice exclusion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-vpclatticeexclusioninput
	//
	VpcLatticeExclusionInput *string `field:"optional" json:"vpcLatticeExclusionInput" yaml:"vpcLatticeExclusionInput"`
	// Used to enable or disable VPC peering exclusion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-vpcpeeringexclusioninput
	//
	VpcPeeringExclusionInput *string `field:"optional" json:"vpcPeeringExclusionInput" yaml:"vpcPeeringExclusionInput"`
}

