package interfacesawscleanroomsml


// A reference to a TrainingDataset resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trainingDatasetReference := &TrainingDatasetReference{
//   	TrainingDatasetArn: jsii.String("trainingDatasetArn"),
//   }
//
type TrainingDatasetReference struct {
	// The TrainingDatasetArn of the TrainingDataset resource.
	TrainingDatasetArn *string `field:"required" json:"trainingDatasetArn" yaml:"trainingDatasetArn"`
}

