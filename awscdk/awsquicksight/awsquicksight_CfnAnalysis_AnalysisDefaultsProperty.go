package awsquicksight


// The configuration for default analysis settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisDefaultsProperty := &AnalysisDefaultsProperty{
//   	DefaultNewSheetConfiguration: &DefaultNewSheetConfigurationProperty{
//   		InteractiveLayoutConfiguration: &DefaultInteractiveLayoutConfigurationProperty{
//   			FreeForm: &DefaultFreeFormLayoutConfigurationProperty{
//   				CanvasSizeOptions: &FreeFormLayoutCanvasSizeOptionsProperty{
//   					ScreenCanvasSizeOptions: &FreeFormLayoutScreenCanvasSizeOptionsProperty{
//   						OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   					},
//   				},
//   			},
//   			Grid: &DefaultGridLayoutConfigurationProperty{
//   				CanvasSizeOptions: &GridLayoutCanvasSizeOptionsProperty{
//   					ScreenCanvasSizeOptions: &GridLayoutScreenCanvasSizeOptionsProperty{
//   						ResizeOption: jsii.String("resizeOption"),
//
//   						// the properties below are optional
//   						OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   					},
//   				},
//   			},
//   		},
//   		PaginatedLayoutConfiguration: &DefaultPaginatedLayoutConfigurationProperty{
//   			SectionBased: &DefaultSectionBasedLayoutConfigurationProperty{
//   				CanvasSizeOptions: &SectionBasedLayoutCanvasSizeOptionsProperty{
//   					PaperCanvasSizeOptions: &SectionBasedLayoutPaperCanvasSizeOptionsProperty{
//   						PaperMargin: &SpacingProperty{
//   							Bottom: jsii.String("bottom"),
//   							Left: jsii.String("left"),
//   							Right: jsii.String("right"),
//   							Top: jsii.String("top"),
//   						},
//   						PaperOrientation: jsii.String("paperOrientation"),
//   						PaperSize: jsii.String("paperSize"),
//   					},
//   				},
//   			},
//   		},
//   		SheetContentType: jsii.String("sheetContentType"),
//   	},
//   }
//
type CfnAnalysis_AnalysisDefaultsProperty struct {
	// The configuration for default new sheet settings.
	DefaultNewSheetConfiguration interface{} `field:"required" json:"defaultNewSheetConfiguration" yaml:"defaultNewSheetConfiguration"`
}

