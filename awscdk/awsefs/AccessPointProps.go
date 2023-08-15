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
//   accessPointProps := &AccessPointProps{
//   	FileSystem: fileSystem,
//
//   	// the properties below are optional
//   	CreateAcl: &Acl{
//   		OwnerGid: jsii.String("ownerGid"),
//   		OwnerUid: jsii.String("ownerUid"),
//   		Permissions: jsii.String("permissions"),
//   	},
//   	Path: jsii.String("path"),
//   	PosixUser: &PosixUser{
//   		Gid: jsii.String("gid"),
//   		Uid: jsii.String("uid"),
//
//   		// the properties below are optional
//   		SecondaryGids: []*string{
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
	// Default: - None. The directory specified by `path` must exist.
	//
	CreateAcl *Acl `field:"optional" json:"createAcl" yaml:"createAcl"`
	// Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system.
	// Default: '/'.
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The full POSIX identity, including the user ID, group ID, and any secondary group IDs, on the access point that is used for all file system operations performed by NFS clients using the access point.
	//
	// Specify this to enforce a user identity using an access point.
	// See:  - [Enforcing a User Identity Using an Access Point](https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html)
	//
	// Default: - user identity not enforced.
	//
	PosixUser *PosixUser `field:"optional" json:"posixUser" yaml:"posixUser"`
	// The efs filesystem.
	FileSystem IFileSystem `field:"required" json:"fileSystem" yaml:"fileSystem"`
}

