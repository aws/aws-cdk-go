package awssagemaker


// The ground truth labels for the dataset used for the monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringGroundTruthS3InputProperty := &monitoringGroundTruthS3InputProperty{
//   	s3Uri: jsii.String("s3Uri"),
//   }
//
type CfnModelQualityJobDefinition_MonitoringGroundTruthS3InputProperty struct {
	// The address of the Amazon S3 location of the ground truth labels.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
}

