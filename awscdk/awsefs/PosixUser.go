package awsefs


// Represents the PosixUser.
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
type PosixUser struct {
	// The POSIX group ID used for all file system operations using this access point.
	Gid *string `field:"required" json:"gid" yaml:"gid"`
	// The POSIX user ID used for all file system operations using this access point.
	Uid *string `field:"required" json:"uid" yaml:"uid"`
	// Secondary POSIX group IDs used for all file system operations using this access point.
	// Default: - None.
	//
	SecondaryGids *[]*string `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
}

