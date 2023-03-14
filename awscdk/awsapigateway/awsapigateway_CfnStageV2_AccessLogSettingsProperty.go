package awsapigateway


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLogSettingsProperty := &accessLogSettingsProperty{
//   	destinationArn: jsii.String("destinationArn"),
//   	format: jsii.String("format"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-stage-accesslogsettings.html
//
// Deprecated: moved to package aws-apigatewayv2.
type CfnStageV2_AccessLogSettingsProperty struct {
	// `CfnStageV2.AccessLogSettingsProperty.DestinationArn`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-stage-accesslogsettings.html#cfn-apigatewayv2-stage-accesslogsettings-destinationarn
	//
	// Deprecated: moved to package aws-apigatewayv2.
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// `CfnStageV2.AccessLogSettingsProperty.Format`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-stage-accesslogsettings.html#cfn-apigatewayv2-stage-accesslogsettings-format
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Format *string `field:"optional" json:"format" yaml:"format"`
}

