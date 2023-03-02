package awsdatabrew


// Represents additional options for correct interpretation of datetime parameters used in the Amazon S3 path of a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datetimeOptionsProperty := &datetimeOptionsProperty{
//   	format: jsii.String("format"),
//
//   	// the properties below are optional
//   	localeCode: jsii.String("localeCode"),
//   	timezoneOffset: jsii.String("timezoneOffset"),
//   }
//
type CfnDataset_DatetimeOptionsProperty struct {
	// Required option, that defines the datetime format used for a date parameter in the Amazon S3 path.
	//
	// Should use only supported datetime specifiers and separation characters, all litera a-z or A-Z character should be escaped with single quotes. E.g. "MM.dd.yyyy-'at'-HH:mm".
	Format *string `field:"required" json:"format" yaml:"format"`
	// Optional value for a non-US locale code, needed for correct interpretation of some date formats.
	LocaleCode *string `field:"optional" json:"localeCode" yaml:"localeCode"`
	// Optional value for a timezone offset of the datetime parameter value in the Amazon S3 path.
	//
	// Shouldn't be used if Format for this parameter includes timezone fields. If no offset specified, UTC is assumed.
	TimezoneOffset *string `field:"optional" json:"timezoneOffset" yaml:"timezoneOffset"`
}

