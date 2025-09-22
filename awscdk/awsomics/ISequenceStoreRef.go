package awsomics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SequenceStore.
// Experimental.
type ISequenceStoreRef interface {
	constructs.IConstruct
	// A reference to a SequenceStore resource.
	// Experimental.
	SequenceStoreRef() *SequenceStoreReference
}

// The jsii proxy for ISequenceStoreRef
type jsiiProxy_ISequenceStoreRef struct {
	internal.Type__constructsIConstruct
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

