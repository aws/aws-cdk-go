package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   decryptStepDetailsProperty := &decryptStepDetailsProperty{
//   	destinationFileLocation: &inputFileLocationProperty{
//   		efsFileLocation: &efsInputFileLocationProperty{
//   			fileSystemId: jsii.String("fileSystemId"),
//   			path: jsii.String("path"),
//   		},
//   		s3FileLocation: &s3InputFileLocationProperty{
//   			bucket: jsii.String("bucket"),
//   			key: jsii.String("key"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	overwriteExisting: jsii.String("overwriteExisting"),
//   	sourceFileLocation: jsii.String("sourceFileLocation"),
//   	type: jsii.String("type"),
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

