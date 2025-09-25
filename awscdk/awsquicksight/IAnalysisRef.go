package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Analysis.
// Experimental.
type IAnalysisRef interface {
	constructs.IConstruct
	// A reference to a Analysis resource.
	// Experimental.
	AnalysisRef() *AnalysisReference
}

// The jsii proxy for IAnalysisRef
type jsiiProxy_IAnalysisRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAnalysisRef) AnalysisRef() *AnalysisReference {
	var returns *AnalysisReference
	_jsii_.Get(
		j,
		"analysisRef",
		&returns,
	)
	return returns
}

