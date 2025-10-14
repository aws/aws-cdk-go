package awsomics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RunGroup.
// Experimental.
type IRunGroupRef interface {
	constructs.IConstruct
	// A reference to a RunGroup resource.
	// Experimental.
	RunGroupRef() *RunGroupReference
}

// The jsii proxy for IRunGroupRef
type jsiiProxy_IRunGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRunGroupRef) RunGroupRef() *RunGroupReference {
	var returns *RunGroupReference
	_jsii_.Get(
		j,
		"runGroupRef",
		&returns,
	)
	return returns
}

