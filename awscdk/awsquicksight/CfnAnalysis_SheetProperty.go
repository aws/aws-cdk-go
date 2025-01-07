package awsquicksight


// A *sheet* , which is an object that contains a set of visuals that are viewed together on one page in Amazon QuickSight.
//
// Every analysis and dashboard contains at least one sheet. Each sheet contains at least one visualization widget, for example a chart, pivot table, or narrative insight. Sheets can be associated with other components, such as controls, filters, and so on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var availabilityStatus interface{}
//
//   sheetProperty := &SheetProperty{
//   	Images: []interface{}{
//   		&SheetImageProperty{
//   			SheetImageId: jsii.String("sheetImageId"),
//   			Source: &SheetImageSourceProperty{
//   				SheetImageStaticFileSource: &SheetImageStaticFileSourceProperty{
//   					StaticFileId: jsii.String("staticFileId"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Actions: []interface{}{
//   				&ImageCustomActionProperty{
//   					ActionOperations: []interface{}{
//   						&ImageCustomActionOperationProperty{
//   							NavigationOperation: &CustomActionNavigationOperationProperty{
//   								LocalNavigationConfiguration: &LocalNavigationConfigurationProperty{
//   									TargetSheetId: jsii.String("targetSheetId"),
//   								},
//   							},
//   							SetParametersOperation: &CustomActionSetParametersOperationProperty{
//   								ParameterValueConfigurations: []interface{}{
//   									&SetParameterValueConfigurationProperty{
//   										DestinationParameterName: jsii.String("destinationParameterName"),
//   										Value: &DestinationParameterValueConfigurationProperty{
//   											CustomValuesConfiguration: &CustomValuesConfigurationProperty{
//   												CustomValues: &CustomParameterValuesProperty{
//   													DateTimeValues: []*string{
//   														jsii.String("dateTimeValues"),
//   													},
//   													DecimalValues: []interface{}{
//   														jsii.Number(123),
//   													},
//   													IntegerValues: []interface{}{
//   														jsii.Number(123),
//   													},
//   													StringValues: []*string{
//   														jsii.String("stringValues"),
//   													},
//   												},
//
//   												// the properties below are optional
//   												IncludeNullValue: jsii.Boolean(false),
//   											},
//   											SelectAllValueOptions: jsii.String("selectAllValueOptions"),
//   											SourceColumn: &ColumnIdentifierProperty{
//   												ColumnName: jsii.String("columnName"),
//   												DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   											},
//   											SourceField: jsii.String("sourceField"),
//   											SourceParameterName: jsii.String("sourceParameterName"),
//   										},
//   									},
//   								},
//   							},
//   							UrlOperation: &CustomActionURLOperationProperty{
//   								UrlTarget: jsii.String("urlTarget"),
//   								UrlTemplate: jsii.String("urlTemplate"),
//   							},
//   						},
//   					},
//   					CustomActionId: jsii.String("customActionId"),
//   					Name: jsii.String("name"),
//   					Trigger: jsii.String("trigger"),
//
//   					// the properties below are optional
//   					Status: jsii.String("status"),
//   				},
//   			},
//   			ImageContentAltText: jsii.String("imageContentAltText"),
//   			Interactions: &ImageInteractionOptionsProperty{
//   				ImageMenuOption: &ImageMenuOptionProperty{
//   					AvailabilityStatus: availabilityStatus,
//   				},
//   			},
//   			Scaling: &SheetImageScalingConfigurationProperty{
//   				ScalingType: jsii.String("scalingType"),
//   			},
//   			Tooltip: &SheetImageTooltipConfigurationProperty{
//   				TooltipText: &SheetImageTooltipTextProperty{
//   					PlainText: jsii.String("plainText"),
//   				},
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	SheetId: jsii.String("sheetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheet.html
//
type CfnAnalysis_SheetProperty struct {
	// A list of images on a sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheet.html#cfn-quicksight-analysis-sheet-images
	//
	Images interface{} `field:"optional" json:"images" yaml:"images"`
	// The name of a sheet.
	//
	// This name is displayed on the sheet's tab in the Amazon QuickSight console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheet.html#cfn-quicksight-analysis-sheet-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The unique identifier associated with a sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheet.html#cfn-quicksight-analysis-sheet-sheetid
	//
	SheetId *string `field:"optional" json:"sheetId" yaml:"sheetId"`
}

