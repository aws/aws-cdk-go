package awskendra


// The value of a document attribute.
//
// You can only provide one value for a document attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentAttributeValueProperty := &documentAttributeValueProperty{
//   	dateValue: jsii.String("dateValue"),
//   	longValue: jsii.Number(123),
//   	stringListValue: []*string{
//   		jsii.String("stringListValue"),
//   	},
//   	stringValue: jsii.String("stringValue"),
//   }
//
type CfnDataSource_DocumentAttributeValueProperty struct {
	// A date expressed as an ISO 8601 string.
	//
	// It is important for the time zone to be included in the ISO 8601 date-time format. For example, 2012-03-25T12:30:10+01:00 is the ISO 8601 date-time format for March 25th 2012 at 12:30PM (plus 10 seconds) in Central European Time.
	DateValue *string `field:"optional" json:"dateValue" yaml:"dateValue"`
	// A long integer value.
	LongValue *float64 `field:"optional" json:"longValue" yaml:"longValue"`
	// A list of strings.
	StringListValue *[]*string `field:"optional" json:"stringListValue" yaml:"stringListValue"`
	// A string, such as "department".
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

