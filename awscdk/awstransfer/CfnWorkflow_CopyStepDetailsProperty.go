package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   copyStepDetailsProperty := &CopyStepDetailsProperty{
//   	DestinationFileLocation: &S3FileLocationProperty{
//   		S3FileLocation: &S3InputFileLocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	OverwriteExisting: jsii.String("overwriteExisting"),
//   	SourceFileLocation: jsii.String("sourceFileLocation"),
//   }
//
type CfnWorkflow_CopyStepDetailsProperty struct {
	// `CfnWorkflow.CopyStepDetailsProperty.DestinationFileLocation`.
	DestinationFileLocation interface{} `field:"optional" json:"destinationFileLocation" yaml:"destinationFileLocation"`
	// `CfnWorkflow.CopyStepDetailsProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnWorkflow.CopyStepDetailsProperty.OverwriteExisting`.
	OverwriteExisting *string `field:"optional" json:"overwriteExisting" yaml:"overwriteExisting"`
	// `CfnWorkflow.CopyStepDetailsProperty.SourceFileLocation`.
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
}

