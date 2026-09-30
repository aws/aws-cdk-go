package awsquicksight


// A visual that contains custom content.
//
// For more information, see [Using custom visual content](https://docs.aws.amazon.com/quicksight/latest/user/custom-visual-content.html) in the *Amazon Quick Suite User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customContentVisualProperty := &CustomContentVisualProperty{
//   	VisualId: jsii.String("visualId"),
//
//   	// the properties below are optional
//   	Actions: []interface{}{
//   		&VisualCustomActionProperty{
//   			ActionOperations: []interface{}{
//   				&VisualCustomActionOperationProperty{
//   					FilterOperation: &CustomActionFilterOperationProperty{
//   						SelectedFieldsConfiguration: &FilterOperationSelectedFieldsConfigurationProperty{
//   							SelectedColumns: []interface{}{
//   								&ColumnIdentifierProperty{
//   									ColumnName: jsii.String("columnName"),
//
//   									// the properties below are optional
//   									DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   									TopicIdentifier: jsii.String("topicIdentifier"),
//   								},
//   							},
//   							SelectedFieldOptions: jsii.String("selectedFieldOptions"),
//   							SelectedFields: []*string{
//   								jsii.String("selectedFields"),
//   							},
//   						},
//   						TargetVisualsConfiguration: &FilterOperationTargetVisualsConfigurationProperty{
//   							SameSheetTargetVisualConfiguration: &SameSheetTargetVisualConfigurationProperty{
//   								TargetVisualOptions: jsii.String("targetVisualOptions"),
//   								TargetVisuals: []*string{
//   									jsii.String("targetVisuals"),
//   								},
//   							},
//   						},
//   					},
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
//
//   										// the properties below are optional
//   										DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   										TopicIdentifier: jsii.String("topicIdentifier"),
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
//   	ChartConfiguration: &CustomContentConfigurationProperty{
//   		ContentType: jsii.String("contentType"),
//   		ContentUrl: jsii.String("contentUrl"),
//   		ImageScaling: jsii.String("imageScaling"),
//   		Interactions: &VisualInteractionOptionsProperty{
//   			ContextMenuOption: &ContextMenuOptionProperty{
//   				AvailabilityStatus: jsii.String("availabilityStatus"),
//   			},
//   			VisualMenuOption: &VisualMenuOptionProperty{
//   				AvailabilityStatus: jsii.String("availabilityStatus"),
//   			},
//   		},
//   	},
//   	DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	Subtitle: &VisualSubtitleLabelOptionsProperty{
//   		FormatText: &LongFormatTextProperty{
//   			PlainText: jsii.String("plainText"),
//   			RichText: jsii.String("richText"),
//   		},
//   		Visibility: jsii.String("visibility"),
//   	},
//   	Title: &VisualTitleLabelOptionsProperty{
//   		FormatText: &ShortFormatTextProperty{
//   			PlainText: jsii.String("plainText"),
//   			RichText: jsii.String("richText"),
//   		},
//   		Visibility: jsii.String("visibility"),
//   	},
//   	TopicIdentifier: jsii.String("topicIdentifier"),
//   	VisualContentAltText: jsii.String("visualContentAltText"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customcontentvisual.html
//
type CfnTemplate_CustomContentVisualProperty struct {
	// The unique identifier of a visual.
	//
	// This identifier must be unique within the context of a dashboard, template, or analysis. Two dashboards, analyses, or templates can have visuals with the same identifiers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customcontentvisual.html#cfn-quicksight-template-customcontentvisual-visualid
	//
	VisualId *string `field:"required" json:"visualId" yaml:"visualId"`
	// The list of custom actions that are configured for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customcontentvisual.html#cfn-quicksight-template-customcontentvisual-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The configuration of a `CustomContentVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customcontentvisual.html#cfn-quicksight-template-customcontentvisual-chartconfiguration
	//
	ChartConfiguration interface{} `field:"optional" json:"chartConfiguration" yaml:"chartConfiguration"`
	// The dataset that is used to create the custom content visual.
	//
	// You can't create a visual without a dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customcontentvisual.html#cfn-quicksight-template-customcontentvisual-datasetidentifier
	//
	// Default: - "".
	//
	DataSetIdentifier *string `field:"optional" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
	// The subtitle that is displayed on the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customcontentvisual.html#cfn-quicksight-template-customcontentvisual-subtitle
	//
	Subtitle interface{} `field:"optional" json:"subtitle" yaml:"subtitle"`
	// The title that is displayed on the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customcontentvisual.html#cfn-quicksight-template-customcontentvisual-title
	//
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customcontentvisual.html#cfn-quicksight-template-customcontentvisual-topicidentifier
	//
	TopicIdentifier *string `field:"optional" json:"topicIdentifier" yaml:"topicIdentifier"`
	// The alt text for the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customcontentvisual.html#cfn-quicksight-template-customcontentvisual-visualcontentalttext
	//
	VisualContentAltText *string `field:"optional" json:"visualContentAltText" yaml:"visualContentAltText"`
}

