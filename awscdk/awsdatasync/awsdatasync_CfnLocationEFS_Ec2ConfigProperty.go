package awsdatasync


// The subnet and the security group that DataSync uses to access the target EFS file system.
//
// The subnet must have at least one mount target for that file system. The security group that you provide must be able to communicate with the security group on the mount target in the subnet specified.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ec2ConfigProperty := &ec2ConfigProperty{
//   	securityGroupArns: []*string{
//   		jsii.String("securityGroupArns"),
//   	},
//   	subnetArn: jsii.String("subnetArn"),
//   }
//
type CfnLocationEFS_Ec2ConfigProperty struct {
	// The Amazon Resource Names (ARNs) of the security groups that are configured for the Amazon EC2 resource.
	//
	// *Pattern* : `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:security-group/.*$`
	SecurityGroupArns *[]*string `field:"required" json:"securityGroupArns" yaml:"securityGroupArns"`
	// The Amazon Resource Name (ARN) of the subnet that DataSync uses to access the target EFS file system.
	SubnetArn *string `field:"required" json:"subnetArn" yaml:"subnetArn"`
}

