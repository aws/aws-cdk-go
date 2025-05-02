package awssagemaker


// The statistics resource for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statisticsResourceProperty := &StatisticsResourceProperty{
//   	S3Uri: jsii.String("s3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-statisticsresource.html
//
type CfnDataQualityJobDefinition_StatisticsResourceProperty struct {
	// The Amazon S3 URI for the statistics resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-statisticsresource.html#cfn-sagemaker-dataqualityjobdefinition-statisticsresource-s3uri
	//
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

