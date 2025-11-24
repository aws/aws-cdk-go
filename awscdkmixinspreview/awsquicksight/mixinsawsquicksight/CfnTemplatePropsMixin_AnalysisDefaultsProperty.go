package mixinsawsquicksight


// The configuration for default analysis settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
//   						OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   						ResizeOption: jsii.String("resizeOption"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-analysisdefaults.html
//
type CfnTemplatePropsMixin_AnalysisDefaultsProperty struct {
	// The configuration for default new sheet settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-analysisdefaults.html#cfn-quicksight-template-analysisdefaults-defaultnewsheetconfiguration
	//
	DefaultNewSheetConfiguration interface{} `field:"optional" json:"defaultNewSheetConfiguration" yaml:"defaultNewSheetConfiguration"`
}

