package interfacesawscleanroomsml

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscleanroomsml/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrainingDataset.
// Experimental.
type ITrainingDatasetRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TrainingDataset resource.
	// Experimental.
	TrainingDatasetRef() *TrainingDatasetReference
}

// The jsii proxy for ITrainingDatasetRef
type jsiiProxy_ITrainingDatasetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITrainingDatasetRef) TrainingDatasetRef() *TrainingDatasetReference {
	var returns *TrainingDatasetReference
	_jsii_.Get(
		j,
		"trainingDatasetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrainingDatasetRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrainingDatasetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

