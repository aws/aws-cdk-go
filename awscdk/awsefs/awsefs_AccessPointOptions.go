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
type AccessPointOptions struct {
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
}

