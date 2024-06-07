package awseks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Endpoint access characteristics.
//
// Example:
//   cluster := eks.NewCluster(this, jsii.String("hello-eks"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_30(),
//   	EndpointAccess: eks.EndpointAccess_PRIVATE(),
//   })
//
type EndpointAccess interface {
	// Restrict public access to specific CIDR blocks.
	//
	// If public access is disabled, this method will result in an error.
	OnlyFrom(cidr ...*string) EndpointAccess
}

// The jsii proxy struct for EndpointAccess
type jsiiProxy_EndpointAccess struct {
	_ byte // padding
}

func EndpointAccess_PRIVATE() EndpointAccess {
	_init_.Initialize()
	var returns EndpointAccess
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.EndpointAccess",
		"PRIVATE",
		&returns,
	)
	return returns
}

func EndpointAccess_PUBLIC() EndpointAccess {
	_init_.Initialize()
	var returns EndpointAccess
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.EndpointAccess",
		"PUBLIC",
		&returns,
	)
	return returns
}

func EndpointAccess_PUBLIC_AND_PRIVATE() EndpointAccess {
	_init_.Initialize()
	var returns EndpointAccess
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.EndpointAccess",
		"PUBLIC_AND_PRIVATE",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EndpointAccess) OnlyFrom(cidr ...*string) EndpointAccess {
	args := []interface{}{}
	for _, a := range cidr {
		args = append(args, a)
	}

	var returns EndpointAccess

	_jsii_.Invoke(
		e,
		"onlyFrom",
		args,
		&returns,
	)

	return returns
}

