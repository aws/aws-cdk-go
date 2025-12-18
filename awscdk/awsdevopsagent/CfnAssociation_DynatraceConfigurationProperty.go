package awsdevopsagent


// Configuration for Dynatrace monitoring integration.
//
// Defines the Dynatrace environment ID, list of resources to monitor, and webhook update settings required for the Agent Space to access metrics, traces, and logs from Dynatrace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynatraceConfigurationProperty := &DynatraceConfigurationProperty{
//   	EnvId: jsii.String("envId"),
//
//   	// the properties below are optional
//   	EnableWebhookUpdates: jsii.Boolean(false),
//   	Resources: []*string{
//   		jsii.String("resources"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-dynatraceconfiguration.html
//
type CfnAssociation_DynatraceConfigurationProperty struct {
	// Dynatrace environment id.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-dynatraceconfiguration.html#cfn-devopsagent-association-dynatraceconfiguration-envid
	//
	EnvId *string `field:"required" json:"envId" yaml:"envId"`
	// When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-dynatraceconfiguration.html#cfn-devopsagent-association-dynatraceconfiguration-enablewebhookupdates
	//
	EnableWebhookUpdates interface{} `field:"optional" json:"enableWebhookUpdates" yaml:"enableWebhookUpdates"`
	// List of Dynatrace resources to monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-dynatraceconfiguration.html#cfn-devopsagent-association-dynatraceconfiguration-resources
	//
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

