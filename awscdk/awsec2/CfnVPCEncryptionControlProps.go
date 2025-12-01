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
	// Specifies whether to exclude egress-only internet gateway traffic from encryption enforcement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-egressonlyinternetgatewayexclusioninput
	//
	EgressOnlyInternetGatewayExclusionInput *string `field:"optional" json:"egressOnlyInternetGatewayExclusionInput" yaml:"egressOnlyInternetGatewayExclusionInput"`
	// Specifies whether to exclude Elastic File System traffic from encryption enforcement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-elasticfilesystemexclusioninput
	//
	ElasticFileSystemExclusionInput *string `field:"optional" json:"elasticFileSystemExclusionInput" yaml:"elasticFileSystemExclusionInput"`
	// Specifies whether to exclude internet gateway traffic from encryption enforcement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-internetgatewayexclusioninput
	//
	InternetGatewayExclusionInput *string `field:"optional" json:"internetGatewayExclusionInput" yaml:"internetGatewayExclusionInput"`
	// Specifies whether to exclude Lambda function traffic from encryption enforcement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-lambdaexclusioninput
	//
	LambdaExclusionInput *string `field:"optional" json:"lambdaExclusionInput" yaml:"lambdaExclusionInput"`
	// The encryption mode for the VPC Encryption Control configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Specifies whether to exclude NAT gateway traffic from encryption enforcement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-natgatewayexclusioninput
	//
	NatGatewayExclusionInput *string `field:"optional" json:"natGatewayExclusionInput" yaml:"natGatewayExclusionInput"`
	// The tags assigned to the VPC Encryption Control configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies whether to exclude virtual private gateway traffic from encryption enforcement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-virtualprivategatewayexclusioninput
	//
	VirtualPrivateGatewayExclusionInput *string `field:"optional" json:"virtualPrivateGatewayExclusionInput" yaml:"virtualPrivateGatewayExclusionInput"`
	// The ID of the VPC for which to create the encryption control configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// Specifies whether to exclude VPC Lattice traffic from encryption enforcement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-vpclatticeexclusioninput
	//
	VpcLatticeExclusionInput *string `field:"optional" json:"vpcLatticeExclusionInput" yaml:"vpcLatticeExclusionInput"`
	// Specifies whether to exclude VPC peering connection traffic from encryption enforcement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcencryptioncontrol.html#cfn-ec2-vpcencryptioncontrol-vpcpeeringexclusioninput
	//
	VpcPeeringExclusionInput *string `field:"optional" json:"vpcPeeringExclusionInput" yaml:"vpcPeeringExclusionInput"`
}

