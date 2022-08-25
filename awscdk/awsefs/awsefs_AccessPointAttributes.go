package awsefs


// Attributes that can be specified when importing an AccessPoint.
//
// Example:
//   efs.accessPoint.fromAccessPointAttributes(this, jsii.String("ap"), &accessPointAttributes{
//   	accessPointId: jsii.String("fsap-1293c4d9832fo0912"),
//   	fileSystem: efs.fileSystem.fromFileSystemAttributes(this, jsii.String("efs"), &fileSystemAttributes{
//   		fileSystemId: jsii.String("fs-099d3e2f"),
//   		securityGroup: ec2.securityGroup.fromSecurityGroupId(this, jsii.String("sg"), jsii.String("sg-51530134")),
//   	}),
//   })
//
type AccessPointAttributes struct {
	// The ARN of the AccessPoint One of this, or {@link accessPointId} is required.
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// The ID of the AccessPoint One of this, or {@link accessPointArn} is required.
	AccessPointId *string `field:"optional" json:"accessPointId" yaml:"accessPointId"`
	// The EFS file system.
	FileSystem IFileSystem `field:"optional" json:"fileSystem" yaml:"fileSystem"`
}

