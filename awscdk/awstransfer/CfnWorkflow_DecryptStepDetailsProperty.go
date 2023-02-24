package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   decryptStepDetailsProperty := &DecryptStepDetailsProperty{
//   	DestinationFileLocation: &InputFileLocationProperty{
//   		EfsFileLocation: &EfsInputFileLocationProperty{
//   			FileSystemId: jsii.String("fileSystemId"),
//   			Path: jsii.String("path"),
//   		},
//   		S3FileLocation: &S3InputFileLocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	OverwriteExisting: jsii.String("overwriteExisting"),
//   	SourceFileLocation: jsii.String("sourceFileLocation"),
//   	Type: jsii.String("type"),
//   }
//
type CfnWorkflow_DecryptStepDetailsProperty struct {
	// `CfnWorkflow.DecryptStepDetailsProperty.DestinationFileLocation`.
	DestinationFileLocation interface{} `field:"optional" json:"destinationFileLocation" yaml:"destinationFileLocation"`
	// `CfnWorkflow.DecryptStepDetailsProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnWorkflow.DecryptStepDetailsProperty.OverwriteExisting`.
	OverwriteExisting *string `field:"optional" json:"overwriteExisting" yaml:"overwriteExisting"`
	// `CfnWorkflow.DecryptStepDetailsProperty.SourceFileLocation`.
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
	// `CfnWorkflow.DecryptStepDetailsProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

