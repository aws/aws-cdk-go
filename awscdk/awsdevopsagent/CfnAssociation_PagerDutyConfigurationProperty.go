package awsdevopsagent


// PagerDuty integration configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pagerDutyConfigurationProperty := &PagerDutyConfigurationProperty{
//   	CustomerEmail: jsii.String("customerEmail"),
//   	Services: []*string{
//   		jsii.String("services"),
//   	},
//
//   	// the properties below are optional
//   	EnableWebhookUpdates: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-pagerdutyconfiguration.html
//
type CfnAssociation_PagerDutyConfigurationProperty struct {
	// Email to be used in PagerDuty API header.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-pagerdutyconfiguration.html#cfn-devopsagent-association-pagerdutyconfiguration-customeremail
	//
	CustomerEmail *string `field:"required" json:"customerEmail" yaml:"customerEmail"`
	// List of PagerDuty service IDs available for the association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-pagerdutyconfiguration.html#cfn-devopsagent-association-pagerdutyconfiguration-services
	//
	Services *[]*string `field:"required" json:"services" yaml:"services"`
	// When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-pagerdutyconfiguration.html#cfn-devopsagent-association-pagerdutyconfiguration-enablewebhookupdates
	//
	EnableWebhookUpdates interface{} `field:"optional" json:"enableWebhookUpdates" yaml:"enableWebhookUpdates"`
}

