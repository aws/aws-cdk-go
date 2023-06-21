package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Mime multipart user data.
//
// This class represents MIME multipart user data, as described in.
// [Specifying Multiple User Data Blocks Using a MIME Multi Part Archive](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/bootstrap_container_instance.html#multi-part_user_data)
//
// Example:
//   bootHookConf := ec2.UserData_ForLinux()
//   bootHookConf.AddCommands(jsii.String("cloud-init-per once docker_options echo 'OPTIONS=\"${OPTIONS} --storage-opt dm.basesize=40G\"' >> /etc/sysconfig/docker"))
//
//   setupCommands := ec2.UserData_ForLinux()
//   setupCommands.AddCommands(jsii.String("sudo yum install awscli && echo Packages installed らと > /var/tmp/setup"))
//
//   multipartUserData := ec2.NewMultipartUserData()
//   // The docker has to be configured at early stage, so content type is overridden to boothook
//   multipartUserData.AddPart(ec2.MultipartBody_FromUserData(bootHookConf, jsii.String("text/cloud-boothook; charset=\"us-ascii\"")))
//   // Execute the rest of setup
//   multipartUserData.AddPart(ec2.MultipartBody_FromUserData(setupCommands))
//
//   ec2.NewLaunchTemplate(this, jsii.String(""), &LaunchTemplateProps{
//   	UserData: multipartUserData,
//   	BlockDevices: []blockDevice{
//   	},
//   })
//
type MultipartUserData interface {
	UserData
	// Add one or more commands to the user data.
	AddCommands(commands ...*string)
	// Adds commands to execute a file.
	AddExecuteFileCommand(params *ExecuteFileOptions)
	// Add one or more commands to the user data that will run when the script exits.
	AddOnExitCommands(commands ...*string)
	// Adds a part to the list of parts.
	AddPart(part MultipartBody)
	// Adds commands to download a file from S3.
	AddS3DownloadCommand(params *S3DownloadOptions) *string
	// Adds a command which will send a cfn-signal when the user data script ends.
	AddSignalOnExitCommand(resource awscdk.Resource)
	// Adds a multipart part based on a UserData object.
	//
	// If `makeDefault` is true, then the UserData added by this method
	// will also be the target of calls to the `add*Command` methods on
	// this MultipartUserData object.
	//
	// If `makeDefault` is false, then this is the same as calling:
	//
	// ```ts
	// declare const multiPart: ec2.MultipartUserData;
	// declare const userData: ec2.UserData;
	// declare const contentType: string;
	//
	// multiPart.addPart(ec2.MultipartBody.fromUserData(userData, contentType));
	// ```
	//
	// An undefined `makeDefault` defaults to either:
	// - `true` if no default UserData has been set yet; or
	// - `false` if there is no default UserData set.
	AddUserDataPart(userData UserData, contentType *string, makeDefault *bool)
	// Render the UserData for use in a construct.
	Render() *string
}

// The jsii proxy struct for MultipartUserData
type jsiiProxy_MultipartUserData struct {
	jsiiProxy_UserData
}

func NewMultipartUserData(opts *MultipartUserDataOptions) MultipartUserData {
	_init_.Initialize()

	if err := validateNewMultipartUserDataParameters(opts); err != nil {
		panic(err)
	}
	j := jsiiProxy_MultipartUserData{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.MultipartUserData",
		[]interface{}{opts},
		&j,
	)

	return &j
}

func NewMultipartUserData_Override(m MultipartUserData, opts *MultipartUserDataOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.MultipartUserData",
		[]interface{}{opts},
		m,
	)
}

// Create a userdata object with custom content.
func MultipartUserData_Custom(content *string) UserData {
	_init_.Initialize()

	if err := validateMultipartUserData_CustomParameters(content); err != nil {
		panic(err)
	}
	var returns UserData

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.MultipartUserData",
		"custom",
		[]interface{}{content},
		&returns,
	)

	return returns
}

// Create a userdata object for Linux hosts.
func MultipartUserData_ForLinux(options *LinuxUserDataOptions) UserData {
	_init_.Initialize()

	if err := validateMultipartUserData_ForLinuxParameters(options); err != nil {
		panic(err)
	}
	var returns UserData

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.MultipartUserData",
		"forLinux",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func MultipartUserData_ForOperatingSystem(os OperatingSystemType) UserData {
	_init_.Initialize()

	if err := validateMultipartUserData_ForOperatingSystemParameters(os); err != nil {
		panic(err)
	}
	var returns UserData

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.MultipartUserData",
		"forOperatingSystem",
		[]interface{}{os},
		&returns,
	)

	return returns
}

// Create a userdata object for Windows hosts.
func MultipartUserData_ForWindows(options *WindowsUserDataOptions) UserData {
	_init_.Initialize()

	if err := validateMultipartUserData_ForWindowsParameters(options); err != nil {
		panic(err)
	}
	var returns UserData

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.MultipartUserData",
		"forWindows",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MultipartUserData) AddCommands(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		m,
		"addCommands",
		args,
	)
}

func (m *jsiiProxy_MultipartUserData) AddExecuteFileCommand(params *ExecuteFileOptions) {
	if err := m.validateAddExecuteFileCommandParameters(params); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addExecuteFileCommand",
		[]interface{}{params},
	)
}

func (m *jsiiProxy_MultipartUserData) AddOnExitCommands(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		m,
		"addOnExitCommands",
		args,
	)
}

func (m *jsiiProxy_MultipartUserData) AddPart(part MultipartBody) {
	if err := m.validateAddPartParameters(part); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addPart",
		[]interface{}{part},
	)
}

func (m *jsiiProxy_MultipartUserData) AddS3DownloadCommand(params *S3DownloadOptions) *string {
	if err := m.validateAddS3DownloadCommandParameters(params); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"addS3DownloadCommand",
		[]interface{}{params},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MultipartUserData) AddSignalOnExitCommand(resource awscdk.Resource) {
	if err := m.validateAddSignalOnExitCommandParameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addSignalOnExitCommand",
		[]interface{}{resource},
	)
}

func (m *jsiiProxy_MultipartUserData) AddUserDataPart(userData UserData, contentType *string, makeDefault *bool) {
	if err := m.validateAddUserDataPartParameters(userData); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addUserDataPart",
		[]interface{}{userData, contentType, makeDefault},
	)
}

func (m *jsiiProxy_MultipartUserData) Render() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

