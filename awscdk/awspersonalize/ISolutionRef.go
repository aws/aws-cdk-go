package awspersonalize

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspersonalize/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Solution.
// Experimental.
type ISolutionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Solution resource.
	// Experimental.
	SolutionRef() *SolutionReference
}

// The jsii proxy for ISolutionRef
type jsiiProxy_ISolutionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISolutionRef) SolutionRef() *SolutionReference {
	var returns *SolutionReference
	_jsii_.Get(
		j,
		"solutionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISolutionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISolutionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

