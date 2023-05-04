package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An object that represents reasons to restart an InitService.
//
// Pass an instance of this object to the `InitFile`, `InitCommand`,
// `InitSource` and `InitPackage` objects, and finally to an `InitService`
// itself to cause the actions (files, commands, sources, and packages)
// to trigger a restart of the service.
//
// For example, the following will run a custom command to install Nginx,
// and trigger the nginx service to be restarted after the command has run.
//
// ```ts
// const handle = new ec2.InitServiceRestartHandle();
// ec2.CloudFormationInit.fromElements(
//   ec2.InitCommand.shellCommand('/usr/bin/custom-nginx-install.sh', { serviceRestartHandles: [handle] }),
//   ec2.InitService.enable('nginx', { serviceRestartHandle: handle }),
// );
// ```.
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
type InitServiceRestartHandle interface {
}

// The jsii proxy struct for InitServiceRestartHandle
type jsiiProxy_InitServiceRestartHandle struct {
	_ byte // padding
}

func NewInitServiceRestartHandle() InitServiceRestartHandle {
	_init_.Initialize()

	j := jsiiProxy_InitServiceRestartHandle{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InitServiceRestartHandle",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewInitServiceRestartHandle_Override(i InitServiceRestartHandle) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InitServiceRestartHandle",
		nil, // no parameters
		i,
	)
}

