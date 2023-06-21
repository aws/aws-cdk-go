package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Base class for all CloudFormation Init elements.
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
type InitElement interface {
	// Returns the init element type for this element.
	ElementType() *string
}

// The jsii proxy struct for InitElement
type jsiiProxy_InitElement struct {
	_ byte // padding
}

func (j *jsiiProxy_InitElement) ElementType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elementType",
		&returns,
	)
	return returns
}


func NewInitElement_Override(i InitElement) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InitElement",
		nil, // no parameters
		i,
	)
}

