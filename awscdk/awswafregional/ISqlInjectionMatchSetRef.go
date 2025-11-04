package awswafregional

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafregional/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SqlInjectionMatchSet.
// Experimental.
type ISqlInjectionMatchSetRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SqlInjectionMatchSet resource.
	// Experimental.
	SqlInjectionMatchSetRef() *SqlInjectionMatchSetReference
}

// The jsii proxy for ISqlInjectionMatchSetRef
type jsiiProxy_ISqlInjectionMatchSetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISqlInjectionMatchSetRef) SqlInjectionMatchSetRef() *SqlInjectionMatchSetReference {
	var returns *SqlInjectionMatchSetReference
	_jsii_.Get(
		j,
		"sqlInjectionMatchSetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISqlInjectionMatchSetRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISqlInjectionMatchSetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

