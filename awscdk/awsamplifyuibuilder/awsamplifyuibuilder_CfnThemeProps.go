package awsamplifyuibuilder


// Properties for defining a `CfnTheme`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var themeValuesProperty_ themeValuesProperty
//
//   cfnThemeProps := &cfnThemeProps{
//   	name: jsii.String("name"),
//   	values: []interface{}{
//   		&themeValuesProperty{
//   			key: jsii.String("key"),
//   			value: &themeValueProperty{
//   				children: []interface{}{
//   					themeValuesProperty_,
//   				},
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	overrides: []interface{}{
//   		&themeValuesProperty{
//   			key: jsii.String("key"),
//   			value: &themeValueProperty{
//   				children: []interface{}{
//   					themeValuesProperty_,
//   				},
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnThemeProps struct {
	// The name of the theme.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of key-value pairs that defines the properties of the theme.
	Values interface{} `field:"required" json:"values" yaml:"values"`
	// Describes the properties that can be overriden to customize a theme.
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// One or more key-value pairs to use when tagging the theme.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

