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
//   flowDefinitionProperty := &flowDefinitionProperty{
//   	flowName: jsii.String("flowName"),
//   	kmsArn: jsii.String("kmsArn"),
//   	sourceFlowConfig: &sourceFlowConfigProperty{
//   		connectorType: jsii.String("connectorType"),
//   		sourceConnectorProperties: &sourceConnectorPropertiesProperty{
//   			marketo: &marketoSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			s3: &s3SourcePropertiesProperty{
//   				bucketName: jsii.String("bucketName"),
//
//   				// the properties below are optional
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   			},
//   			salesforce: &salesforceSourcePropertiesProperty{
//   				object: jsii.String("object"),
//
//   				// the properties below are optional
//   				enableDynamicFieldUpdate: jsii.Boolean(false),
//   				includeDeletedRecords: jsii.Boolean(false),
//   			},
//   			serviceNow: &serviceNowSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			zendesk: &zendeskSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		connectorProfileName: jsii.String("connectorProfileName"),
//   		incrementalPullConfig: &incrementalPullConfigProperty{
//   			datetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   		},
//   	},
//   	tasks: []interface{}{
//   		&taskProperty{
//   			sourceFields: []*string{
//   				jsii.String("sourceFields"),
//   			},
//   			taskType: jsii.String("taskType"),
//
//   			// the properties below are optional
//   			connectorOperator: &connectorOperatorProperty{
//   				marketo: jsii.String("marketo"),
//   				s3: jsii.String("s3"),
//   				salesforce: jsii.String("salesforce"),
//   				serviceNow: jsii.String("serviceNow"),
//   				zendesk: jsii.String("zendesk"),
//   			},
//   			destinationField: jsii.String("destinationField"),
//   			taskProperties: []interface{}{
//   				&taskPropertiesMapProperty{
//   					operatorPropertyKey: jsii.String("operatorPropertyKey"),
//   					property: jsii.String("property"),
//   				},
//   			},
//   		},
//   	},
//   	triggerConfig: &triggerConfigProperty{
//   		triggerType: jsii.String("triggerType"),
//
//   		// the properties below are optional
//   		triggerProperties: &triggerPropertiesProperty{
//   			scheduled: &scheduledTriggerPropertiesProperty{
//   				scheduleExpression: jsii.String("scheduleExpression"),
//
//   				// the properties below are optional
//   				dataPullMode: jsii.String("dataPullMode"),
//   				firstExecutionFrom: jsii.Number(123),
//   				scheduleEndTime: jsii.Number(123),
//   				scheduleOffset: jsii.Number(123),
//   				scheduleStartTime: jsii.Number(123),
//   				timezone: jsii.String("timezone"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
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

