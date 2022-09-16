package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnIntegration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIntegrationProps := &cfnIntegrationProps{
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	flowDefinition: &flowDefinitionProperty{
//   		flowName: jsii.String("flowName"),
//   		kmsArn: jsii.String("kmsArn"),
//   		sourceFlowConfig: &sourceFlowConfigProperty{
//   			connectorType: jsii.String("connectorType"),
//   			sourceConnectorProperties: &sourceConnectorPropertiesProperty{
//   				marketo: &marketoSourcePropertiesProperty{
//   					object: jsii.String("object"),
//   				},
//   				s3: &s3SourcePropertiesProperty{
//   					bucketName: jsii.String("bucketName"),
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   				},
//   				salesforce: &salesforceSourcePropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					enableDynamicFieldUpdate: jsii.Boolean(false),
//   					includeDeletedRecords: jsii.Boolean(false),
//   				},
//   				serviceNow: &serviceNowSourcePropertiesProperty{
//   					object: jsii.String("object"),
//   				},
//   				zendesk: &zendeskSourcePropertiesProperty{
//   					object: jsii.String("object"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			connectorProfileName: jsii.String("connectorProfileName"),
//   			incrementalPullConfig: &incrementalPullConfigProperty{
//   				datetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   			},
//   		},
//   		tasks: []interface{}{
//   			&taskProperty{
//   				sourceFields: []*string{
//   					jsii.String("sourceFields"),
//   				},
//   				taskType: jsii.String("taskType"),
//
//   				// the properties below are optional
//   				connectorOperator: &connectorOperatorProperty{
//   					marketo: jsii.String("marketo"),
//   					s3: jsii.String("s3"),
//   					salesforce: jsii.String("salesforce"),
//   					serviceNow: jsii.String("serviceNow"),
//   					zendesk: jsii.String("zendesk"),
//   				},
//   				destinationField: jsii.String("destinationField"),
//   				taskProperties: []interface{}{
//   					&taskPropertiesMapProperty{
//   						operatorPropertyKey: jsii.String("operatorPropertyKey"),
//   						property: jsii.String("property"),
//   					},
//   				},
//   			},
//   		},
//   		triggerConfig: &triggerConfigProperty{
//   			triggerType: jsii.String("triggerType"),
//
//   			// the properties below are optional
//   			triggerProperties: &triggerPropertiesProperty{
//   				scheduled: &scheduledTriggerPropertiesProperty{
//   					scheduleExpression: jsii.String("scheduleExpression"),
//
//   					// the properties below are optional
//   					dataPullMode: jsii.String("dataPullMode"),
//   					firstExecutionFrom: jsii.Number(123),
//   					scheduleEndTime: jsii.Number(123),
//   					scheduleOffset: jsii.Number(123),
//   					scheduleStartTime: jsii.Number(123),
//   					timezone: jsii.String("timezone"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   	},
//   	objectTypeName: jsii.String("objectTypeName"),
//   	objectTypeNames: []interface{}{
//   		&objectTypeMappingProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	uri: jsii.String("uri"),
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

