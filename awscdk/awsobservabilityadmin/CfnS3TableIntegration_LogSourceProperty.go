package awsobservabilityadmin


// CloudWatch Logs data source to associate with the S3 Table Integration.
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
	// The name of the CloudWatch Logs data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-logsource.html#cfn-observabilityadmin-s3tableintegration-logsource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the CloudWatch Logs data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-logsource.html#cfn-observabilityadmin-s3tableintegration-logsource-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The ID of the CloudWatch Logs data source association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-logsource.html#cfn-observabilityadmin-s3tableintegration-logsource-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
}

