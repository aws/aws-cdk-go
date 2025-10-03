package awscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AnalysisTemplate.
// Experimental.
type IAnalysisTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAnalysisTemplateRef
type jsiiProxy_IAnalysisTemplateRef struct {
	internal.Type__constructsIConstruct
}

