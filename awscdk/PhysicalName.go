package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Includes special markers for automatic generation of physical names.
type PhysicalName interface {
}

// The jsii proxy struct for PhysicalName
type jsiiProxy_PhysicalName struct {
	_ byte // padding
}

func PhysicalName_GENERATE_IF_NEEDED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.PhysicalName",
		"GENERATE_IF_NEEDED",
		&returns,
	)
	return returns
}

