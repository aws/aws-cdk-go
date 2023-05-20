package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for enabling Continuous Logging for Glue Jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var logGroup logGroup
//
//   continuousLoggingProps := &ContinuousLoggingProps{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	ConversionPattern: jsii.String("conversionPattern"),
//   	LogGroup: logGroup,
//   	LogStreamPrefix: jsii.String("logStreamPrefix"),
//   	Quiet: jsii.Boolean(false),
//   }
//
// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
//
// Experimental.
type ContinuousLoggingProps struct {
	// Enable continouous logging.
	// Experimental.
	Enabled *bool `field:"required" json:"enabled" yaml:"enabled"`
	// Apply the provided conversion pattern.
	//
	// This is a Log4j Conversion Pattern to customize driver and executor logs.
	// Experimental.
	ConversionPattern *string `field:"optional" json:"conversionPattern" yaml:"conversionPattern"`
	// Specify a custom CloudWatch log group name.
	// Experimental.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Specify a custom CloudWatch log stream prefix.
	// Experimental.
	LogStreamPrefix *string `field:"optional" json:"logStreamPrefix" yaml:"logStreamPrefix"`
	// Filter out non-useful Apache Spark driver/executor and Apache Hadoop YARN heartbeat log messages.
	// Experimental.
	Quiet *bool `field:"optional" json:"quiet" yaml:"quiet"`
}

