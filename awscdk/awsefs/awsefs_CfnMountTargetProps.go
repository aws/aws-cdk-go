package awsefs


// Properties for defining a `CfnMountTarget`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMountTargetProps := &cfnMountTargetProps{
//   	fileSystemId: jsii.String("fileSystemId"),
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	subnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	ipAddress: jsii.String("ipAddress"),
//   }
//
type CfnMountTargetProps struct {
	// The ID of the file system for which to create the mount target.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Up to five VPC security group IDs, of the form `sg-xxxxxxxx` .
	//
	// These must be for the same VPC as subnet specified.
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The ID of the subnet to add the mount target in.
	//
	// For file systems that use One Zone storage classes, use the subnet that is associated with the file system's Availability Zone.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Valid IPv4 address within the address range of the specified subnet.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}

