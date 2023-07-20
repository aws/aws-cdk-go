package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLogSettingProperty := &AccessLogSettingProperty{
//   	DestinationArn: jsii.String("destinationArn"),
//   	Format: jsii.String("format"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-accesslogsetting.html
//
type CfnApi_AccessLogSettingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-accesslogsetting.html#cfn-serverless-api-accesslogsetting-destinationarn
	//
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-accesslogsetting.html#cfn-serverless-api-accesslogsetting-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
}

