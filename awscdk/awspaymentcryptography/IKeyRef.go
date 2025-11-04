package awspaymentcryptography

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspaymentcryptography/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Key.
// Experimental.
type IKeyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Key resource.
	// Experimental.
	KeyRef() *KeyReference
}

// The jsii proxy for IKeyRef
type jsiiProxy_IKeyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IKeyRef) KeyRef() *KeyReference {
	var returns *KeyReference
	_jsii_.Get(
		j,
		"keyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

