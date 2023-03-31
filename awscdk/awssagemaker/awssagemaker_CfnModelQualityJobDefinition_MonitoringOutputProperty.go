package awssagemaker


// The output object for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringOutputProperty := &monitoringOutputProperty{
//   	s3Output: &s3OutputProperty{
//   		localPath: jsii.String("localPath"),
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		s3UploadMode: jsii.String("s3UploadMode"),
//   	},
//   }
//
type CfnModelQualityJobDefinition_MonitoringOutputProperty struct {
	// The Amazon S3 storage location where the results of a monitoring job are saved.
	S3Output interface{} `field:"required" json:"s3Output" yaml:"s3Output"`
}

