package awscleanrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AnalysisTemplate.
// Experimental.
type IAnalysisTemplateRef interface {
	constructs.IConstruct
	// A reference to a AnalysisTemplate resource.
	// Experimental.
	AnalysisTemplateRef() *AnalysisTemplateReference
}

// The jsii proxy for IAnalysisTemplateRef
type jsiiProxy_IAnalysisTemplateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAnalysisTemplateRef) AnalysisTemplateRef() *AnalysisTemplateReference {
	var returns *AnalysisTemplateReference
	_jsii_.Get(
		j,
		"analysisTemplateRef",
		&returns,
	)
	return returns
}

