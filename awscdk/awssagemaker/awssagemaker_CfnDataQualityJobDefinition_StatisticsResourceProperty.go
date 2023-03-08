package awssagemaker


// The statistics resource for a monitoring job.
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
type CfnDataQualityJobDefinition_StatisticsResourceProperty struct {
	// The Amazon S3 URI for the statistics resource.
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

