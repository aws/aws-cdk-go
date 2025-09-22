package awsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConformancePack.
// Experimental.
type IConformancePackRef interface {
	constructs.IConstruct
	// A reference to a ConformancePack resource.
	// Experimental.
	ConformancePackRef() *ConformancePackReference
}

// The jsii proxy for IConformancePackRef
type jsiiProxy_IConformancePackRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConformancePackRef) ConformancePackRef() *ConformancePackReference {
	var returns *ConformancePackReference
	_jsii_.Get(
		j,
		"conformancePackRef",
		&returns,
	)
	return returns
}

