package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Insight.
// Experimental.
type IInsightRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Insight resource.
	// Experimental.
	InsightRef() *InsightReference
}

// The jsii proxy for IInsightRef
type jsiiProxy_IInsightRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IInsightRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInsightRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

