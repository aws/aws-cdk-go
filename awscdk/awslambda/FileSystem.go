package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsefs"
)

// Represents the filesystem for the Lambda function.
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
type FileSystem interface {
	// the FileSystem configurations for the Lambda function.
	Config() *FileSystemConfig
}

// The jsii proxy struct for FileSystem
type jsiiProxy_FileSystem struct {
	_ byte // padding
}

func (j *jsiiProxy_FileSystem) Config() *FileSystemConfig {
	var returns *FileSystemConfig
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}


func NewFileSystem(config *FileSystemConfig) FileSystem {
	_init_.Initialize()

	if err := validateNewFileSystemParameters(config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FileSystem{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.FileSystem",
		[]interface{}{config},
		&j,
	)

	return &j
}

func NewFileSystem_Override(f FileSystem, config *FileSystemConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.FileSystem",
		[]interface{}{config},
		f,
	)
}

// mount the filesystem from Amazon EFS.
func FileSystem_FromEfsAccessPoint(ap awsefs.IAccessPoint, mountPath *string) FileSystem {
	_init_.Initialize()

	if err := validateFileSystem_FromEfsAccessPointParameters(ap, mountPath); err != nil {
		panic(err)
	}
	var returns FileSystem

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.FileSystem",
		"fromEfsAccessPoint",
		[]interface{}{ap, mountPath},
		&returns,
	)

	return returns
}

