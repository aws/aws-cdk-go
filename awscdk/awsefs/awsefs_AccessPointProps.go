package awsefs


// Properties for the AccessPoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fileSystem fileSystem
//
//   accessPointProps := &accessPointProps{
//   	fileSystem: fileSystem,
//
//   	// the properties below are optional
//   	createAcl: &acl{
//   		ownerGid: jsii.String("ownerGid"),
//   		ownerUid: jsii.String("ownerUid"),
//   		permissions: jsii.String("permissions"),
//   	},
//   	path: jsii.String("path"),
//   	posixUser: &posixUser{
//   		gid: jsii.String("gid"),
//   		uid: jsii.String("uid"),
//
//   		// the properties below are optional
//   		secondaryGids: []*string{
//   			jsii.String("secondaryGids"),
//   		},
//   	},
//   }
//
type AccessPointProps struct {
	// Specifies the POSIX IDs and permissions to apply when creating the access point's root directory.
	//
	// If the
	// root directory specified by `path` does not exist, EFS creates the root directory and applies the
	// permissions specified here. If the specified `path` does not exist, you must specify `createAcl`.
	CreateAcl *Acl `field:"optional" json:"createAcl" yaml:"createAcl"`
	// Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The full POSIX identity, including the user ID, group ID, and any secondary group IDs, on the access point that is used for all file system operations performed by NFS clients using the access point.
	//
	// Specify this to enforce a user identity using an access point.
	// See: - [Enforcing a User Identity Using an Access Point](https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html)
	//
	PosixUser *PosixUser `field:"optional" json:"posixUser" yaml:"posixUser"`
	// The efs filesystem.
	FileSystem IFileSystem `field:"required" json:"fileSystem" yaml:"fileSystem"`
}

