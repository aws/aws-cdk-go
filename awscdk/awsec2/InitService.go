package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A services that be enabled, disabled or restarted when the instance is launched.
//
// Example:
//   var myBucket bucket
//
//
//   handle := ec2.NewInitServiceRestartHandle()
//
//   ec2.CloudFormationInit_FromElements(ec2.InitFile_FromString(jsii.String("/etc/nginx/nginx.conf"), jsii.String("..."), &InitFileOptions{
//   	ServiceRestartHandles: []initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.InitSource_FromS3Object(jsii.String("/var/www/html"), myBucket, jsii.String("html.zip"), &InitSourceOptions{
//   	ServiceRestartHandles: []*initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.InitService_Enable(jsii.String("nginx"), &InitServiceOptions{
//   	ServiceRestartHandle: handle,
//   }))
//
type InitService interface {
	InitElement
	// Returns the init element type for this element.
	ElementType() *string
}

// The jsii proxy struct for InitService
type jsiiProxy_InitService struct {
	jsiiProxy_InitElement
}

func (j *jsiiProxy_InitService) ElementType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elementType",
		&returns,
	)
	return returns
}


// Disable and stop the given service.
func InitService_Disable(serviceName *string) InitService {
	_init_.Initialize()

	if err := validateInitService_DisableParameters(serviceName); err != nil {
		panic(err)
	}
	var returns InitService

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitService",
		"disable",
		[]interface{}{serviceName},
		&returns,
	)

	return returns
}

// Enable and start the given service, optionally restarting it.
func InitService_Enable(serviceName *string, options *InitServiceOptions) InitService {
	_init_.Initialize()

	if err := validateInitService_EnableParameters(serviceName, options); err != nil {
		panic(err)
	}
	var returns InitService

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitService",
		"enable",
		[]interface{}{serviceName, options},
		&returns,
	)

	return returns
}

// Install a systemd-compatible config file for the given service.
//
// This is a helper function to create a simple systemd configuration
// file that will allow running a service on the machine using `InitService.enable()`.
//
// Systemd allows many configuration options; this function does not pretend
// to expose all of them. If you need advanced configuration options, you
// can use `InitFile` to create exactly the configuration file you need
// at `/etc/systemd/system/${serviceName}.service`.
func InitService_SystemdConfigFile(serviceName *string, options *SystemdConfigFileOptions) InitFile {
	_init_.Initialize()

	if err := validateInitService_SystemdConfigFileParameters(serviceName, options); err != nil {
		panic(err)
	}
	var returns InitFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitService",
		"systemdConfigFile",
		[]interface{}{serviceName, options},
		&returns,
	)

	return returns
}

