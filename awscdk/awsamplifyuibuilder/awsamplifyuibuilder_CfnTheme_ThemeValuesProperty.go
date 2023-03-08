package awsamplifyuibuilder


// The `ThemeValues` property specifies key-value pair that defines a property of a theme.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var themeValueProperty_ themeValueProperty
//
//   themeValuesProperty := &themeValuesProperty{
//   	key: jsii.String("key"),
//   	value: &themeValueProperty{
//   		children: []interface{}{
//   			&themeValuesProperty{
//   				key: jsii.String("key"),
//   				value: themeValueProperty_,
//   			},
//   		},
//   		value: jsii.String("value"),
//   	},
//   }
//
type CfnTheme_ThemeValuesProperty struct {
	// The name of the property.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the property.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

