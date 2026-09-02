package awsec2


// Describes the exclusion configurations for various resource types in VPC Encryption Control.
//
// For more information, see [Enforce VPC encryption in transit](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-encryption-controls.html) in the *Amazon VPC User Guide*.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   vpcEncryptionControlExclusionsProperty := &VpcEncryptionControlExclusionsProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrolexclusions.html
//
type CfnVPCPropsMixin_VpcEncryptionControlExclusionsProperty struct {
	// Describes an exclusion configuration for VPC Encryption Control.
	//
	// For more information, see [Enforce VPC encryption in transit](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-encryption-controls.html) in the *Amazon VPC User Guide*.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrolexclusions.html#cfn-ec2-vpc-vpcencryptioncontrolexclusions-egressonlyinternetgateway
	//
	EgressOnlyInternetGateway interface{} `field:"optional" json:"egressOnlyInternetGateway" yaml:"egressOnlyInternetGateway"`
	// Describes an exclusion configuration for VPC Encryption Control.
	//
	// For more information, see [Enforce VPC encryption in transit](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-encryption-controls.html) in the *Amazon VPC User Guide*.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrolexclusions.html#cfn-ec2-vpc-vpcencryptioncontrolexclusions-elasticfilesystem
	//
	ElasticFileSystem interface{} `field:"optional" json:"elasticFileSystem" yaml:"elasticFileSystem"`
	// Describes an exclusion configuration for VPC Encryption Control.
	//
	// For more information, see [Enforce VPC encryption in transit](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-encryption-controls.html) in the *Amazon VPC User Guide*.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrolexclusions.html#cfn-ec2-vpc-vpcencryptioncontrolexclusions-internetgateway
	//
	InternetGateway interface{} `field:"optional" json:"internetGateway" yaml:"internetGateway"`
	// Describes an exclusion configuration for VPC Encryption Control.
	//
	// For more information, see [Enforce VPC encryption in transit](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-encryption-controls.html) in the *Amazon VPC User Guide*.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrolexclusions.html#cfn-ec2-vpc-vpcencryptioncontrolexclusions-lambda
	//
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// Describes an exclusion configuration for VPC Encryption Control.
	//
	// For more information, see [Enforce VPC encryption in transit](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-encryption-controls.html) in the *Amazon VPC User Guide*.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrolexclusions.html#cfn-ec2-vpc-vpcencryptioncontrolexclusions-natgateway
	//
	NatGateway interface{} `field:"optional" json:"natGateway" yaml:"natGateway"`
	// Describes an exclusion configuration for VPC Encryption Control.
	//
	// For more information, see [Enforce VPC encryption in transit](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-encryption-controls.html) in the *Amazon VPC User Guide*.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrolexclusions.html#cfn-ec2-vpc-vpcencryptioncontrolexclusions-virtualprivategateway
	//
	VirtualPrivateGateway interface{} `field:"optional" json:"virtualPrivateGateway" yaml:"virtualPrivateGateway"`
	// Describes an exclusion configuration for VPC Encryption Control.
	//
	// For more information, see [Enforce VPC encryption in transit](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-encryption-controls.html) in the *Amazon VPC User Guide*.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrolexclusions.html#cfn-ec2-vpc-vpcencryptioncontrolexclusions-vpclattice
	//
	VpcLattice interface{} `field:"optional" json:"vpcLattice" yaml:"vpcLattice"`
	// Describes an exclusion configuration for VPC Encryption Control.
	//
	// For more information, see [Enforce VPC encryption in transit](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-encryption-controls.html) in the *Amazon VPC User Guide*.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrolexclusions.html#cfn-ec2-vpc-vpcencryptioncontrolexclusions-vpcpeering
	//
	VpcPeering interface{} `field:"optional" json:"vpcPeering" yaml:"vpcPeering"`
}

