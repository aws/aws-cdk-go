package awsssmincidents


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pagerDutyIncidentConfigurationProperty := &pagerDutyIncidentConfigurationProperty{
//   	serviceId: jsii.String("serviceId"),
//   }
//
type CfnResponsePlan_PagerDutyIncidentConfigurationProperty struct {
	// `CfnResponsePlan.PagerDutyIncidentConfigurationProperty.ServiceId`.
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
}

