package awssagemaker


// The baseline statistics file in Amazon S3 that the current monitoring job should be validated against.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statisticsResourceProperty := &statisticsResourceProperty{
//   	s3Uri: jsii.String("s3Uri"),
//   }
//
type CfnMonitoringSchedule_StatisticsResourceProperty struct {
	// The S3 URI for the statistics resource.
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

