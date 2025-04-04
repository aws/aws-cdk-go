package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Instance User Data.
//
// Example:
//   var cluster cluster
//
//   userData := ec2.UserData_ForLinux()
//   userData.AddCommands(jsii.String("set -o xtrace"),
//   fmt.Sprintf("/etc/eks/bootstrap.sh %v", cluster.ClusterName))
//   lt := ec2.NewCfnLaunchTemplate(this, jsii.String("LaunchTemplate"), &CfnLaunchTemplateProps{
//   	LaunchTemplateData: &LaunchTemplateDataProperty{
//   		ImageId: jsii.String("some-ami-id"),
//   		 // custom AMI
//   		InstanceType: jsii.String("t3.small"),
//   		UserData: awscdk.Fn_Base64(userData.Render()),
//   	},
//   })
//   cluster.AddNodegroupCapacity(jsii.String("extra-ng"), &NodegroupOptions{
//   	LaunchTemplateSpec: &LaunchTemplateSpec{
//   		Id: lt.ref,
//   		Version: lt.AttrLatestVersionNumber,
//   	},
//   })
//
type UserData interface {
	// Add one or more commands to the user data.
	AddCommands(commands ...*string)
	// Adds commands to execute a file.
	AddExecuteFileCommand(params *ExecuteFileOptions)
	// Add one or more commands to the user data that will run when the script exits.
	AddOnExitCommands(commands ...*string)
	// Adds commands to download a file from S3.
	//
	// Returns: : The local path that the file will be downloaded to.
	AddS3DownloadCommand(params *S3DownloadOptions) *string
	// Adds a command which will send a cfn-signal when the user data script ends.
	AddSignalOnExitCommand(resource awscdk.Resource)
	// Render the UserData for use in a construct.
	Render() *string
}

// The jsii proxy struct for UserData
type jsiiProxy_UserData struct {
	_ byte // padding
}

func NewUserData_Override(u UserData) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.UserData",
		nil, // no parameters
		u,
	)
}

// Create a userdata object with custom content.
func UserData_Custom(content *string) UserData {
	_init_.Initialize()

	if err := validateUserData_CustomParameters(content); err != nil {
		panic(err)
	}
	var returns UserData

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.UserData",
		"custom",
		[]interface{}{content},
		&returns,
	)

	return returns
}

// Create a userdata object for Linux hosts.
func UserData_ForLinux(options *LinuxUserDataOptions) UserData {
	_init_.Initialize()

	if err := validateUserData_ForLinuxParameters(options); err != nil {
		panic(err)
	}
	var returns UserData

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.UserData",
		"forLinux",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func UserData_ForOperatingSystem(os OperatingSystemType) UserData {
	_init_.Initialize()

	if err := validateUserData_ForOperatingSystemParameters(os); err != nil {
		panic(err)
	}
	var returns UserData

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.UserData",
		"forOperatingSystem",
		[]interface{}{os},
		&returns,
	)

	return returns
}

// Create a userdata object for Windows hosts.
func UserData_ForWindows(options *WindowsUserDataOptions) UserData {
	_init_.Initialize()

	if err := validateUserData_ForWindowsParameters(options); err != nil {
		panic(err)
	}
	var returns UserData

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.UserData",
		"forWindows",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserData) AddCommands(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		u,
		"addCommands",
		args,
	)
}

func (u *jsiiProxy_UserData) AddExecuteFileCommand(params *ExecuteFileOptions) {
	if err := u.validateAddExecuteFileCommandParameters(params); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addExecuteFileCommand",
		[]interface{}{params},
	)
}

func (u *jsiiProxy_UserData) AddOnExitCommands(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		u,
		"addOnExitCommands",
		args,
	)
}

func (u *jsiiProxy_UserData) AddS3DownloadCommand(params *S3DownloadOptions) *string {
	if err := u.validateAddS3DownloadCommandParameters(params); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"addS3DownloadCommand",
		[]interface{}{params},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserData) AddSignalOnExitCommand(resource awscdk.Resource) {
	if err := u.validateAddSignalOnExitCommandParameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addSignalOnExitCommand",
		[]interface{}{resource},
	)
}

func (u *jsiiProxy_UserData) Render() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

