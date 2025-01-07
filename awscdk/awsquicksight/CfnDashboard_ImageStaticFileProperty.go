package awsquicksight


// A static file that contains an image.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-imagestaticfile.html
//
type CfnDashboard_ImageStaticFileProperty struct {
	// The ID of the static file that contains an image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-imagestaticfile.html#cfn-quicksight-dashboard-imagestaticfile-staticfileid
	//
	StaticFileId *string `field:"required" json:"staticFileId" yaml:"staticFileId"`
	// The source of the image static file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-imagestaticfile.html#cfn-quicksight-dashboard-imagestaticfile-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

