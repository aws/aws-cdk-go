package previewawsec2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnVPCEncryptionControlPropsMixin_ResourceExclusionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-resourceexclusions.html#cfn-ec2-vpcencryptioncontrol-resourceexclusions-egressonlyinternetgateway
	//
	EgressOnlyInternetGateway interface{} `field:"optional" json:"egressOnlyInternetGateway" yaml:"egressOnlyInternetGateway"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-resourceexclusions.html#cfn-ec2-vpcencryptioncontrol-resourceexclusions-elasticfilesystem
	//
	ElasticFileSystem interface{} `field:"optional" json:"elasticFileSystem" yaml:"elasticFileSystem"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-resourceexclusions.html#cfn-ec2-vpcencryptioncontrol-resourceexclusions-internetgateway
	//
	InternetGateway interface{} `field:"optional" json:"internetGateway" yaml:"internetGateway"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-resourceexclusions.html#cfn-ec2-vpcencryptioncontrol-resourceexclusions-lambda
	//
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-resourceexclusions.html#cfn-ec2-vpcencryptioncontrol-resourceexclusions-natgateway
	//
	NatGateway interface{} `field:"optional" json:"natGateway" yaml:"natGateway"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-resourceexclusions.html#cfn-ec2-vpcencryptioncontrol-resourceexclusions-virtualprivategateway
	//
	VirtualPrivateGateway interface{} `field:"optional" json:"virtualPrivateGateway" yaml:"virtualPrivateGateway"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-resourceexclusions.html#cfn-ec2-vpcencryptioncontrol-resourceexclusions-vpclattice
	//
	VpcLattice interface{} `field:"optional" json:"vpcLattice" yaml:"vpcLattice"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-resourceexclusions.html#cfn-ec2-vpcencryptioncontrol-resourceexclusions-vpcpeering
	//
	VpcPeering interface{} `field:"optional" json:"vpcPeering" yaml:"vpcPeering"`
}

