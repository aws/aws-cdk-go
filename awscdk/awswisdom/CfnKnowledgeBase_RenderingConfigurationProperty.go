package awswisdom


// Information about how to render the content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   renderingConfigurationProperty := &RenderingConfigurationProperty{
//   	TemplateUri: jsii.String("templateUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-renderingconfiguration.html
//
type CfnKnowledgeBase_RenderingConfigurationProperty struct {
	// A URI template containing exactly one variable in `${variableName}` format.
	//
	// This can only be set for `EXTERNAL` knowledge bases. For Salesforce, ServiceNow, and Zendesk, the variable must be one of the following:
	//
	// - Salesforce: `Id` , `ArticleNumber` , `VersionNumber` , `Title` , `PublishStatus` , or `IsDeleted`
	// - ServiceNow: `number` , `short_description` , `sys_mod_count` , `workflow_state` , or `active`
	// - Zendesk: `id` , `title` , `updated_at` , or `draft`
	//
	// The variable is replaced with the actual value for a piece of content when calling [GetContent](https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_GetContent.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-renderingconfiguration.html#cfn-wisdom-knowledgebase-renderingconfiguration-templateuri
	//
	TemplateUri *string `field:"optional" json:"templateUri" yaml:"templateUri"`
}

