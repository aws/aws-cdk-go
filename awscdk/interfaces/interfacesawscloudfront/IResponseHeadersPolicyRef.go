package interfacesawscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResponseHeadersPolicy.
// Experimental.
type IResponseHeadersPolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ResponseHeadersPolicy resource.
	// Experimental.
	ResponseHeadersPolicyRef() *ResponseHeadersPolicyReference
}

// The jsii proxy for IResponseHeadersPolicyRef
type jsiiProxy_IResponseHeadersPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IResponseHeadersPolicyRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IResponseHeadersPolicyRef) ResponseHeadersPolicyRef() *ResponseHeadersPolicyReference {
	var returns *ResponseHeadersPolicyReference
	_jsii_.Get(
		j,
		"responseHeadersPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResponseHeadersPolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResponseHeadersPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

