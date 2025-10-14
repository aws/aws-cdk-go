package awscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InsightRule.
// Experimental.
type IInsightRuleRef interface {
	constructs.IConstruct
	// A reference to a InsightRule resource.
	// Experimental.
	InsightRuleRef() *InsightRuleReference
}

// The jsii proxy for IInsightRuleRef
type jsiiProxy_IInsightRuleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IInsightRuleRef) InsightRuleRef() *InsightRuleReference {
	var returns *InsightRuleReference
	_jsii_.Get(
		j,
		"insightRuleRef",
		&returns,
	)
	return returns
}

