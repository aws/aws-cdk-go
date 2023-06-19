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
// Experimental.
type AccessPointAttributes struct {
	// The ARN of the AccessPoint One of this, or {@link accessPointId} is required.
	// Experimental.
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// The ID of the AccessPoint One of this, or {@link accessPointArn} is required.
	// Experimental.
	AccessPointId *string `field:"optional" json:"accessPointId" yaml:"accessPointId"`
	// The EFS file system.
	// Experimental.
	FileSystem IFileSystem `field:"optional" json:"fileSystem" yaml:"fileSystem"`
}

