package awslogs


// Properties for creating AWS vended log parsers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vendedLogParserProps := &VendedLogParserProps{
//   	LogType: awscdk.Aws_logs.VendedLogType_CLOUDFRONT,
//
//   	// the properties below are optional
//   	Source: jsii.String("source"),
//   }
//
type VendedLogParserProps struct {
	// The type of AWS vended log to parse.
	LogType VendedLogType `field:"required" json:"logType" yaml:"logType"`
	// Source field to parse.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

