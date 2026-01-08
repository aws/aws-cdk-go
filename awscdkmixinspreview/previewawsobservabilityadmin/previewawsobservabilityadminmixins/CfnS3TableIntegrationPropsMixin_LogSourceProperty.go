package previewawsobservabilityadminmixins


// A data source with an S3 Table integration for query access in the `logs` namespace.
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
	// The unique identifier for the association between the data source and S3 Table integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-logsource.html#cfn-observabilityadmin-s3tableintegration-logsource-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// The name of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-logsource.html#cfn-observabilityadmin-s3tableintegration-logsource-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-logsource.html#cfn-observabilityadmin-s3tableintegration-logsource-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

