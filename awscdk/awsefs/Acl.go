package awsefs


// Permissions as POSIX ACL.
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
type Acl struct {
	// Specifies the POSIX group ID to apply to the RootDirectory.
	//
	// Accepts values from 0 to 2^32 (4294967295).
	OwnerGid *string `field:"required" json:"ownerGid" yaml:"ownerGid"`
	// Specifies the POSIX user ID to apply to the RootDirectory.
	//
	// Accepts values from 0 to 2^32 (4294967295).
	OwnerUid *string `field:"required" json:"ownerUid" yaml:"ownerUid"`
	// Specifies the POSIX permissions to apply to the RootDirectory, in the format of an octal number representing the file's mode bits.
	Permissions *string `field:"required" json:"permissions" yaml:"permissions"`
}

