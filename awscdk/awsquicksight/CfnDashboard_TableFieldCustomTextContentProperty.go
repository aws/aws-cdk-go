package awsquicksight


// The custom text content (value, font configuration) for the table link content configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableFieldCustomTextContentProperty := &TableFieldCustomTextContentProperty{
//   	FontConfiguration: &FontConfigurationProperty{
//   		FontColor: jsii.String("fontColor"),
//   		FontDecoration: jsii.String("fontDecoration"),
//   		FontSize: &FontSizeProperty{
//   			Relative: jsii.String("relative"),
//   		},
//   		FontStyle: jsii.String("fontStyle"),
//   		FontWeight: &FontWeightProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Value: jsii.String("value"),
//   }
//
type CfnDashboard_TableFieldCustomTextContentProperty struct {
	// The font configuration of the custom text content for the table URL link content.
	FontConfiguration interface{} `field:"required" json:"fontConfiguration" yaml:"fontConfiguration"`
	// The string value of the custom text content for the table URL link content.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

