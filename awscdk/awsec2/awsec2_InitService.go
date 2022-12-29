package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A services that be enabled, disabled or restarted when the instance is launched.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var myBucket bucket
//
//
//   handle := ec2.NewInitServiceRestartHandle()
//
//   ec2.cloudFormationInit.fromElements(ec2.initFile.fromString(jsii.String("/etc/nginx/nginx.conf"), jsii.String("..."), &initFileOptions{
//   	serviceRestartHandles: []initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.initSource.fromS3Object(jsii.String("/var/www/html"), myBucket, jsii.String("html.zip"), &initSourceOptions{
//   	serviceRestartHandles: []*initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.initService.enable(jsii.String("nginx"), &initServiceOptions{
//   	serviceRestartHandle: handle,
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

