package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QueryDefinition.
// Experimental.
type IQueryDefinitionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a QueryDefinition resource.
	// Experimental.
	QueryDefinitionRef() *QueryDefinitionReference
}

// The jsii proxy for IQueryDefinitionRef
type jsiiProxy_IQueryDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IQueryDefinitionRef) QueryDefinitionRef() *QueryDefinitionReference {
	var returns *QueryDefinitionReference
	_jsii_.Get(
		j,
		"queryDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueryDefinitionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueryDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

