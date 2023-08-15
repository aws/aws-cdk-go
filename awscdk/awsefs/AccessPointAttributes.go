package awsefs


// Attributes that can be specified when importing an AccessPoint.
//
// Example:
//   efs.AccessPoint_FromAccessPointAttributes(this, jsii.String("ap"), &AccessPointAttributes{
//   	AccessPointId: jsii.String("fsap-1293c4d9832fo0912"),
//   	FileSystem: efs.FileSystem_FromFileSystemAttributes(this, jsii.String("efs"), &FileSystemAttributes{
//   		FileSystemId: jsii.String("fs-099d3e2f"),
//   		SecurityGroup: ec2.SecurityGroup_FromSecurityGroupId(this, jsii.String("sg"), jsii.String("sg-51530134")),
//   	}),
//   })
//
type AccessPointAttributes struct {
	// The ARN of the AccessPoint One of this, or `accessPointId` is required.
	// Default: - determined based on accessPointId.
	//
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// The ID of the AccessPoint One of this, or `accessPointArn` is required.
	// Default: - determined based on accessPointArn.
	//
	AccessPointId *string `field:"optional" json:"accessPointId" yaml:"accessPointId"`
	// The EFS file system.
	// Default: - no EFS file system.
	//
	FileSystem IFileSystem `field:"optional" json:"fileSystem" yaml:"fileSystem"`
}

