package awsomics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SequenceStore.
// Experimental.
type ISequenceStoreRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SequenceStore resource.
	// Experimental.
	SequenceStoreRef() *SequenceStoreReference
}

// The jsii proxy for ISequenceStoreRef
type jsiiProxy_ISequenceStoreRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISequenceStoreRef) SequenceStoreRef() *SequenceStoreReference {
	var returns *SequenceStoreReference
	_jsii_.Get(
		j,
		"sequenceStoreRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISequenceStoreRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISequenceStoreRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

