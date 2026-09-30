package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   geospatialNullDataSettingsProperty := &GeospatialNullDataSettingsProperty{
//   	SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   		FillColor: jsii.String("fillColor"),
//   		StrokeColor: jsii.String("strokeColor"),
//   		StrokeWidth: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialnulldatasettings.html
//
type CfnTemplatePropsMixin_GeospatialNullDataSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialnulldatasettings.html#cfn-quicksight-template-geospatialnulldatasettings-symbolstyle
	//
	SymbolStyle interface{} `field:"optional" json:"symbolStyle" yaml:"symbolStyle"`
}

