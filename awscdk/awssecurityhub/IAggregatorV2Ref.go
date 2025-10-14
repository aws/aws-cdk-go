package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AggregatorV2.
// Experimental.
type IAggregatorV2Ref interface {
	constructs.IConstruct
	// A reference to a AggregatorV2 resource.
	// Experimental.
	AggregatorV2Ref() *AggregatorV2Reference
}

// The jsii proxy for IAggregatorV2Ref
type jsiiProxy_IAggregatorV2Ref struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAggregatorV2Ref) AggregatorV2Ref() *AggregatorV2Reference {
	var returns *AggregatorV2Reference
	_jsii_.Get(
		j,
		"aggregatorV2Ref",
		&returns,
	)
	return returns
}

