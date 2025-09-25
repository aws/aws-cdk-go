package awscleanroomsml

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanroomsml/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrainingDataset.
// Experimental.
type ITrainingDatasetRef interface {
	constructs.IConstruct
	// A reference to a TrainingDataset resource.
	// Experimental.
	TrainingDatasetRef() *TrainingDatasetReference
}

// The jsii proxy for ITrainingDatasetRef
type jsiiProxy_ITrainingDatasetRef struct {
	internal.Type__constructsIConstruct
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

