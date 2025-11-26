package previewawsquicksightmixins


// The configuration for default new sheet settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
//   					OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   					ResizeOption: jsii.String("resizeOption"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultnewsheetconfiguration.html
//
type CfnAnalysisPropsMixin_DefaultNewSheetConfigurationProperty struct {
	// The options that determine the default settings for interactive layout configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultnewsheetconfiguration.html#cfn-quicksight-analysis-defaultnewsheetconfiguration-interactivelayoutconfiguration
	//
	InteractiveLayoutConfiguration interface{} `field:"optional" json:"interactiveLayoutConfiguration" yaml:"interactiveLayoutConfiguration"`
	// The options that determine the default settings for a paginated layout configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultnewsheetconfiguration.html#cfn-quicksight-analysis-defaultnewsheetconfiguration-paginatedlayoutconfiguration
	//
	PaginatedLayoutConfiguration interface{} `field:"optional" json:"paginatedLayoutConfiguration" yaml:"paginatedLayoutConfiguration"`
	// The option that determines the sheet content type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultnewsheetconfiguration.html#cfn-quicksight-analysis-defaultnewsheetconfiguration-sheetcontenttype
	//
	SheetContentType *string `field:"optional" json:"sheetContentType" yaml:"sheetContentType"`
}

