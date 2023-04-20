package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gridLayoutConfigurationProperty := &GridLayoutConfigurationProperty{
//   	Elements: []interface{}{
//   		&GridLayoutElementProperty{
//   			ColumnSpan: jsii.Number(123),
//   			ElementId: jsii.String("elementId"),
//   			ElementType: jsii.String("elementType"),
//   			RowSpan: jsii.Number(123),
//
//   			// the properties below are optional
//   			ColumnIndex: jsii.Number(123),
//   			RowIndex: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	CanvasSizeOptions: &GridLayoutCanvasSizeOptionsProperty{
//   		ScreenCanvasSizeOptions: &GridLayoutScreenCanvasSizeOptionsProperty{
//   			ResizeOption: jsii.String("resizeOption"),
//
//   			// the properties below are optional
//   			OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   		},
//   	},
//   }
//
type CfnDashboard_GridLayoutConfigurationProperty struct {
	// `CfnDashboard.GridLayoutConfigurationProperty.Elements`.
	Elements interface{} `field:"required" json:"elements" yaml:"elements"`
	// `CfnDashboard.GridLayoutConfigurationProperty.CanvasSizeOptions`.
	CanvasSizeOptions interface{} `field:"optional" json:"canvasSizeOptions" yaml:"canvasSizeOptions"`
}

