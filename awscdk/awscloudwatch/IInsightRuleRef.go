package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InsightRule.
// Experimental.
type IInsightRuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IInsightRuleRef
type jsiiProxy_IInsightRuleRef struct {
	internal.Type__constructsIConstruct
}

