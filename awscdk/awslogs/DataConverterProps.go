package awslogs


// Properties for creating data converter processors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataConverterProps := &DataConverterProps{
//   	Type: awscdk.Aws_logs.DataConverterType_TYPE_CONVERTER,
//
//   	// the properties below are optional
//   	DateTimeConverterOptions: &DateTimeConverterProperty{
//   		Locale: jsii.String("locale"),
//   		MatchPatterns: []*string{
//   			jsii.String("matchPatterns"),
//   		},
//   		Source: jsii.String("source"),
//   		Target: jsii.String("target"),
//
//   		// the properties below are optional
//   		SourceTimezone: jsii.String("sourceTimezone"),
//   		TargetFormat: jsii.String("targetFormat"),
//   		TargetTimezone: jsii.String("targetTimezone"),
//   	},
//   	TypeConverterOptions: &TypeConverterProperty{
//   		Entries: []typeConverterEntryProperty{
//   			&typeConverterEntryProperty{
//   				Key: jsii.String("key"),
//   				Type: awscdk.*Aws_logs.TypeConverterType_BOOLEAN,
//   			},
//   		},
//   	},
//   }
//
type DataConverterProps struct {
	// The type of data conversion operation.
	Type DataConverterType `field:"required" json:"type" yaml:"type"`
	// Options for datetime conversion.
	//
	// Required when type is DATETIME_CONVERTER.
	// Default: - No date time converter processor is created if not set.
	//
	DateTimeConverterOptions *DateTimeConverterProperty `field:"optional" json:"dateTimeConverterOptions" yaml:"dateTimeConverterOptions"`
	// Options for type conversion.
	//
	// Required when type is TYPE_CONVERTER.
	// Default: - No type convertor processor is created if not set.
	//
	TypeConverterOptions *TypeConverterProperty `field:"optional" json:"typeConverterOptions" yaml:"typeConverterOptions"`
}

