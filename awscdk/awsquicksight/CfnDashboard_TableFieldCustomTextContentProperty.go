package awsquicksight


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
	// `CfnDashboard.TableFieldCustomTextContentProperty.FontConfiguration`.
	FontConfiguration interface{} `field:"required" json:"fontConfiguration" yaml:"fontConfiguration"`
	// `CfnDashboard.TableFieldCustomTextContentProperty.Value`.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

