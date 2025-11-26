package previewawsquicksightmixins


// The options that determine the null value format configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   nullValueFormatConfigurationProperty := &NullValueFormatConfigurationProperty{
//   	NullString: jsii.String("nullString"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-nullvalueformatconfiguration.html
//
type CfnTemplatePropsMixin_NullValueFormatConfigurationProperty struct {
	// Determines the null string of null values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-nullvalueformatconfiguration.html#cfn-quicksight-template-nullvalueformatconfiguration-nullstring
	//
	NullString *string `field:"optional" json:"nullString" yaml:"nullString"`
}

