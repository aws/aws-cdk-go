package awsomics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReferenceStore.
// Experimental.
type IReferenceStoreRef interface {
	constructs.IConstruct
	// A reference to a ReferenceStore resource.
	// Experimental.
	ReferenceStoreRef() *ReferenceStoreReference
}

// The jsii proxy for IReferenceStoreRef
type jsiiProxy_IReferenceStoreRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IReferenceStoreRef) ReferenceStoreRef() *ReferenceStoreReference {
	var returns *ReferenceStoreReference
	_jsii_.Get(
		j,
		"referenceStoreRef",
		&returns,
	)
	return returns
}

