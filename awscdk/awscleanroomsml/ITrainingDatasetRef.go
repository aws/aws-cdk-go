package awscleanroomsml

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanroomsml/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrainingDataset.
// Experimental.
type ITrainingDatasetRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITrainingDatasetRef
type jsiiProxy_ITrainingDatasetRef struct {
	internal.Type__constructsIConstruct
}

