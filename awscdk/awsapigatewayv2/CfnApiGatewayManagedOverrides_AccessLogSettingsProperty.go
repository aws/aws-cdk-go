package awsapigatewayv2


// The `AccessLogSettings` property overrides the access log settings for an API Gateway-managed stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLogSettingsProperty := &AccessLogSettingsProperty{
//   	DestinationArn: jsii.String("destinationArn"),
//   	Format: jsii.String("format"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-apigatewaymanagedoverrides-accesslogsettings.html
//
type CfnApiGatewayManagedOverrides_AccessLogSettingsProperty struct {
	// The ARN of the CloudWatch Logs log group to receive access logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-apigatewaymanagedoverrides-accesslogsettings.html#cfn-apigatewayv2-apigatewaymanagedoverrides-accesslogsettings-destinationarn
	//
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// A single line format of the access logs of data, as specified by selected $context variables.
	//
	// The format must include at least $context.requestId.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-apigatewaymanagedoverrides-accesslogsettings.html#cfn-apigatewayv2-apigatewaymanagedoverrides-accesslogsettings-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
}

