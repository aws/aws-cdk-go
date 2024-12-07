package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageStaticFileProperty := &ImageStaticFileProperty{
//   	StaticFileId: jsii.String("staticFileId"),
//
//   	// the properties below are optional
//   	Source: &StaticFileSourceProperty{
//   		S3Options: &StaticFileS3SourceOptionsProperty{
//   			BucketName: jsii.String("bucketName"),
//   			ObjectKey: jsii.String("objectKey"),
//   			Region: jsii.String("region"),
//   		},
//   		UrlOptions: &StaticFileUrlSourceOptionsProperty{
//   			Url: jsii.String("url"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-imagestaticfile.html
//
type CfnAnalysis_ImageStaticFileProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-imagestaticfile.html#cfn-quicksight-analysis-imagestaticfile-staticfileid
	//
	StaticFileId *string `field:"required" json:"staticFileId" yaml:"staticFileId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-imagestaticfile.html#cfn-quicksight-analysis-imagestaticfile-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

