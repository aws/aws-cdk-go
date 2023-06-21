package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIntegration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIntegrationProps := &CfnIntegrationProps{
//   	DomainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	FlowDefinition: &FlowDefinitionProperty{
//   		FlowName: jsii.String("flowName"),
//   		KmsArn: jsii.String("kmsArn"),
//   		SourceFlowConfig: &SourceFlowConfigProperty{
//   			ConnectorType: jsii.String("connectorType"),
//   			SourceConnectorProperties: &SourceConnectorPropertiesProperty{
//   				Marketo: &MarketoSourcePropertiesProperty{
//   					Object: jsii.String("object"),
//   				},
//   				S3: &S3SourcePropertiesProperty{
//   					BucketName: jsii.String("bucketName"),
//
//   					// the properties below are optional
//   					BucketPrefix: jsii.String("bucketPrefix"),
//   				},
//   				Salesforce: &SalesforceSourcePropertiesProperty{
//   					Object: jsii.String("object"),
//
//   					// the properties below are optional
//   					EnableDynamicFieldUpdate: jsii.Boolean(false),
//   					IncludeDeletedRecords: jsii.Boolean(false),
//   				},
//   				ServiceNow: &ServiceNowSourcePropertiesProperty{
//   					Object: jsii.String("object"),
//   				},
//   				Zendesk: &ZendeskSourcePropertiesProperty{
//   					Object: jsii.String("object"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			ConnectorProfileName: jsii.String("connectorProfileName"),
//   			IncrementalPullConfig: &IncrementalPullConfigProperty{
//   				DatetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   			},
//   		},
//   		Tasks: []interface{}{
//   			&TaskProperty{
//   				SourceFields: []*string{
//   					jsii.String("sourceFields"),
//   				},
//   				TaskType: jsii.String("taskType"),
//
//   				// the properties below are optional
//   				ConnectorOperator: &ConnectorOperatorProperty{
//   					Marketo: jsii.String("marketo"),
//   					S3: jsii.String("s3"),
//   					Salesforce: jsii.String("salesforce"),
//   					ServiceNow: jsii.String("serviceNow"),
//   					Zendesk: jsii.String("zendesk"),
//   				},
//   				DestinationField: jsii.String("destinationField"),
//   				TaskProperties: []interface{}{
//   					&TaskPropertiesMapProperty{
//   						OperatorPropertyKey: jsii.String("operatorPropertyKey"),
//   						Property: jsii.String("property"),
//   					},
//   				},
//   			},
//   		},
//   		TriggerConfig: &TriggerConfigProperty{
//   			TriggerType: jsii.String("triggerType"),
//
//   			// the properties below are optional
//   			TriggerProperties: &TriggerPropertiesProperty{
//   				Scheduled: &ScheduledTriggerPropertiesProperty{
//   					ScheduleExpression: jsii.String("scheduleExpression"),
//
//   					// the properties below are optional
//   					DataPullMode: jsii.String("dataPullMode"),
//   					FirstExecutionFrom: jsii.Number(123),
//   					ScheduleEndTime: jsii.Number(123),
//   					ScheduleOffset: jsii.Number(123),
//   					ScheduleStartTime: jsii.Number(123),
//   					Timezone: jsii.String("timezone"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   	},
//   	ObjectTypeName: jsii.String("objectTypeName"),
//   	ObjectTypeNames: []interface{}{
//   		&ObjectTypeMappingProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Uri: jsii.String("uri"),
//   }
//
type CfnIntegrationProps struct {
	// The unique name of the domain.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The configuration that controls how Customer Profiles retrieves data from the source.
	FlowDefinition interface{} `field:"optional" json:"flowDefinition" yaml:"flowDefinition"`
	// The name of the profile object type mapping to use.
	ObjectTypeName *string `field:"optional" json:"objectTypeName" yaml:"objectTypeName"`
	// The object type mapping.
	ObjectTypeNames interface{} `field:"optional" json:"objectTypeNames" yaml:"objectTypeNames"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The URI of the S3 bucket or any other type of data source.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

