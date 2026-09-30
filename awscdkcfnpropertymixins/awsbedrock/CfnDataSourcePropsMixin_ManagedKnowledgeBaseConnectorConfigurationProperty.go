package awsbedrock


// Configuration for managed knowledge base connector data sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var connectorParameters interface{}
//   var daily interface{}
//   var lastDayOfMonth interface{}
//
//   managedKnowledgeBaseConnectorConfigurationProperty := &ManagedKnowledgeBaseConnectorConfigurationProperty{
//   	ConnectorParameters: connectorParameters,
//   	DeletionProtectionConfiguration: &DeletionProtectionConfigurationProperty{
//   		DeletionProtectionStatus: jsii.String("deletionProtectionStatus"),
//   		DeletionProtectionThreshold: jsii.Number(123),
//   	},
//   	MediaExtractionConfiguration: &MediaExtractionConfigurationProperty{
//   		AudioExtractionConfiguration: &AudioExtractionConfigurationProperty{
//   			AudioExtractionStatus: jsii.String("audioExtractionStatus"),
//   		},
//   		ImageExtractionConfiguration: &ImageExtractionConfigurationProperty{
//   			ImageExtractionStatus: jsii.String("imageExtractionStatus"),
//   		},
//   		VideoExtractionConfiguration: &VideoExtractionConfigurationProperty{
//   			VideoExtractionStatus: jsii.String("videoExtractionStatus"),
//   		},
//   	},
//   	SyncSchedule: &SyncScheduleProperty{
//   		Daily: daily,
//   		Monthly: &MonthlyScheduleProperty{
//   			DayOfMonth: &DayOfMonthProperty{
//   				DayNumber: jsii.Number(123),
//   				LastDayOfMonth: lastDayOfMonth,
//   			},
//   		},
//   		Weekly: &WeeklyScheduleProperty{
//   			DayOfWeek: jsii.String("dayOfWeek"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-managedknowledgebaseconnectorconfiguration.html
//
type CfnDataSourcePropsMixin_ManagedKnowledgeBaseConnectorConfigurationProperty struct {
	// Connector-specific parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-managedknowledgebaseconnectorconfiguration.html#cfn-bedrock-datasource-managedknowledgebaseconnectorconfiguration-connectorparameters
	//
	ConnectorParameters interface{} `field:"optional" json:"connectorParameters" yaml:"connectorParameters"`
	// Configuration for deletion protection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-managedknowledgebaseconnectorconfiguration.html#cfn-bedrock-datasource-managedknowledgebaseconnectorconfiguration-deletionprotectionconfiguration
	//
	DeletionProtectionConfiguration interface{} `field:"optional" json:"deletionProtectionConfiguration" yaml:"deletionProtectionConfiguration"`
	// Configuration for media extraction settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-managedknowledgebaseconnectorconfiguration.html#cfn-bedrock-datasource-managedknowledgebaseconnectorconfiguration-mediaextractionconfiguration
	//
	MediaExtractionConfiguration interface{} `field:"optional" json:"mediaExtractionConfiguration" yaml:"mediaExtractionConfiguration"`
	// Recurring schedule on which the connector automatically refreshes ingested content.
	//
	// Exactly one frequency variant is set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-managedknowledgebaseconnectorconfiguration.html#cfn-bedrock-datasource-managedknowledgebaseconnectorconfiguration-syncschedule
	//
	SyncSchedule interface{} `field:"optional" json:"syncSchedule" yaml:"syncSchedule"`
}

