package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
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
// Experimental.
type InitElement interface {
	// Returns the init element type for this element.
	// Experimental.
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


// Experimental.
func NewInitElement_Override(i InitElement) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.InitElement",
		nil, // no parameters
		i,
	)
}

