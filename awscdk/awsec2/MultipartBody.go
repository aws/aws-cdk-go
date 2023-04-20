package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The base class for all classes which can be used as `MultipartUserData`.
//
// Example:
//   multipartUserData := ec2.NewMultipartUserData()
//   commandsUserData := ec2.UserData_ForLinux()
//   multipartUserData.AddUserDataPart(commandsUserData, ec2.MultipartBody_SHELL_SCRIPT(), jsii.Boolean(true))
//
//   // Adding commands to the multipartUserData adds them to commandsUserData, and vice-versa.
//   multipartUserData.AddCommands(jsii.String("touch /root/multi.txt"))
//   commandsUserData.AddCommands(jsii.String("touch /root/userdata.txt"))
//
type MultipartBody interface {
	// Render body part as the string.
	//
	// Subclasses should not add leading nor trailing new line characters (\r \n).
	RenderBodyPart() *[]*string
}

// The jsii proxy struct for MultipartBody
type jsiiProxy_MultipartBody struct {
	_ byte // padding
}

func NewMultipartBody_Override(m MultipartBody) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.MultipartBody",
		nil, // no parameters
		m,
	)
}

// Constructs the raw `MultipartBody` using specified body, content type and transfer encoding.
//
// When transfer encoding is specified (typically as Base64), it's caller responsibility to convert body to
// Base64 either by wrapping with `Fn.base64` or by converting it by other converters.
func MultipartBody_FromRawBody(opts *MultipartBodyOptions) MultipartBody {
	_init_.Initialize()

	if err := validateMultipartBody_FromRawBodyParameters(opts); err != nil {
		panic(err)
	}
	var returns MultipartBody

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.MultipartBody",
		"fromRawBody",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

// Constructs the new `MultipartBody` wrapping existing `UserData`. Modification to `UserData` are reflected in subsequent renders of the part.
//
// For more information about content types see `MultipartBodyOptions.contentType`.
func MultipartBody_FromUserData(userData UserData, contentType *string) MultipartBody {
	_init_.Initialize()

	if err := validateMultipartBody_FromUserDataParameters(userData); err != nil {
		panic(err)
	}
	var returns MultipartBody

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.MultipartBody",
		"fromUserData",
		[]interface{}{userData, contentType},
		&returns,
	)

	return returns
}

func MultipartBody_CLOUD_BOOTHOOK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.MultipartBody",
		"CLOUD_BOOTHOOK",
		&returns,
	)
	return returns
}

func MultipartBody_SHELL_SCRIPT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.MultipartBody",
		"SHELL_SCRIPT",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MultipartBody) RenderBodyPart() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"renderBodyPart",
		nil, // no parameters
		&returns,
	)

	return returns
}

