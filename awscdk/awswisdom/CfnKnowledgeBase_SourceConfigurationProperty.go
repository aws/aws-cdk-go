package awswisdom


// Configuration information about the external data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceConfigurationProperty := &SourceConfigurationProperty{
//   	AppIntegrations: &AppIntegrationsConfigurationProperty{
//   		AppIntegrationArn: jsii.String("appIntegrationArn"),
//
//   		// the properties below are optional
//   		ObjectFields: []*string{
//   			jsii.String("objectFields"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-sourceconfiguration.html
//
type CfnKnowledgeBase_SourceConfigurationProperty struct {
	// Configuration information for Amazon AppIntegrations to automatically ingest content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-sourceconfiguration.html#cfn-wisdom-knowledgebase-sourceconfiguration-appintegrations
	//
	AppIntegrations interface{} `field:"required" json:"appIntegrations" yaml:"appIntegrations"`
}

