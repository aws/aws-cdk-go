package awscustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CalculatedAttributeDefinition.
// Experimental.
type ICalculatedAttributeDefinitionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CalculatedAttributeDefinition resource.
	// Experimental.
	CalculatedAttributeDefinitionRef() *CalculatedAttributeDefinitionReference
}

// The jsii proxy for ICalculatedAttributeDefinitionRef
type jsiiProxy_ICalculatedAttributeDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICalculatedAttributeDefinitionRef) CalculatedAttributeDefinitionRef() *CalculatedAttributeDefinitionReference {
	var returns *CalculatedAttributeDefinitionReference
	_jsii_.Get(
		j,
		"calculatedAttributeDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICalculatedAttributeDefinitionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICalculatedAttributeDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

