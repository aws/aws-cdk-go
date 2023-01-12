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
//   fileSystem := efs.NewFileSystem(this, jsii.String("Efs"), &fileSystemProps{
//   	vpc: vpc,
//   })
//
//   // create a new access point from the filesystem
//   accessPoint := fileSystem.addAccessPoint(jsii.String("AccessPoint"), &accessPointOptions{
//   	// set /export/lambda as the root of the access point
//   	path: jsii.String("/export/lambda"),
//   	// as /export/lambda does not exist in a new efs filesystem, the efs will create the directory with the following createAcl
//   	createAcl: &acl{
//   		ownerUid: jsii.String("1001"),
//   		ownerGid: jsii.String("1001"),
//   		permissions: jsii.String("750"),
//   	},
//   	// enforce the POSIX identity so lambda function will access with this identity
//   	posixUser: &posixUser{
//   		uid: jsii.String("1001"),
//   		gid: jsii.String("1001"),
//   	},
//   })
//
//   fn := lambda.NewFunction(this, jsii.String("MyLambda"), &functionProps{
//   	// mount the access point to /mnt/msg in the lambda runtime environment
//   	filesystem: lambda.fileSystem.fromEfsAccessPoint(accessPoint, jsii.String("/mnt/msg")),
//   	runtime: lambda.runtime_NODEJS_18_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	vpc: vpc,
//   })
//
type PosixUser struct {
	// The POSIX group ID used for all file system operations using this access point.
	Gid *string `field:"required" json:"gid" yaml:"gid"`
	// The POSIX user ID used for all file system operations using this access point.
	Uid *string `field:"required" json:"uid" yaml:"uid"`
	// Secondary POSIX group IDs used for all file system operations using this access point.
	SecondaryGids *[]*string `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
}

