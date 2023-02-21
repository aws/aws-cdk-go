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
//   		ObjectFields: []*string{
//   			jsii.String("objectFields"),
//   		},
//   	},
//   }
//
type CfnKnowledgeBase_SourceConfigurationProperty struct {
	// Configuration information for Amazon AppIntegrations to automatically ingest content.
	AppIntegrations interface{} `field:"required" json:"appIntegrations" yaml:"appIntegrations"`
}

