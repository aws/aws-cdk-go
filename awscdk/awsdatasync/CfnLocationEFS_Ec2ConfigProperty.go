package awsdatasync


// The subnet and security groups that AWS DataSync uses to access your Amazon EFS file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ec2ConfigProperty := &Ec2ConfigProperty{
//   	SecurityGroupArns: []*string{
//   		jsii.String("securityGroupArns"),
//   	},
//   	SubnetArn: jsii.String("subnetArn"),
//   }
//
type CfnLocationEFS_Ec2ConfigProperty struct {
	// Specifies the Amazon Resource Names (ARNs) of the security groups associated with an Amazon EFS file system's mount target.
	SecurityGroupArns *[]*string `field:"required" json:"securityGroupArns" yaml:"securityGroupArns"`
	// Specifies the ARN of a subnet where DataSync creates the [network interfaces](https://docs.aws.amazon.com/datasync/latest/userguide/datasync-network.html#required-network-interfaces) for managing traffic during your transfer.
	//
	// The subnet must be located:
	//
	// - In the same virtual private cloud (VPC) as the Amazon EFS file system.
	// - In the same Availability Zone as at least one mount target for the Amazon EFS file system.
	//
	// > You don't need to specify a subnet that includes a file system mount target.
	SubnetArn *string `field:"required" json:"subnetArn" yaml:"subnetArn"`
}

