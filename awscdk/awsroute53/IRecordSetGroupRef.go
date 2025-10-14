package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RecordSetGroup.
// Experimental.
type IRecordSetGroupRef interface {
	constructs.IConstruct
	// A reference to a RecordSetGroup resource.
	// Experimental.
	RecordSetGroupRef() *RecordSetGroupReference
}

// The jsii proxy for IRecordSetGroupRef
type jsiiProxy_IRecordSetGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRecordSetGroupRef) RecordSetGroupRef() *RecordSetGroupReference {
	var returns *RecordSetGroupReference
	_jsii_.Get(
		j,
		"recordSetGroupRef",
		&returns,
	)
	return returns
}

