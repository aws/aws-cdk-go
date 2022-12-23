package awsssmincidents


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pagerDutyConfigurationProperty := &pagerDutyConfigurationProperty{
//   	name: jsii.String("name"),
//   	pagerDutyIncidentConfiguration: &pagerDutyIncidentConfigurationProperty{
//   		serviceId: jsii.String("serviceId"),
//   	},
//   	secretId: jsii.String("secretId"),
//   }
//
type CfnResponsePlan_PagerDutyConfigurationProperty struct {
	// `CfnResponsePlan.PagerDutyConfigurationProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnResponsePlan.PagerDutyConfigurationProperty.PagerDutyIncidentConfiguration`.
	PagerDutyIncidentConfiguration interface{} `field:"required" json:"pagerDutyIncidentConfiguration" yaml:"pagerDutyIncidentConfiguration"`
	// `CfnResponsePlan.PagerDutyConfigurationProperty.SecretId`.
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
}

