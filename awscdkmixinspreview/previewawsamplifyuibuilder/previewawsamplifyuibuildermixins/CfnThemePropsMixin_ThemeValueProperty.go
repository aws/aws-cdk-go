package previewawsamplifyuibuildermixins


// The `ThemeValue` property specifies the configuration of a theme's properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var themeValueProperty_ ThemeValueProperty
//
//   themeValueProperty := &ThemeValueProperty{
//   	Children: []interface{}{
//   		&ThemeValuesProperty{
//   			Key: jsii.String("key"),
//   			Value: themeValueProperty_,
//   		},
//   	},
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-theme-themevalue.html
//
type CfnThemePropsMixin_ThemeValueProperty struct {
	// A list of key-value pairs that define the theme's properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-theme-themevalue.html#cfn-amplifyuibuilder-theme-themevalue-children
	//
	Children interface{} `field:"optional" json:"children" yaml:"children"`
	// The value of a theme property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-theme-themevalue.html#cfn-amplifyuibuilder-theme-themevalue-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

