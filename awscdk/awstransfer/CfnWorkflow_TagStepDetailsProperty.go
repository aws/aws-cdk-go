package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagStepDetailsProperty := &TagStepDetailsProperty{
//   	Name: jsii.String("name"),
//   	SourceFileLocation: jsii.String("sourceFileLocation"),
//   	Tags: []s3TagProperty{
//   		&s3TagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnWorkflow_TagStepDetailsProperty struct {
	// `CfnWorkflow.TagStepDetailsProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnWorkflow.TagStepDetailsProperty.SourceFileLocation`.
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
	// `CfnWorkflow.TagStepDetailsProperty.Tags`.
	Tags *[]*CfnWorkflow_S3TagProperty `field:"optional" json:"tags" yaml:"tags"`
}

