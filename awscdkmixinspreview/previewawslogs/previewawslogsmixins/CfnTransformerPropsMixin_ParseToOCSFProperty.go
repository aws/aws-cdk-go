package previewawslogsmixins


// This processor converts logs into [Open Cybersecurity Schema Framework (OCSF)](https://docs.aws.amazon.com/https://ocsf.io) events.
//
// For more information about this processor including examples, see [parseToOCSF](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-parseToOCSF) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parseToOCSFProperty := &ParseToOCSFProperty{
//   	EventSource: jsii.String("eventSource"),
//   	MappingVersion: jsii.String("mappingVersion"),
//   	OcsfVersion: jsii.String("ocsfVersion"),
//   	Source: jsii.String("source"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsetoocsf.html
//
type CfnTransformerPropsMixin_ParseToOCSFProperty struct {
	// Specify the service or process that produces the log events that will be converted with this processor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsetoocsf.html#cfn-logs-transformer-parsetoocsf-eventsource
	//
	EventSource *string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsetoocsf.html#cfn-logs-transformer-parsetoocsf-mappingversion
	//
	MappingVersion *string `field:"optional" json:"mappingVersion" yaml:"mappingVersion"`
	// Specify which version of the OCSF schema to use for the transformed log events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsetoocsf.html#cfn-logs-transformer-parsetoocsf-ocsfversion
	//
	OcsfVersion *string `field:"optional" json:"ocsfVersion" yaml:"ocsfVersion"`
	// The path to the field in the log event that you want to parse.
	//
	// If you omit this value, the whole log message is parsed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsetoocsf.html#cfn-logs-transformer-parsetoocsf-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

