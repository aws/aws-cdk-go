package awsefs


// Options to create an AccessPoint.
//
// Example:
//   fileSystem.AddAccessPoint(jsii.String("MyAccessPoint"), &AccessPointOptions{
//   	// create a unique access point via an optional client token
//   	ClientToken: jsii.String("client-token"),
//   })
//
type AccessPointOptions struct {
	// The opaque string specified in the request to ensure idempotent creation.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-accesspoint.html#cfn-efs-accesspoint-clienttoken
	//
	// Default: - No client token.
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
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
}

