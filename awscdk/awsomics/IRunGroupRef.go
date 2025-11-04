package awsomics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RunGroup.
// Experimental.
type IRunGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RunGroup resource.
	// Experimental.
	RunGroupRef() *RunGroupReference
}

// The jsii proxy for IRunGroupRef
type jsiiProxy_IRunGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IRunGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRunGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

