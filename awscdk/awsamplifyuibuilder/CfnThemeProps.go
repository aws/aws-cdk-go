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
//   cfnThemeProps := &CfnThemeProps{
//   	Name: jsii.String("name"),
//   	Values: []interface{}{
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
//
//   	// the properties below are optional
//   	AppId: jsii.String("appId"),
//   	EnvironmentName: jsii.String("environmentName"),
//   	Overrides: []interface{}{
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
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-theme.html
//
type CfnThemeProps struct {
	// The name of the theme.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-theme.html#cfn-amplifyuibuilder-theme-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of key-value pairs that defines the properties of the theme.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-theme.html#cfn-amplifyuibuilder-theme-values
	//
	Values interface{} `field:"required" json:"values" yaml:"values"`
	// The unique ID for the Amplify app associated with the theme.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-theme.html#cfn-amplifyuibuilder-theme-appid
	//
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// The name of the backend environment that is a part of the Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-theme.html#cfn-amplifyuibuilder-theme-environmentname
	//
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// Describes the properties that can be overriden to customize a theme.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-theme.html#cfn-amplifyuibuilder-theme-overrides
	//
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// One or more key-value pairs to use when tagging the theme.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-theme.html#cfn-amplifyuibuilder-theme-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

