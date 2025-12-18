package awsdevopsagent


// Configuration for ServiceNow integration.
//
// Defines the ServiceNow instance URL, instance ID, and webhook update settings required for the Agent Space to create, update, and manage incidents and change requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNowConfigurationProperty := &ServiceNowConfigurationProperty{
//   	EnableWebhookUpdates: jsii.Boolean(false),
//   	InstanceId: jsii.String("instanceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-servicenowconfiguration.html
//
type CfnAssociation_ServiceNowConfigurationProperty struct {
	// When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-servicenowconfiguration.html#cfn-devopsagent-association-servicenowconfiguration-enablewebhookupdates
	//
	EnableWebhookUpdates interface{} `field:"optional" json:"enableWebhookUpdates" yaml:"enableWebhookUpdates"`
	// ServiceNow instance ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-servicenowconfiguration.html#cfn-devopsagent-association-servicenowconfiguration-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
}

