package awsfis


// The S3 destination for the experiment report.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   experimentReportS3ConfigurationProperty := &ExperimentReportS3ConfigurationProperty{
//   	BucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimentreports3configuration.html
//
type CfnExperimentTemplate_ExperimentReportS3ConfigurationProperty struct {
	// The name of the S3 bucket where the experiment report will be stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimentreports3configuration.html#cfn-fis-experimenttemplate-experimentreports3configuration-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The prefix of the S3 bucket where the experiment report will be stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimentreports3configuration.html#cfn-fis-experimenttemplate-experimentreports3configuration-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

