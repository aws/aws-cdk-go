package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
)

// An option group.
type IOptionGroup interface {
	awscdk.IResource
	// Adds a configuration to this OptionGroup.
	//
	// This method is a no-op for an imported OptionGroup.
	//
	// Returns: true if the OptionConfiguration was successfully added.
	AddConfiguration(configuration *OptionConfiguration) *bool
	// The name of the option group.
	OptionGroupName() *string
}

// The jsii proxy for IOptionGroup
type jsiiProxy_IOptionGroup struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IOptionGroup) AddConfiguration(configuration *OptionConfiguration) *bool {
	if err := i.validateAddConfigurationParameters(configuration); err != nil {
		panic(err)
	}
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

