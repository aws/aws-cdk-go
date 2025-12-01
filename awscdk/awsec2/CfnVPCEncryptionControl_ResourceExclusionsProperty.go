package awsec2


// Information about resource exclusions for the VPC Encryption Control configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceExclusionsProperty := &ResourceExclusionsProperty{
//   	EgressOnlyInternetGateway: &VpcEncryptionControlExclusionProperty{
//   		State: jsii.String("state"),
//   		StateMessage: jsii.String("stateMessage"),
//   	},
//   	ElasticFileSystem: &VpcEncryptionControlExclusionProperty{
//   		State: jsii.String("state"),
//   		StateMessage: jsii.String("stateMessage"),
//   	},
//   	InternetGateway: &VpcEncryptionControlExclusionProperty{
//   		State: jsii.String("state"),
//   		StateMessage: jsii.String("stateMessage"),
//   	},
//   	Lambda: &VpcEncryptionControlExclusionProperty{
//   		State: jsii.String("state"),
//   		StateMessage: jsii.String("stateMessage"),
//   	},
//   	NatGateway: &VpcEncryptionControlExclusionProperty{
//   		State: jsii.String("state"),
//   		StateMessage: jsii.String("stateMessage"),
//   	},
//   	VirtualPrivateGateway: &VpcEncryptionControlExclusionProperty{
//   		State: jsii.String("state"),
//   		StateMessage: jsii.String("stateMessage"),
//   	},
//   	VpcLattice: &VpcEncryptionControlExclusionProperty{
//   		State: jsii.String("state"),
//   		StateMessage: jsii.String("stateMessage"),
//   	},
//   	VpcPeering: &VpcEncryptionControlExclusionProperty{
//   		State: jsii.String("state"),
//   		StateMessage: jsii.String("stateMessage"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-resourceexclusions.html
//
type CfnVPCEncryptionControl_ResourceExclusionsProperty struct {
	// Specifies whether to exclude egress-only internet gateway traffic from encryption enforcement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-resourceexclusions.html#cfn-ec2-vpcencryptioncontrol-resourceexclusions-egressonlyinternetgateway
	//
	EgressOnlyInternetGateway interface{} `field:"optional" json:"egressOnlyInternetGateway" yaml:"egressOnlyInternetGateway"`
	// Specifies whether to exclude Elastic File System traffic from encryption enforcement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-resourceexclusions.html#cfn-ec2-vpcencryptioncontrol-resourceexclusions-elasticfilesystem
	//
	ElasticFileSystem interface{} `field:"optional" json:"elasticFileSystem" yaml:"elasticFileSystem"`
	// Specifies whether to exclude internet gateway traffic from encryption enforcement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-resourceexclusions.html#cfn-ec2-vpcencryptioncontrol-resourceexclusions-internetgateway
	//
	InternetGateway interface{} `field:"optional" json:"internetGateway" yaml:"internetGateway"`
	// Specifies whether to exclude Lambda function traffic from encryption enforcement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-resourceexclusions.html#cfn-ec2-vpcencryptioncontrol-resourceexclusions-lambda
	//
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// Specifies whether to exclude NAT gateway traffic from encryption enforcement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-resourceexclusions.html#cfn-ec2-vpcencryptioncontrol-resourceexclusions-natgateway
	//
	NatGateway interface{} `field:"optional" json:"natGateway" yaml:"natGateway"`
	// Specifies whether to exclude virtual private gateway traffic from encryption enforcement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-resourceexclusions.html#cfn-ec2-vpcencryptioncontrol-resourceexclusions-virtualprivategateway
	//
	VirtualPrivateGateway interface{} `field:"optional" json:"virtualPrivateGateway" yaml:"virtualPrivateGateway"`
	// Specifies whether to exclude VPC Lattice traffic from encryption enforcement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-resourceexclusions.html#cfn-ec2-vpcencryptioncontrol-resourceexclusions-vpclattice
	//
	VpcLattice interface{} `field:"optional" json:"vpcLattice" yaml:"vpcLattice"`
	// Specifies whether to exclude VPC peering connection traffic from encryption enforcement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-resourceexclusions.html#cfn-ec2-vpcencryptioncontrol-resourceexclusions-vpcpeering
	//
	VpcPeering interface{} `field:"optional" json:"vpcPeering" yaml:"vpcPeering"`
}

