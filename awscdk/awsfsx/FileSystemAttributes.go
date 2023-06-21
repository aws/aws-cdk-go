package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties that describe an existing FSx file system.
//
// Example:
//   sg := ec2.SecurityGroup_FromSecurityGroupId(this, jsii.String("FsxSecurityGroup"), jsii.String("{SECURITY-GROUP-ID}"))
//   fs := fsx.LustreFileSystem_FromLustreFileSystemAttributes(this, jsii.String("FsxLustreFileSystem"), &FileSystemAttributes{
//   	DnsName: jsii.String("{FILE-SYSTEM-DNS-NAME}"),
//   	FileSystemId: jsii.String("{FILE-SYSTEM-ID}"),
//   	SecurityGroup: sg,
//   })
//
//   vpc := ec2.Vpc_FromVpcAttributes(this, jsii.String("Vpc"), &VpcAttributes{
//   	AvailabilityZones: []*string{
//   		jsii.String("us-west-2a"),
//   		jsii.String("us-west-2b"),
//   	},
//   	PublicSubnetIds: []*string{
//   		jsii.String("{US-WEST-2A-SUBNET-ID}"),
//   		jsii.String("{US-WEST-2B-SUBNET-ID}"),
//   	},
//   	VpcId: jsii.String("{VPC-ID}"),
//   })
//
//   inst := ec2.NewInstance(this, jsii.String("inst"), &InstanceProps{
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_T2, ec2.InstanceSize_LARGE),
//   	MachineImage: ec2.NewAmazonLinuxImage(&AmazonLinuxImageProps{
//   		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
//   	}),
//   	Vpc: Vpc,
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	},
//   })
//
//   fs.Connections.AllowDefaultPortFrom(inst)
//
type FileSystemAttributes struct {
	// The DNS name assigned to this file system.
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// The ID of the file system, assigned by Amazon FSx.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// The security group of the file system.
	SecurityGroup awsec2.ISecurityGroup `field:"required" json:"securityGroup" yaml:"securityGroup"`
}

