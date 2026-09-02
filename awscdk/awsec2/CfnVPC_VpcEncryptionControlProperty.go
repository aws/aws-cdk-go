package awsec2


// Describes the configuration and state of VPC encryption controls.
//
// For more information, see [Enforce VPC encryption in transit](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-encryption-controls.html) in the *Amazon VPC User Guide*.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcEncryptionControlProperty := &VpcEncryptionControlProperty{
//   	EgressOnlyInternetGatewayExclusion: jsii.String("egressOnlyInternetGatewayExclusion"),
//   	ElasticFileSystemExclusion: jsii.String("elasticFileSystemExclusion"),
//   	InternetGatewayExclusion: jsii.String("internetGatewayExclusion"),
//   	LambdaExclusion: jsii.String("lambdaExclusion"),
//   	Mode: jsii.String("mode"),
//   	NatGatewayExclusion: jsii.String("natGatewayExclusion"),
//   	ResourceExclusions: &VpcEncryptionControlExclusionsProperty{
//   		EgressOnlyInternetGateway: &VpcEncryptionControlExclusionProperty{
//   			State: jsii.String("state"),
//   			StateMessage: jsii.String("stateMessage"),
//   		},
//   		ElasticFileSystem: &VpcEncryptionControlExclusionProperty{
//   			State: jsii.String("state"),
//   			StateMessage: jsii.String("stateMessage"),
//   		},
//   		InternetGateway: &VpcEncryptionControlExclusionProperty{
//   			State: jsii.String("state"),
//   			StateMessage: jsii.String("stateMessage"),
//   		},
//   		Lambda: &VpcEncryptionControlExclusionProperty{
//   			State: jsii.String("state"),
//   			StateMessage: jsii.String("stateMessage"),
//   		},
//   		NatGateway: &VpcEncryptionControlExclusionProperty{
//   			State: jsii.String("state"),
//   			StateMessage: jsii.String("stateMessage"),
//   		},
//   		VirtualPrivateGateway: &VpcEncryptionControlExclusionProperty{
//   			State: jsii.String("state"),
//   			StateMessage: jsii.String("stateMessage"),
//   		},
//   		VpcLattice: &VpcEncryptionControlExclusionProperty{
//   			State: jsii.String("state"),
//   			StateMessage: jsii.String("stateMessage"),
//   		},
//   		VpcPeering: &VpcEncryptionControlExclusionProperty{
//   			State: jsii.String("state"),
//   			StateMessage: jsii.String("stateMessage"),
//   		},
//   	},
//   	State: jsii.String("state"),
//   	StateMessage: jsii.String("stateMessage"),
//   	VirtualPrivateGatewayExclusion: jsii.String("virtualPrivateGatewayExclusion"),
//   	VpcEncryptionControlId: jsii.String("vpcEncryptionControlId"),
//   	VpcId: jsii.String("vpcId"),
//   	VpcLatticeExclusion: jsii.String("vpcLatticeExclusion"),
//   	VpcPeeringExclusion: jsii.String("vpcPeeringExclusion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrol.html
//
type CfnVPC_VpcEncryptionControlProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrol.html#cfn-ec2-vpc-vpcencryptioncontrol-egressonlyinternetgatewayexclusion
	//
	EgressOnlyInternetGatewayExclusion *string `field:"optional" json:"egressOnlyInternetGatewayExclusion" yaml:"egressOnlyInternetGatewayExclusion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrol.html#cfn-ec2-vpc-vpcencryptioncontrol-elasticfilesystemexclusion
	//
	ElasticFileSystemExclusion *string `field:"optional" json:"elasticFileSystemExclusion" yaml:"elasticFileSystemExclusion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrol.html#cfn-ec2-vpc-vpcencryptioncontrol-internetgatewayexclusion
	//
	InternetGatewayExclusion *string `field:"optional" json:"internetGatewayExclusion" yaml:"internetGatewayExclusion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrol.html#cfn-ec2-vpc-vpcencryptioncontrol-lambdaexclusion
	//
	LambdaExclusion *string `field:"optional" json:"lambdaExclusion" yaml:"lambdaExclusion"`
	// The encryption mode for the VPC Encryption Control configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrol.html#cfn-ec2-vpc-vpcencryptioncontrol-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrol.html#cfn-ec2-vpc-vpcencryptioncontrol-natgatewayexclusion
	//
	NatGatewayExclusion *string `field:"optional" json:"natGatewayExclusion" yaml:"natGatewayExclusion"`
	// Describes the exclusion configurations for various resource types in VPC Encryption Control.
	//
	// For more information, see [Enforce VPC encryption in transit](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-encryption-controls.html) in the *Amazon VPC User Guide*.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrol.html#cfn-ec2-vpc-vpcencryptioncontrol-resourceexclusions
	//
	ResourceExclusions interface{} `field:"optional" json:"resourceExclusions" yaml:"resourceExclusions"`
	// The current state of the VPC Encryption Control configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrol.html#cfn-ec2-vpc-vpcencryptioncontrol-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// A message providing additional information about the encryption control state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrol.html#cfn-ec2-vpc-vpcencryptioncontrol-statemessage
	//
	StateMessage *string `field:"optional" json:"stateMessage" yaml:"stateMessage"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrol.html#cfn-ec2-vpc-vpcencryptioncontrol-virtualprivategatewayexclusion
	//
	VirtualPrivateGatewayExclusion *string `field:"optional" json:"virtualPrivateGatewayExclusion" yaml:"virtualPrivateGatewayExclusion"`
	// The ID of the VPC Encryption Control configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrol.html#cfn-ec2-vpc-vpcencryptioncontrol-vpcencryptioncontrolid
	//
	VpcEncryptionControlId *string `field:"optional" json:"vpcEncryptionControlId" yaml:"vpcEncryptionControlId"`
	// The ID of the VPC associated with the encryption control configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrol.html#cfn-ec2-vpc-vpcencryptioncontrol-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrol.html#cfn-ec2-vpc-vpcencryptioncontrol-vpclatticeexclusion
	//
	VpcLatticeExclusion *string `field:"optional" json:"vpcLatticeExclusion" yaml:"vpcLatticeExclusion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrol.html#cfn-ec2-vpc-vpcencryptioncontrol-vpcpeeringexclusion
	//
	VpcPeeringExclusion *string `field:"optional" json:"vpcPeeringExclusion" yaml:"vpcPeeringExclusion"`
}

