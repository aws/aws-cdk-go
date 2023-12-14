package awsefs


// Specifies the directory on the Amazon EFS file system that the access point provides access to.
//
// The access point exposes the specified file system path as the root directory of your file system to applications using the access point. NFS clients using the access point can only access data in the access point's `RootDirectory` and its subdirectories.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rootDirectoryProperty := &RootDirectoryProperty{
//   	CreationInfo: &CreationInfoProperty{
//   		OwnerGid: jsii.String("ownerGid"),
//   		OwnerUid: jsii.String("ownerUid"),
//   		Permissions: jsii.String("permissions"),
//   	},
//   	Path: jsii.String("path"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-rootdirectory.html
//
type CfnAccessPoint_RootDirectoryProperty struct {
	// (Optional) Specifies the POSIX IDs and permissions to apply to the access point's `RootDirectory` .
	//
	// If the `RootDirectory` > `Path` specified does not exist, EFS creates the root directory using the `CreationInfo` settings when a client connects to an access point. When specifying the `CreationInfo` , you must provide values for all properties.
	//
	// > If you do not provide `CreationInfo` and the specified `RootDirectory` > `Path` does not exist, attempts to mount the file system using the access point will fail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-rootdirectory.html#cfn-efs-accesspoint-rootdirectory-creationinfo
	//
	CreationInfo interface{} `field:"optional" json:"creationInfo" yaml:"creationInfo"`
	// Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system.
	//
	// A path can have up to four subdirectories. If the specified path does not exist, you are required to provide the `CreationInfo` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-rootdirectory.html#cfn-efs-accesspoint-rootdirectory-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
}

