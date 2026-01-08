package awsobservabilityadmin


// A data source with an S3 Table integration for query access in the `logs` namespace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logSourceProperty := &LogSourceProperty{
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Identifier: jsii.String("identifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-logsource.html
//
type CfnS3TableIntegration_LogSourceProperty struct {
	// The name of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-logsource.html#cfn-observabilityadmin-s3tableintegration-logsource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-logsource.html#cfn-observabilityadmin-s3tableintegration-logsource-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The unique identifier for the association between the data source and S3 Table integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-logsource.html#cfn-observabilityadmin-s3tableintegration-logsource-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
}

