package awslogs


// This processor uses pattern matching to parse and structure unstructured data.
//
// This processor can also extract fields from log messages.
// For more information about this processor including examples, see grok in the CloudWatch Logs User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grokProperty := &GrokProperty{
//   	Match: jsii.String("match"),
//
//   	// the properties below are optional
//   	Source: jsii.String("source"),
//   }
//
type GrokProperty struct {
	// The grok pattern to match against the log event.
	//
	// For a list of supported grok patterns,
	// see Supported grok patterns in the CloudWatch Logs User Guide.
	Match *string `field:"required" json:"match" yaml:"match"`
	// The path to the field in the log event that you want to parse.
	// Default: '@message'.
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

