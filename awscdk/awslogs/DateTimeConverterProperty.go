package awslogs


// This processor converts a datetime string into a format that you specify.
//
// For more information about this processor including examples, see datetimeConverter in the CloudWatch Logs User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateTimeConverterProperty := &DateTimeConverterProperty{
//   	Locale: jsii.String("locale"),
//   	MatchPatterns: []*string{
//   		jsii.String("matchPatterns"),
//   	},
//   	Source: jsii.String("source"),
//   	Target: jsii.String("target"),
//
//   	// the properties below are optional
//   	SourceTimezone: jsii.String("sourceTimezone"),
//   	TargetFormat: jsii.String("targetFormat"),
//   	TargetTimezone: jsii.String("targetTimezone"),
//   }
//
type DateTimeConverterProperty struct {
	// The locale of the source field.
	Locale *string `field:"required" json:"locale" yaml:"locale"`
	// A list of patterns to match against the source field.
	MatchPatterns *[]*string `field:"required" json:"matchPatterns" yaml:"matchPatterns"`
	// The key to apply the date conversion to.
	Source *string `field:"required" json:"source" yaml:"source"`
	// The JSON field to store the result in.
	Target *string `field:"required" json:"target" yaml:"target"`
	// The time zone of the source field.
	// Default: UTC.
	//
	SourceTimezone *string `field:"optional" json:"sourceTimezone" yaml:"sourceTimezone"`
	// The datetime format to use for the converted data in the target field.
	// Default: "yyyy-MM-dd'T'HH:mm:ss.SSS'Z'"
	//
	TargetFormat *string `field:"optional" json:"targetFormat" yaml:"targetFormat"`
	// The time zone of the target field.
	// Default: UTC.
	//
	TargetTimezone *string `field:"optional" json:"targetTimezone" yaml:"targetTimezone"`
}

