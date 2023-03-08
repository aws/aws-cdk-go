package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagStepDetailsProperty := &tagStepDetailsProperty{
//   	name: jsii.String("name"),
//   	sourceFileLocation: jsii.String("sourceFileLocation"),
//   	tags: []s3TagProperty{
//   		&s3TagProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

