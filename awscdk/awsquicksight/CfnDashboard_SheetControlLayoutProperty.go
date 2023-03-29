package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetControlLayoutProperty := &SheetControlLayoutProperty{
//   	Configuration: &SheetControlLayoutConfigurationProperty{
//   		GridLayout: &GridLayoutConfigurationProperty{
//   			Elements: []interface{}{
//   				&GridLayoutElementProperty{
//   					ColumnSpan: jsii.Number(123),
//   					ElementId: jsii.String("elementId"),
//   					ElementType: jsii.String("elementType"),
//   					RowSpan: jsii.Number(123),
//
//   					// the properties below are optional
//   					ColumnIndex: jsii.Number(123),
//   					RowIndex: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			CanvasSizeOptions: &GridLayoutCanvasSizeOptionsProperty{
//   				ScreenCanvasSizeOptions: &GridLayoutScreenCanvasSizeOptionsProperty{
//   					ResizeOption: jsii.String("resizeOption"),
//
//   					// the properties below are optional
//   					OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDashboard_SheetControlLayoutProperty struct {
	// `CfnDashboard.SheetControlLayoutProperty.Configuration`.
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
}

