package mixinsawsquicksight


// A grid layout to define the placement of sheet control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sheetControlLayoutProperty := &SheetControlLayoutProperty{
//   	Configuration: &SheetControlLayoutConfigurationProperty{
//   		GridLayout: &GridLayoutConfigurationProperty{
//   			CanvasSizeOptions: &GridLayoutCanvasSizeOptionsProperty{
//   				ScreenCanvasSizeOptions: &GridLayoutScreenCanvasSizeOptionsProperty{
//   					OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   					ResizeOption: jsii.String("resizeOption"),
//   				},
//   			},
//   			Elements: []interface{}{
//   				&GridLayoutElementProperty{
//   					ColumnIndex: jsii.Number(123),
//   					ColumnSpan: jsii.Number(123),
//   					ElementId: jsii.String("elementId"),
//   					ElementType: jsii.String("elementType"),
//   					RowIndex: jsii.Number(123),
//   					RowSpan: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetcontrollayout.html
//
type CfnTemplatePropsMixin_SheetControlLayoutProperty struct {
	// The configuration that determines the elements and canvas size options of sheet control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetcontrollayout.html#cfn-quicksight-template-sheetcontrollayout-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
}

