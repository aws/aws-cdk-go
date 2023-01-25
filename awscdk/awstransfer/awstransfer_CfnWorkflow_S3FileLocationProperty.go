package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3FileLocationProperty := &s3FileLocationProperty{
//   	s3FileLocation: &s3InputFileLocationProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//   	},
//   }
//
type CfnWorkflow_S3FileLocationProperty struct {
	// `CfnWorkflow.S3FileLocationProperty.S3FileLocation`.
	S3FileLocation interface{} `field:"optional" json:"s3FileLocation" yaml:"s3FileLocation"`
}

