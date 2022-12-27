package awsssmincidents


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationProperty := &integrationProperty{
//   	pagerDutyConfiguration: &pagerDutyConfigurationProperty{
//   		name: jsii.String("name"),
//   		pagerDutyIncidentConfiguration: &pagerDutyIncidentConfigurationProperty{
//   			serviceId: jsii.String("serviceId"),
//   		},
//   		secretId: jsii.String("secretId"),
//   	},
//   }
//
type CfnResponsePlan_IntegrationProperty struct {
	// `CfnResponsePlan.IntegrationProperty.PagerDutyConfiguration`.
	PagerDutyConfiguration interface{} `field:"required" json:"pagerDutyConfiguration" yaml:"pagerDutyConfiguration"`
}

