package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   	},
//   	Scaling: &SheetImageScalingConfigurationProperty{
//   		ScalingType: jsii.String("scalingType"),
//   	},
//   	Tooltip: &SheetImageTooltipConfigurationProperty{
//   		TooltipText: &SheetImageTooltipTextProperty{
//   			PlainText: jsii.String("plainText"),
//   		},
//   		Visibility: jsii.String("visibility"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetimage.html
//
type CfnAnalysis_SheetImageProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetimage.html#cfn-quicksight-analysis-sheetimage-sheetimageid
	//
	SheetImageId *string `field:"required" json:"sheetImageId" yaml:"sheetImageId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetimage.html#cfn-quicksight-analysis-sheetimage-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetimage.html#cfn-quicksight-analysis-sheetimage-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetimage.html#cfn-quicksight-analysis-sheetimage-imagecontentalttext
	//
	ImageContentAltText *string `field:"optional" json:"imageContentAltText" yaml:"imageContentAltText"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetimage.html#cfn-quicksight-analysis-sheetimage-interactions
	//
	Interactions interface{} `field:"optional" json:"interactions" yaml:"interactions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetimage.html#cfn-quicksight-analysis-sheetimage-scaling
	//
	Scaling interface{} `field:"optional" json:"scaling" yaml:"scaling"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetimage.html#cfn-quicksight-analysis-sheetimage-tooltip
	//
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
}

