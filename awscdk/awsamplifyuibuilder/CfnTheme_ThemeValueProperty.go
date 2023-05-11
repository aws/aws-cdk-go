package awsamplifyuibuilder


// The `ThemeValue` property specifies the configuration of a theme's properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var themeValuesProperty_ themeValuesProperty
//
//   themeValueProperty := &ThemeValueProperty{
//   	Children: []interface{}{
//   		&themeValuesProperty{
//   			Key: jsii.String("key"),
//   			Value: &ThemeValueProperty{
//   				Children: []interface{}{
//   					themeValuesProperty_,
//   				},
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	Value: jsii.String("value"),
//   }
//
type CfnTheme_ThemeValueProperty struct {
	// A list of key-value pairs that define the theme's properties.
	Children interface{} `field:"optional" json:"children" yaml:"children"`
	// The value of a theme property.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

