package awscustomerprofiles


// The configurations that control how Customer Profiles retrieves data from the source, Amazon AppFlow.
//
// Customer Profiles uses this information to create an AppFlow flow on behalf of customers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowDefinitionProperty := &FlowDefinitionProperty{
//   	FlowName: jsii.String("flowName"),
//   	KmsArn: jsii.String("kmsArn"),
//   	SourceFlowConfig: &SourceFlowConfigProperty{
//   		ConnectorType: jsii.String("connectorType"),
//   		SourceConnectorProperties: &SourceConnectorPropertiesProperty{
//   			Marketo: &MarketoSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			S3: &S3SourcePropertiesProperty{
//   				BucketName: jsii.String("bucketName"),
//
//   				// the properties below are optional
//   				BucketPrefix: jsii.String("bucketPrefix"),
//   			},
//   			Salesforce: &SalesforceSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//
//   				// the properties below are optional
//   				EnableDynamicFieldUpdate: jsii.Boolean(false),
//   				IncludeDeletedRecords: jsii.Boolean(false),
//   			},
//   			ServiceNow: &ServiceNowSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			Zendesk: &ZendeskSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		ConnectorProfileName: jsii.String("connectorProfileName"),
//   		IncrementalPullConfig: &IncrementalPullConfigProperty{
//   			DatetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   		},
//   	},
//   	Tasks: []interface{}{
//   		&TaskProperty{
//   			SourceFields: []*string{
//   				jsii.String("sourceFields"),
//   			},
//   			TaskType: jsii.String("taskType"),
//
//   			// the properties below are optional
//   			ConnectorOperator: &ConnectorOperatorProperty{
//   				Marketo: jsii.String("marketo"),
//   				S3: jsii.String("s3"),
//   				Salesforce: jsii.String("salesforce"),
//   				ServiceNow: jsii.String("serviceNow"),
//   				Zendesk: jsii.String("zendesk"),
//   			},
//   			DestinationField: jsii.String("destinationField"),
//   			TaskProperties: []interface{}{
//   				&TaskPropertiesMapProperty{
//   					OperatorPropertyKey: jsii.String("operatorPropertyKey"),
//   					Property: jsii.String("property"),
//   				},
//   			},
//   		},
//   	},
//   	TriggerConfig: &TriggerConfigProperty{
//   		TriggerType: jsii.String("triggerType"),
//
//   		// the properties below are optional
//   		TriggerProperties: &TriggerPropertiesProperty{
//   			Scheduled: &ScheduledTriggerPropertiesProperty{
//   				ScheduleExpression: jsii.String("scheduleExpression"),
//
//   				// the properties below are optional
//   				DataPullMode: jsii.String("dataPullMode"),
//   				FirstExecutionFrom: jsii.Number(123),
//   				ScheduleEndTime: jsii.Number(123),
//   				ScheduleOffset: jsii.Number(123),
//   				ScheduleStartTime: jsii.Number(123),
//   				Timezone: jsii.String("timezone"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
type CfnIntegration_FlowDefinitionProperty struct {
	// The specified name of the flow.
	//
	// Use underscores (_) or hyphens (-) only. Spaces are not allowed.
	FlowName *string `field:"required" json:"flowName" yaml:"flowName"`
	// The Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key you provide for encryption.
	KmsArn *string `field:"required" json:"kmsArn" yaml:"kmsArn"`
	// The configuration that controls how Customer Profiles retrieves data from the source.
	SourceFlowConfig interface{} `field:"required" json:"sourceFlowConfig" yaml:"sourceFlowConfig"`
	// A list of tasks that Customer Profiles performs while transferring the data in the flow run.
	Tasks interface{} `field:"required" json:"tasks" yaml:"tasks"`
	// The trigger settings that determine how and when the flow runs.
	TriggerConfig interface{} `field:"required" json:"triggerConfig" yaml:"triggerConfig"`
	// A description of the flow you want to create.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

