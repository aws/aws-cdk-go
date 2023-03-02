package awsapigatewayv2


// The `AccessLogSettings` property overrides the access log settings for an API Gateway-managed stage.
//
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
type CfnApiGatewayManagedOverrides_AccessLogSettingsProperty struct {
	// The ARN of the CloudWatch Logs log group to receive access logs.
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// A single line format of the access logs of data, as specified by selected $context variables.
	//
	// The format must include at least $context.requestId.
	Format *string `field:"optional" json:"format" yaml:"format"`
}

