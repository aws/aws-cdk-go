package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OptionGroup.
// Experimental.
type IOptionGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a OptionGroup resource.
	// Experimental.
	OptionGroupRef() *OptionGroupReference
}

// The jsii proxy for IOptionGroupRef
type jsiiProxy_IOptionGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IOptionGroupRef) OptionGroupRef() *OptionGroupReference {
	var returns *OptionGroupReference
	_jsii_.Get(
		j,
		"optionGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOptionGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOptionGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

