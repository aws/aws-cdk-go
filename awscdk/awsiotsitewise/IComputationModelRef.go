package awsiotsitewise

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotsitewise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ComputationModel.
// Experimental.
type IComputationModelRef interface {
	constructs.IConstruct
	// A reference to a ComputationModel resource.
	// Experimental.
	ComputationModelRef() *ComputationModelReference
}

// The jsii proxy for IComputationModelRef
type jsiiProxy_IComputationModelRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IComputationModelRef) ComputationModelRef() *ComputationModelReference {
	var returns *ComputationModelReference
	_jsii_.Get(
		j,
		"computationModelRef",
		&returns,
	)
	return returns
}

