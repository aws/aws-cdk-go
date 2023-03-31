package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Properties that describe an existing FSx file system.
//
// Example:
//   sg := ec2.securityGroup.fromSecurityGroupId(this, jsii.String("FsxSecurityGroup"), jsii.String("{SECURITY-GROUP-ID}"))
//   fs := fsx.lustreFileSystem.fromLustreFileSystemAttributes(this, jsii.String("FsxLustreFileSystem"), &fileSystemAttributes{
//   	dnsName: jsii.String("{FILE-SYSTEM-DNS-NAME}"),
//   	fileSystemId: jsii.String("{FILE-SYSTEM-ID}"),
//   	securityGroup: sg,
//   })
//
//   vpc := ec2.vpc.fromVpcAttributes(this, jsii.String("Vpc"), &vpcAttributes{
//   	availabilityZones: []*string{
//   		jsii.String("us-west-2a"),
//   		jsii.String("us-west-2b"),
//   	},
//   	publicSubnetIds: []*string{
//   		jsii.String("{US-WEST-2A-SUBNET-ID}"),
//   		jsii.String("{US-WEST-2B-SUBNET-ID}"),
//   	},
//   	vpcId: jsii.String("{VPC-ID}"),
//   })
//
//   inst := ec2.NewInstance(this, jsii.String("inst"), &instanceProps{
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_T2, ec2.instanceSize_LARGE),
//   	machineImage: ec2.NewAmazonLinuxImage(&amazonLinuxImageProps{
//   		generation: ec2.amazonLinuxGeneration_AMAZON_LINUX_2,
//   	}),
//   	vpc: vpc,
//   	vpcSubnets: &subnetSelection{
//   		subnetType: ec2.subnetType_PUBLIC,
//   	},
//   })
//
//   fs.connections.allowDefaultPortFrom(inst)
//
// Experimental.
type FileSystemAttributes struct {
	// The DNS name assigned to this file system.
	// Experimental.
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// The ID of the file system, assigned by Amazon FSx.
	// Experimental.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// The security group of the file system.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"required" json:"securityGroup" yaml:"securityGroup"`
}

