package awssmsvoice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssmsvoice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SenderId.
// Experimental.
type ISenderIdRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SenderId resource.
	// Experimental.
	SenderIdRef() *SenderIdReference
}

// The jsii proxy for ISenderIdRef
type jsiiProxy_ISenderIdRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISenderIdRef) SenderIdRef() *SenderIdReference {
	var returns *SenderIdReference
	_jsii_.Get(
		j,
		"senderIdRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISenderIdRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISenderIdRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

