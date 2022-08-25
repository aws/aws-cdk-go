package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsrds/internal"
)

// An option group.
// Experimental.
type IOptionGroup interface {
	awscdk.IResource
	// Adds a configuration to this OptionGroup.
	//
	// This method is a no-op for an imported OptionGroup.
	//
	// Returns: true if the OptionConfiguration was successfully added.
	// Experimental.
	AddConfiguration(configuration *OptionConfiguration) *bool
	// The name of the option group.
	// Experimental.
	OptionGroupName() *string
}

// The jsii proxy for IOptionGroup
type jsiiProxy_IOptionGroup struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IOptionGroup) AddConfiguration(configuration *OptionConfiguration) *bool {
	var returns *bool

	_jsii_.Invoke(
		i,
		"addConfiguration",
		[]interface{}{configuration},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IOptionGroup) OptionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupName",
		&returns,
	)
	return returns
}

