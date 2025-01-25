package awslogs


// This processor converts a datetime string into a format that you specify.
//
// For more information about this processor including examples, see [datetimeConverter](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-datetimeConverter) in the *CloudWatch Logs User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateTimeConverterProperty := &DateTimeConverterProperty{
//   	MatchPatterns: []*string{
//   		jsii.String("matchPatterns"),
//   	},
//   	Source: jsii.String("source"),
//   	Target: jsii.String("target"),
//
//   	// the properties below are optional
//   	Locale: jsii.String("locale"),
//   	SourceTimezone: jsii.String("sourceTimezone"),
//   	TargetFormat: jsii.String("targetFormat"),
//   	TargetTimezone: jsii.String("targetTimezone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-datetimeconverter.html
//
type CfnTransformer_DateTimeConverterProperty struct {
	// A list of patterns to match against the `source` field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-datetimeconverter.html#cfn-logs-transformer-datetimeconverter-matchpatterns
	//
	MatchPatterns *[]*string `field:"required" json:"matchPatterns" yaml:"matchPatterns"`
	// The key to apply the date conversion to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-datetimeconverter.html#cfn-logs-transformer-datetimeconverter-source
	//
	Source *string `field:"required" json:"source" yaml:"source"`
	// The JSON field to store the result in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-datetimeconverter.html#cfn-logs-transformer-datetimeconverter-target
	//
	Target *string `field:"required" json:"target" yaml:"target"`
	// The locale of the source field.
	//
	// If you omit this, the default of `locale.ROOT` is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-datetimeconverter.html#cfn-logs-transformer-datetimeconverter-locale
	//
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// The time zone of the source field.
	//
	// If you omit this, the default used is the UTC zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-datetimeconverter.html#cfn-logs-transformer-datetimeconverter-sourcetimezone
	//
	SourceTimezone *string `field:"optional" json:"sourceTimezone" yaml:"sourceTimezone"`
	// The datetime format to use for the converted data in the target field.
	//
	// If you omit this, the default of `yyyy-MM-dd'T'HH:mm:ss.SSS'Z` is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-datetimeconverter.html#cfn-logs-transformer-datetimeconverter-targetformat
	//
	TargetFormat *string `field:"optional" json:"targetFormat" yaml:"targetFormat"`
	// The time zone of the target field.
	//
	// If you omit this, the default used is the UTC zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-datetimeconverter.html#cfn-logs-transformer-datetimeconverter-targettimezone
	//
	TargetTimezone *string `field:"optional" json:"targetTimezone" yaml:"targetTimezone"`
}

