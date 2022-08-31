package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Instance User Data.
//
// Example:
//   multipartUserData := ec2.NewMultipartUserData()
//   commandsUserData := ec2.userData.forLinux()
//   multipartUserData.addUserDataPart(commandsUserData, ec2.multipartBody_SHELL_SCRIPT(), jsii.Boolean(true))
//
//   // Adding commands to the multipartUserData adds them to commandsUserData, and vice-versa.
//   multipartUserData.addCommands(jsii.String("touch /root/multi.txt"))
//   commandsUserData.addCommands(jsii.String("touch /root/userdata.txt"))
//
// Experimental.
type UserData interface {
	// Add one or more commands to the user data.
	// Experimental.
	AddCommands(commands ...*string)
	// Adds commands to execute a file.
	// Experimental.
	AddExecuteFileCommand(params *ExecuteFileOptions)
	// Add one or more commands to the user data that will run when the script exits.
	// Experimental.
	AddOnExitCommands(commands ...*string)
	// Adds commands to download a file from S3.
	//
	// Returns: : The local path that the file will be downloaded to.
	// Experimental.
	AddS3DownloadCommand(params *S3DownloadOptions) *string
	// Adds a command which will send a cfn-signal when the user data script ends.
	// Experimental.
	AddSignalOnExitCommand(resource awscdk.Resource)
	// Render the UserData for use in a construct.
	// Experimental.
	Render() *string
}

// The jsii proxy struct for UserData
type jsiiProxy_UserData struct {
	_ byte // padding
}

// Experimental.
func NewUserData_Override(u UserData) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.UserData",
		nil, // no parameters
		u,
	)
}

// Create a userdata object with custom content.
// Experimental.
func UserData_Custom(content *string) UserData {
	_init_.Initialize()

	var returns UserData

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.UserData",
		"custom",
		[]interface{}{content},
		&returns,
	)

	return returns
}

// Create a userdata object for Linux hosts.
// Experimental.
func UserData_ForLinux(options *LinuxUserDataOptions) UserData {
	_init_.Initialize()

	var returns UserData

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.UserData",
		"forLinux",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Experimental.
func UserData_ForOperatingSystem(os OperatingSystemType) UserData {
	_init_.Initialize()

	var returns UserData

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.UserData",
		"forOperatingSystem",
		[]interface{}{os},
		&returns,
	)

	return returns
}

// Create a userdata object for Windows hosts.
// Experimental.
func UserData_ForWindows() UserData {
	_init_.Initialize()

	var returns UserData

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.UserData",
		"forWindows",
		nil, // no parameters
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

