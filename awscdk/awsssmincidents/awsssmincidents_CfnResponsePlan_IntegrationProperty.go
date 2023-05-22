package awsssmincidents


// Information about third-party services integrated into a response plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationProperty := &IntegrationProperty{
//   	PagerDutyConfiguration: &PagerDutyConfigurationProperty{
//   		Name: jsii.String("name"),
//   		PagerDutyIncidentConfiguration: &PagerDutyIncidentConfigurationProperty{
//   			ServiceId: jsii.String("serviceId"),
//   		},
//   		SecretId: jsii.String("secretId"),
//   	},
//   }
//
type CfnResponsePlan_IntegrationProperty struct {
	// Information about the PagerDuty service where the response plan creates an incident.
	PagerDutyConfiguration interface{} `field:"required" json:"pagerDutyConfiguration" yaml:"pagerDutyConfiguration"`
}

