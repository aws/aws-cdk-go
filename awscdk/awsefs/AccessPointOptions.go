package awsefs


// Options to create an AccessPoint.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import efs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   // create a new VPC
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//
//   // create a new Amazon EFS filesystem
//   fileSystem := efs.NewFileSystem(this, jsii.String("Efs"), &FileSystemProps{
//   	Vpc: Vpc,
//   })
//
//   // create a new access point from the filesystem
//   accessPoint := fileSystem.AddAccessPoint(jsii.String("AccessPoint"), &AccessPointOptions{
//   	// set /export/lambda as the root of the access point
//   	Path: jsii.String("/export/lambda"),
//   	// as /export/lambda does not exist in a new efs filesystem, the efs will create the directory with the following createAcl
//   	CreateAcl: &Acl{
//   		OwnerUid: jsii.String("1001"),
//   		OwnerGid: jsii.String("1001"),
//   		Permissions: jsii.String("750"),
//   	},
//   	// enforce the POSIX identity so lambda function will access with this identity
//   	PosixUser: &PosixUser{
//   		Uid: jsii.String("1001"),
//   		Gid: jsii.String("1001"),
//   	},
//   })
//
//   fn := lambda.NewFunction(this, jsii.String("MyLambda"), &FunctionProps{
//   	// mount the access point to /mnt/msg in the lambda runtime environment
//   	Filesystem: lambda.FileSystem_FromEfsAccessPoint(accessPoint, jsii.String("/mnt/msg")),
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	Vpc: Vpc,
//   })
//
type AccessPointOptions struct {
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

