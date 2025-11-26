package previewawsapigatewayv2mixins


// Settings for logging access in a stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   accessLogSettingsProperty := &AccessLogSettingsProperty{
//   	DestinationArn: jsii.String("destinationArn"),
//   	Format: jsii.String("format"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-stage-accesslogsettings.html
//
type CfnStagePropsMixin_AccessLogSettingsProperty struct {
	// The ARN of the CloudWatch Logs log group to receive access logs.
	//
	// This parameter is required to enable access logging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-stage-accesslogsettings.html#cfn-apigatewayv2-stage-accesslogsettings-destinationarn
	//
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// A single line format of the access logs of data, as specified by selected $context variables.
	//
	// The format must include at least $context.requestId. This parameter is required to enable access logging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-stage-accesslogsettings.html#cfn-apigatewayv2-stage-accesslogsettings-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
}

