package awsssmincidents


// Details about the PagerDuty service where the response plan creates an incident.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pagerDutyIncidentConfigurationProperty := &PagerDutyIncidentConfigurationProperty{
//   	ServiceId: jsii.String("serviceId"),
//   }
//
type CfnResponsePlan_PagerDutyIncidentConfigurationProperty struct {
	// The ID of the PagerDuty service that the response plan associates with an incident when it launches.
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
}

