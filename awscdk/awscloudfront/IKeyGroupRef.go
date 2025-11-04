package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a KeyGroup.
// Experimental.
type IKeyGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a KeyGroup resource.
	// Experimental.
	KeyGroupRef() *KeyGroupReference
}

// The jsii proxy for IKeyGroupRef
type jsiiProxy_IKeyGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IKeyGroupRef) KeyGroupRef() *KeyGroupReference {
	var returns *KeyGroupReference
	_jsii_.Get(
		j,
		"keyGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeyGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeyGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

