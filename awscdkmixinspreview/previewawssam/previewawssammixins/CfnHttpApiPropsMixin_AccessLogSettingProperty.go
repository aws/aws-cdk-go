package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   accessLogSettingProperty := &AccessLogSettingProperty{
//   	DestinationArn: jsii.String("destinationArn"),
//   	Format: jsii.String("format"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-accesslogsetting.html
//
type CfnHttpApiPropsMixin_AccessLogSettingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-accesslogsetting.html#cfn-serverless-httpapi-accesslogsetting-destinationarn
	//
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-accesslogsetting.html#cfn-serverless-httpapi-accesslogsetting-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
}

