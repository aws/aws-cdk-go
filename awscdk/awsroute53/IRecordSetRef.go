package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RecordSet.
// Experimental.
type IRecordSetRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RecordSet resource.
	// Experimental.
	RecordSetRef() *RecordSetReference
}

// The jsii proxy for IRecordSetRef
type jsiiProxy_IRecordSetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IRecordSetRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecordSetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

