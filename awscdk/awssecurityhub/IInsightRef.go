package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Insight.
// Experimental.
type IInsightRef interface {
	constructs.IConstruct
	// A reference to a Insight resource.
	// Experimental.
	InsightRef() *InsightReference
}

// The jsii proxy for IInsightRef
type jsiiProxy_IInsightRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IInsightRef) InsightRef() *InsightReference {
	var returns *InsightReference
	_jsii_.Get(
		j,
		"insightRef",
		&returns,
	)
	return returns
}

