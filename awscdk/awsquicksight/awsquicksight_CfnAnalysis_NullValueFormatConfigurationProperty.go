package awsquicksight


// The options that determine the null value format configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nullValueFormatConfigurationProperty := &NullValueFormatConfigurationProperty{
//   	NullString: jsii.String("nullString"),
//   }
//
type CfnAnalysis_NullValueFormatConfigurationProperty struct {
	// Determines the null string of null values.
	NullString *string `field:"required" json:"nullString" yaml:"nullString"`
}

