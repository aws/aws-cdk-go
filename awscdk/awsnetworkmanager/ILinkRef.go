package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Link.
// Experimental.
type ILinkRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Link resource.
	// Experimental.
	LinkRef() *LinkReference
}

// The jsii proxy for ILinkRef
type jsiiProxy_ILinkRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ILinkRef) LinkRef() *LinkReference {
	var returns *LinkReference
	_jsii_.Get(
		j,
		"linkRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILinkRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILinkRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

