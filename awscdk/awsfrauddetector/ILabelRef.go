package awsfrauddetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfrauddetector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Label.
// Experimental.
type ILabelRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Label resource.
	// Experimental.
	LabelRef() *LabelReference
}

// The jsii proxy for ILabelRef
type jsiiProxy_ILabelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ILabelRef) LabelRef() *LabelReference {
	var returns *LabelReference
	_jsii_.Get(
		j,
		"labelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILabelRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILabelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

