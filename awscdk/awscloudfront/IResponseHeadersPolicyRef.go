package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResponseHeadersPolicy.
// Experimental.
type IResponseHeadersPolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ResponseHeadersPolicy resource.
	// Experimental.
	ResponseHeadersPolicyRef() *ResponseHeadersPolicyReference
}

// The jsii proxy for IResponseHeadersPolicyRef
type jsiiProxy_IResponseHeadersPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IResponseHeadersPolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

