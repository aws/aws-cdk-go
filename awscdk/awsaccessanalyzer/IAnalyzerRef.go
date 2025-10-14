package awsaccessanalyzer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsaccessanalyzer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Analyzer.
// Experimental.
type IAnalyzerRef interface {
	constructs.IConstruct
	// A reference to a Analyzer resource.
	// Experimental.
	AnalyzerRef() *AnalyzerReference
}

// The jsii proxy for IAnalyzerRef
type jsiiProxy_IAnalyzerRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAnalyzerRef) AnalyzerRef() *AnalyzerReference {
	var returns *AnalyzerReference
	_jsii_.Get(
		j,
		"analyzerRef",
		&returns,
	)
	return returns
}

