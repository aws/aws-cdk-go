package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   staticFileSourceProperty := &StaticFileSourceProperty{
//   	S3Options: &StaticFileS3SourceOptionsProperty{
//   		BucketName: jsii.String("bucketName"),
//   		ObjectKey: jsii.String("objectKey"),
//   		Region: jsii.String("region"),
//   	},
//   	UrlOptions: &StaticFileUrlSourceOptionsProperty{
//   		Url: jsii.String("url"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-staticfilesource.html
//
type CfnDashboard_StaticFileSourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-staticfilesource.html#cfn-quicksight-dashboard-staticfilesource-s3options
	//
	S3Options interface{} `field:"optional" json:"s3Options" yaml:"s3Options"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-staticfilesource.html#cfn-quicksight-dashboard-staticfilesource-urloptions
	//
	UrlOptions interface{} `field:"optional" json:"urlOptions" yaml:"urlOptions"`
}

