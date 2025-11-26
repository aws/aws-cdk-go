package previewawsdatasyncmixins


// The subnet and security groups that AWS DataSync uses to connect to one of your Amazon EFS file system's [mount targets](https://docs.aws.amazon.com/efs/latest/ug/accessing-fs.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ec2ConfigProperty := &Ec2ConfigProperty{
//   	SecurityGroupArns: []*string{
//   		jsii.String("securityGroupArns"),
//   	},
//   	SubnetArn: jsii.String("subnetArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationefs-ec2config.html
//
type CfnLocationEFSPropsMixin_Ec2ConfigProperty struct {
	// Specifies the Amazon Resource Names (ARNs) of the security groups associated with an Amazon EFS file system's mount target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationefs-ec2config.html#cfn-datasync-locationefs-ec2config-securitygrouparns
	//
	SecurityGroupArns *[]*string `field:"optional" json:"securityGroupArns" yaml:"securityGroupArns"`
	// Specifies the ARN of a subnet where DataSync creates the [network interfaces](https://docs.aws.amazon.com/datasync/latest/userguide/datasync-network.html#required-network-interfaces.html) for managing traffic during your transfer.
	//
	// The subnet must be located:
	//
	// - In the same virtual private cloud (VPC) as the Amazon EFS file system.
	// - In the same Availability Zone as at least one mount target for the Amazon EFS file system.
	//
	// > You don't need to specify a subnet that includes a file system mount target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationefs-ec2config.html#cfn-datasync-locationefs-ec2config-subnetarn
	//
	SubnetArn *string `field:"optional" json:"subnetArn" yaml:"subnetArn"`
}

