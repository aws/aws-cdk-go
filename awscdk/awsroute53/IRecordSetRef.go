package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RecordSet.
// Experimental.
type IRecordSetRef interface {
	constructs.IConstruct
	// A reference to a RecordSet resource.
	// Experimental.
	RecordSetRef() *RecordSetReference
}

// The jsii proxy for IRecordSetRef
type jsiiProxy_IRecordSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRecordSetRef) RecordSetRef() *RecordSetReference {
	var returns *RecordSetReference
	_jsii_.Get(
		j,
		"recordSetRef",
		&returns,
	)
	return returns
}

