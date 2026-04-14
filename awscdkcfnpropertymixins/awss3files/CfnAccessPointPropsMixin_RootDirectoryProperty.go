package awss3files


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   rootDirectoryProperty := &RootDirectoryProperty{
//   	CreationPermissions: &CreationPermissionsProperty{
//   		OwnerGid: jsii.String("ownerGid"),
//   		OwnerUid: jsii.String("ownerUid"),
//   		Permissions: jsii.String("permissions"),
//   	},
//   	Path: jsii.String("path"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-accesspoint-rootdirectory.html
//
type CfnAccessPointPropsMixin_RootDirectoryProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-accesspoint-rootdirectory.html#cfn-s3files-accesspoint-rootdirectory-creationpermissions
	//
	CreationPermissions interface{} `field:"optional" json:"creationPermissions" yaml:"creationPermissions"`
	// Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system.
	//
	// A path can have up to four subdirectories. If the specified path does not exist, you are required to provide the CreationPermissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-accesspoint-rootdirectory.html#cfn-s3files-accesspoint-rootdirectory-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
}

