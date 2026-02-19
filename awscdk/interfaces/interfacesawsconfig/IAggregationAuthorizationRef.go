package interfacesawsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AggregationAuthorization.
// Experimental.
type IAggregationAuthorizationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AggregationAuthorization resource.
	// Experimental.
	AggregationAuthorizationRef() *AggregationAuthorizationReference
}

// The jsii proxy for IAggregationAuthorizationRef
type jsiiProxy_IAggregationAuthorizationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAggregationAuthorizationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAggregationAuthorizationRef) AggregationAuthorizationRef() *AggregationAuthorizationReference {
	var returns *AggregationAuthorizationReference
	_jsii_.Get(
		j,
		"aggregationAuthorizationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAggregationAuthorizationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAggregationAuthorizationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

