package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FindingAggregator.
// Experimental.
type IFindingAggregatorRef interface {
	constructs.IConstruct
	// A reference to a FindingAggregator resource.
	// Experimental.
	FindingAggregatorRef() *FindingAggregatorReference
}

// The jsii proxy for IFindingAggregatorRef
type jsiiProxy_IFindingAggregatorRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFindingAggregatorRef) FindingAggregatorRef() *FindingAggregatorReference {
	var returns *FindingAggregatorReference
	_jsii_.Get(
		j,
		"findingAggregatorRef",
		&returns,
	)
	return returns
}

