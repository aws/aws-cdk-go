package previewawsqbusinessmixins


// The value of a document attribute.
//
// You can only provide one value for a document attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   documentAttributeValueProperty := &DocumentAttributeValueProperty{
//   	DateValue: jsii.String("dateValue"),
//   	LongValue: jsii.Number(123),
//   	StringListValue: []*string{
//   		jsii.String("stringListValue"),
//   	},
//   	StringValue: jsii.String("stringValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-documentattributevalue.html
//
type CfnDataAccessorPropsMixin_DocumentAttributeValueProperty struct {
	// A date expressed as an ISO 8601 string.
	//
	// It's important for the time zone to be included in the ISO 8601 date-time format. For example, 2012-03-25T12:30:10+01:00 is the ISO 8601 date-time format for March 25th 2012 at 12:30PM (plus 10 seconds) in Central European Time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-documentattributevalue.html#cfn-qbusiness-dataaccessor-documentattributevalue-datevalue
	//
	DateValue *string `field:"optional" json:"dateValue" yaml:"dateValue"`
	// A long integer value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-documentattributevalue.html#cfn-qbusiness-dataaccessor-documentattributevalue-longvalue
	//
	LongValue *float64 `field:"optional" json:"longValue" yaml:"longValue"`
	// A list of strings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-documentattributevalue.html#cfn-qbusiness-dataaccessor-documentattributevalue-stringlistvalue
	//
	StringListValue *[]*string `field:"optional" json:"stringListValue" yaml:"stringListValue"`
	// A string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-documentattributevalue.html#cfn-qbusiness-dataaccessor-documentattributevalue-stringvalue
	//
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

