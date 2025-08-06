package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A class used to generate resource arns for AppSync Event APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appSyncEventResource := awscdk.Aws_appsync.AppSyncEventResource_All()
//
type AppSyncEventResource interface {
	// Return the Resource ARN.
	ResourceArns(api EventApiBase) *[]*string
}

// The jsii proxy struct for AppSyncEventResource
type jsiiProxy_AppSyncEventResource struct {
	_ byte // padding
}

// Generate the resource names that accepts all types: `*`.
func AppSyncEventResource_All() AppSyncEventResource {
	_init_.Initialize()

	var returns AppSyncEventResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.AppSyncEventResource",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Generate the resource names that accepts all channel namespaces: `*`.
func AppSyncEventResource_AllChannelNamespaces() AppSyncEventResource {
	_init_.Initialize()

	var returns AppSyncEventResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.AppSyncEventResource",
		"allChannelNamespaces",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Generate a resource for the calling API.
func AppSyncEventResource_ForAPI() AppSyncEventResource {
	_init_.Initialize()

	var returns AppSyncEventResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.AppSyncEventResource",
		"forAPI",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Generate the resource names given a channel namespace.
func AppSyncEventResource_OfChannelNamespace(channelNamespace *string) AppSyncEventResource {
	_init_.Initialize()

	if err := validateAppSyncEventResource_OfChannelNamespaceParameters(channelNamespace); err != nil {
		panic(err)
	}
	var returns AppSyncEventResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.AppSyncEventResource",
		"ofChannelNamespace",
		[]interface{}{channelNamespace},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncEventResource) ResourceArns(api EventApiBase) *[]*string {
	if err := a.validateResourceArnsParameters(api); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"resourceArns",
		[]interface{}{api},
		&returns,
	)

	return returns
}

