package awswafregional

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafregional/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ByteMatchSet.
// Experimental.
type IByteMatchSetRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ByteMatchSet resource.
	// Experimental.
	ByteMatchSetRef() *ByteMatchSetReference
}

// The jsii proxy for IByteMatchSetRef
type jsiiProxy_IByteMatchSetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IByteMatchSetRef) ByteMatchSetRef() *ByteMatchSetReference {
	var returns *ByteMatchSetReference
	_jsii_.Get(
		j,
		"byteMatchSetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IByteMatchSetRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IByteMatchSetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

