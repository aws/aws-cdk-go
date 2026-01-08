package previewawsquicksightmixins


// The configuration that determines the elements and canvas size options of sheet control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sheetControlLayoutConfigurationProperty := &SheetControlLayoutConfigurationProperty{
//   	GridLayout: &GridLayoutConfigurationProperty{
//   		CanvasSizeOptions: &GridLayoutCanvasSizeOptionsProperty{
//   			ScreenCanvasSizeOptions: &GridLayoutScreenCanvasSizeOptionsProperty{
//   				OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   				ResizeOption: jsii.String("resizeOption"),
//   			},
//   		},
//   		Elements: []interface{}{
//   			&GridLayoutElementProperty{
//   				BackgroundStyle: &GridLayoutElementBackgroundStyleProperty{
//   					Color: jsii.String("color"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   				BorderRadius: jsii.String("borderRadius"),
//   				BorderStyle: &GridLayoutElementBorderStyleProperty{
//   					Color: jsii.String("color"),
//   					Visibility: jsii.String("visibility"),
//   					Width: jsii.String("width"),
//   				},
//   				ColumnIndex: jsii.Number(123),
//   				ColumnSpan: jsii.Number(123),
//   				ElementId: jsii.String("elementId"),
//   				ElementType: jsii.String("elementType"),
//   				LoadingAnimation: &LoadingAnimationProperty{
//   					Visibility: jsii.String("visibility"),
//   				},
//   				Padding: jsii.String("padding"),
//   				RowIndex: jsii.Number(123),
//   				RowSpan: jsii.Number(123),
//   				SelectedBorderStyle: &GridLayoutElementBorderStyleProperty{
//   					Color: jsii.String("color"),
//   					Visibility: jsii.String("visibility"),
//   					Width: jsii.String("width"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetcontrollayoutconfiguration.html
//
type CfnDashboardPropsMixin_SheetControlLayoutConfigurationProperty struct {
	// The configuration that determines the elements and canvas size options of sheet control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetcontrollayoutconfiguration.html#cfn-quicksight-dashboard-sheetcontrollayoutconfiguration-gridlayout
	//
	GridLayout interface{} `field:"optional" json:"gridLayout" yaml:"gridLayout"`
}

