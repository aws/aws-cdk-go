package previewawsobservabilityadminmixins


// CloudWatch Logs data source to associate with the S3 Table Integration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   logSourceProperty := &LogSourceProperty{
//   	Identifier: jsii.String("identifier"),
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-logsource.html
//
type CfnS3TableIntegrationPropsMixin_LogSourceProperty struct {
	// The ID of the CloudWatch Logs data source association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-logsource.html#cfn-observabilityadmin-s3tableintegration-logsource-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// The name of the CloudWatch Logs data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-logsource.html#cfn-observabilityadmin-s3tableintegration-logsource-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of the CloudWatch Logs data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-logsource.html#cfn-observabilityadmin-s3tableintegration-logsource-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

