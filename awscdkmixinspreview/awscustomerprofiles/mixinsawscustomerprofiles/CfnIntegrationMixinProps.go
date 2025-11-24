package mixinsawscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnIntegrationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIntegrationMixinProps := &CfnIntegrationMixinProps{
//   	DomainName: jsii.String("domainName"),
//   	EventTriggerNames: []*string{
//   		jsii.String("eventTriggerNames"),
//   	},
//   	FlowDefinition: &FlowDefinitionProperty{
//   		Description: jsii.String("description"),
//   		FlowName: jsii.String("flowName"),
//   		KmsArn: jsii.String("kmsArn"),
//   		SourceFlowConfig: &SourceFlowConfigProperty{
//   			ConnectorProfileName: jsii.String("connectorProfileName"),
//   			ConnectorType: jsii.String("connectorType"),
//   			IncrementalPullConfig: &IncrementalPullConfigProperty{
//   				DatetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   			},
//   			SourceConnectorProperties: &SourceConnectorPropertiesProperty{
//   				Marketo: &MarketoSourcePropertiesProperty{
//   					Object: jsii.String("object"),
//   				},
//   				S3: &S3SourcePropertiesProperty{
//   					BucketName: jsii.String("bucketName"),
//   					BucketPrefix: jsii.String("bucketPrefix"),
//   				},
//   				Salesforce: &SalesforceSourcePropertiesProperty{
//   					EnableDynamicFieldUpdate: jsii.Boolean(false),
//   					IncludeDeletedRecords: jsii.Boolean(false),
//   					Object: jsii.String("object"),
//   				},
//   				ServiceNow: &ServiceNowSourcePropertiesProperty{
//   					Object: jsii.String("object"),
//   				},
//   				Zendesk: &ZendeskSourcePropertiesProperty{
//   					Object: jsii.String("object"),
//   				},
//   			},
//   		},
//   		Tasks: []interface{}{
//   			&TaskProperty{
//   				ConnectorOperator: &ConnectorOperatorProperty{
//   					Marketo: jsii.String("marketo"),
//   					S3: jsii.String("s3"),
//   					Salesforce: jsii.String("salesforce"),
//   					ServiceNow: jsii.String("serviceNow"),
//   					Zendesk: jsii.String("zendesk"),
//   				},
//   				DestinationField: jsii.String("destinationField"),
//   				SourceFields: []*string{
//   					jsii.String("sourceFields"),
//   				},
//   				TaskProperties: []interface{}{
//   					&TaskPropertiesMapProperty{
//   						OperatorPropertyKey: jsii.String("operatorPropertyKey"),
//   						Property: jsii.String("property"),
//   					},
//   				},
//   				TaskType: jsii.String("taskType"),
//   			},
//   		},
//   		TriggerConfig: &TriggerConfigProperty{
//   			TriggerProperties: &TriggerPropertiesProperty{
//   				Scheduled: &ScheduledTriggerPropertiesProperty{
//   					DataPullMode: jsii.String("dataPullMode"),
//   					FirstExecutionFrom: jsii.Number(123),
//   					ScheduleEndTime: jsii.Number(123),
//   					ScheduleExpression: jsii.String("scheduleExpression"),
//   					ScheduleOffset: jsii.Number(123),
//   					ScheduleStartTime: jsii.Number(123),
//   					Timezone: jsii.String("timezone"),
//   				},
//   			},
//   			TriggerType: jsii.String("triggerType"),
//   		},
//   	},
//   	ObjectTypeName: jsii.String("objectTypeName"),
//   	ObjectTypeNames: []interface{}{
//   		&ObjectTypeMappingProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Uri: jsii.String("uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-integration.html
//
type CfnIntegrationMixinProps struct {
	// The unique name of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-integration.html#cfn-customerprofiles-integration-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// A list of unique names for active event triggers associated with the integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-integration.html#cfn-customerprofiles-integration-eventtriggernames
	//
	EventTriggerNames *[]*string `field:"optional" json:"eventTriggerNames" yaml:"eventTriggerNames"`
	// The configuration that controls how Customer Profiles retrieves data from the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-integration.html#cfn-customerprofiles-integration-flowdefinition
	//
	FlowDefinition interface{} `field:"optional" json:"flowDefinition" yaml:"flowDefinition"`
	// The name of the profile object type mapping to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-integration.html#cfn-customerprofiles-integration-objecttypename
	//
	ObjectTypeName *string `field:"optional" json:"objectTypeName" yaml:"objectTypeName"`
	// The object type mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-integration.html#cfn-customerprofiles-integration-objecttypenames
	//
	ObjectTypeNames interface{} `field:"optional" json:"objectTypeNames" yaml:"objectTypeNames"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-integration.html#cfn-customerprofiles-integration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The URI of the S3 bucket or any other type of data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-integration.html#cfn-customerprofiles-integration-uri
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

