package awsamplifyuibuilder


// The `ThemeValues` property specifies key-value pair that defines a property of a theme.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var themeValuesProperty_ themeValuesProperty
//
//   themeValuesProperty := &themeValuesProperty{
//   	Key: jsii.String("key"),
//   	Value: &ThemeValueProperty{
//   		Children: []interface{}{
//   			themeValuesProperty_,
//   		},
//   		Value: jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-theme-themevalues.html
//
type CfnTheme_ThemeValuesProperty struct {
	// The name of the property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-theme-themevalues.html#cfn-amplifyuibuilder-theme-themevalues-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-theme-themevalues.html#cfn-amplifyuibuilder-theme-themevalues-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

