package awsquicksight


// An image that is located on a sheet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var availabilityStatus interface{}
//   var visibility interface{}
//
//   sheetImageProperty := &SheetImageProperty{
//   	SheetImageId: jsii.String("sheetImageId"),
//   	Source: &SheetImageSourceProperty{
//   		SheetImageStaticFileSource: &SheetImageStaticFileSourceProperty{
//   			StaticFileId: jsii.String("staticFileId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Actions: []interface{}{
//   		&ImageCustomActionProperty{
//   			ActionOperations: []interface{}{
//   				&ImageCustomActionOperationProperty{
//   					NavigationOperation: &CustomActionNavigationOperationProperty{
//   						LocalNavigationConfiguration: &LocalNavigationConfigurationProperty{
//   							TargetSheetId: jsii.String("targetSheetId"),
//   						},
//   					},
//   					SetParametersOperation: &CustomActionSetParametersOperationProperty{
//   						ParameterValueConfigurations: []interface{}{
//   							&SetParameterValueConfigurationProperty{
//   								DestinationParameterName: jsii.String("destinationParameterName"),
//   								Value: &DestinationParameterValueConfigurationProperty{
//   									CustomValuesConfiguration: &CustomValuesConfigurationProperty{
//   										CustomValues: &CustomParameterValuesProperty{
//   											DateTimeValues: []*string{
//   												jsii.String("dateTimeValues"),
//   											},
//   											DecimalValues: []interface{}{
//   												jsii.Number(123),
//   											},
//   											IntegerValues: []interface{}{
//   												jsii.Number(123),
//   											},
//   											StringValues: []*string{
//   												jsii.String("stringValues"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										IncludeNullValue: jsii.Boolean(false),
//   									},
//   									SelectAllValueOptions: jsii.String("selectAllValueOptions"),
//   									SourceColumn: &ColumnIdentifierProperty{
//   										ColumnName: jsii.String("columnName"),
//   										DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   									},
//   									SourceField: jsii.String("sourceField"),
//   									SourceParameterName: jsii.String("sourceParameterName"),
//   								},
//   							},
//   						},
//   					},
//   					UrlOperation: &CustomActionURLOperationProperty{
//   						UrlTarget: jsii.String("urlTarget"),
//   						UrlTemplate: jsii.String("urlTemplate"),
//   					},
//   				},
//   			},
//   			CustomActionId: jsii.String("customActionId"),
//   			Name: jsii.String("name"),
//   			Trigger: jsii.String("trigger"),
//
//   			// the properties below are optional
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	ImageContentAltText: jsii.String("imageContentAltText"),
//   	Interactions: &ImageInteractionOptionsProperty{
//   		ImageMenuOption: &ImageMenuOptionProperty{
//   			AvailabilityStatus: availabilityStatus,
//   		},
//   	},
//   	Scaling: &SheetImageScalingConfigurationProperty{
//   		ScalingType: jsii.String("scalingType"),
//   	},
//   	Tooltip: &SheetImageTooltipConfigurationProperty{
//   		TooltipText: &SheetImageTooltipTextProperty{
//   			PlainText: jsii.String("plainText"),
//   		},
//   		Visibility: visibility,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimage.html
//
type CfnTemplate_SheetImageProperty struct {
	// The ID of the sheet image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimage.html#cfn-quicksight-template-sheetimage-sheetimageid
	//
	SheetImageId *string `field:"required" json:"sheetImageId" yaml:"sheetImageId"`
	// The source of the image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimage.html#cfn-quicksight-template-sheetimage-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// A list of custom actions that are configured for an image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimage.html#cfn-quicksight-template-sheetimage-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The alt text for the image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimage.html#cfn-quicksight-template-sheetimage-imagecontentalttext
	//
	ImageContentAltText *string `field:"optional" json:"imageContentAltText" yaml:"imageContentAltText"`
	// The general image interactions setup for an image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimage.html#cfn-quicksight-template-sheetimage-interactions
	//
	Interactions interface{} `field:"optional" json:"interactions" yaml:"interactions"`
	// Determines how the image is scaled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimage.html#cfn-quicksight-template-sheetimage-scaling
	//
	Scaling interface{} `field:"optional" json:"scaling" yaml:"scaling"`
	// The tooltip to be shown when hovering over the image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimage.html#cfn-quicksight-template-sheetimage-tooltip
	//
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
}

