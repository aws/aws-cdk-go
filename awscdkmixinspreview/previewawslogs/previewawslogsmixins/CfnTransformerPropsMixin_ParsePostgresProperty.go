package previewawslogsmixins


// Use this processor to parse RDS for PostgreSQL vended logs, extract fields, and and convert them into a JSON format.
//
// This processor always processes the entire log event message. For more information about this processor including examples, see [parsePostGres](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parsePostGres) .
//
// For more information about RDS for PostgreSQL log format, see [RDS for PostgreSQL database log filesTCP flag sequence](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.Concepts.PostgreSQL.html#USER_LogAccess.Concepts.PostgreSQL.Log_Format.log-line-prefix) .
//
// > If you use this processor, it must be the first processor in your transformer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parsePostgresProperty := &ParsePostgresProperty{
//   	Source: jsii.String("source"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsepostgres.html
//
type CfnTransformerPropsMixin_ParsePostgresProperty struct {
	// Omit this parameter and the whole log message will be processed by this processor.
	//
	// No other value than `@message` is allowed for `source` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsepostgres.html#cfn-logs-transformer-parsepostgres-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

