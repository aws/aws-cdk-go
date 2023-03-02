package awswisdom


// Configuration information for Amazon AppIntegrations to automatically ingest content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appIntegrationsConfigurationProperty := &appIntegrationsConfigurationProperty{
//   	appIntegrationArn: jsii.String("appIntegrationArn"),
//   	objectFields: []*string{
//   		jsii.String("objectFields"),
//   	},
//   }
//
type CfnKnowledgeBase_AppIntegrationsConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the AppIntegrations DataIntegration to use for ingesting content.
	AppIntegrationArn *string `field:"required" json:"appIntegrationArn" yaml:"appIntegrationArn"`
	// The fields from the source that are made available to your agents in Wisdom.
	//
	// - For [Salesforce](https://docs.aws.amazon.com/https://developer.salesforce.com/docs/atlas.en-us.knowledge_dev.meta/knowledge_dev/sforce_api_objects_knowledge__kav.htm) , you must include at least `Id` , `ArticleNumber` , `VersionNumber` , `Title` , `PublishStatus` , and `IsDeleted` .
	// - For [ServiceNow](https://docs.aws.amazon.com/https://developer.servicenow.com/dev.do#!/reference/api/rome/rest/knowledge-management-api) , you must include at least `number` , `short_description` , `sys_mod_count` , `workflow_state` , and `active` .
	//
	// Make sure to include additional fields. These fields are indexed and used to source recommendations.
	ObjectFields *[]*string `field:"required" json:"objectFields" yaml:"objectFields"`
}

