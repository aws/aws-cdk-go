package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultNewSheetConfigurationProperty := &DefaultNewSheetConfigurationProperty{
//   	InteractiveLayoutConfiguration: &DefaultInteractiveLayoutConfigurationProperty{
//   		FreeForm: &DefaultFreeFormLayoutConfigurationProperty{
//   			CanvasSizeOptions: &FreeFormLayoutCanvasSizeOptionsProperty{
//   				ScreenCanvasSizeOptions: &FreeFormLayoutScreenCanvasSizeOptionsProperty{
//   					OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   				},
//   			},
//   		},
//   		Grid: &DefaultGridLayoutConfigurationProperty{
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
//   	PaginatedLayoutConfiguration: &DefaultPaginatedLayoutConfigurationProperty{
//   		SectionBased: &DefaultSectionBasedLayoutConfigurationProperty{
//   			CanvasSizeOptions: &SectionBasedLayoutCanvasSizeOptionsProperty{
//   				PaperCanvasSizeOptions: &SectionBasedLayoutPaperCanvasSizeOptionsProperty{
//   					PaperMargin: &SpacingProperty{
//   						Bottom: jsii.String("bottom"),
//   						Left: jsii.String("left"),
//   						Right: jsii.String("right"),
//   						Top: jsii.String("top"),
//   					},
//   					PaperOrientation: jsii.String("paperOrientation"),
//   					PaperSize: jsii.String("paperSize"),
//   				},
//   			},
//   		},
//   	},
//   	SheetContentType: jsii.String("sheetContentType"),
//   }
//
type CfnAnalysis_DefaultNewSheetConfigurationProperty struct {
	// `CfnAnalysis.DefaultNewSheetConfigurationProperty.InteractiveLayoutConfiguration`.
	InteractiveLayoutConfiguration interface{} `field:"optional" json:"interactiveLayoutConfiguration" yaml:"interactiveLayoutConfiguration"`
	// `CfnAnalysis.DefaultNewSheetConfigurationProperty.PaginatedLayoutConfiguration`.
	PaginatedLayoutConfiguration interface{} `field:"optional" json:"paginatedLayoutConfiguration" yaml:"paginatedLayoutConfiguration"`
	// `CfnAnalysis.DefaultNewSheetConfigurationProperty.SheetContentType`.
	SheetContentType *string `field:"optional" json:"sheetContentType" yaml:"sheetContentType"`
}

