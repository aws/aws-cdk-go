package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   colorsConfigurationProperty := &ColorsConfigurationProperty{
//   	CustomColors: []interface{}{
//   		&CustomColorProperty{
//   			Color: jsii.String("color"),
//
//   			// the properties below are optional
//   			FieldValue: jsii.String("fieldValue"),
//   			SpecialValue: jsii.String("specialValue"),
//   		},
//   	},
//   }
//
type CfnDashboard_ColorsConfigurationProperty struct {
	// `CfnDashboard.ColorsConfigurationProperty.CustomColors`.
	CustomColors interface{} `field:"optional" json:"customColors" yaml:"customColors"`
}

