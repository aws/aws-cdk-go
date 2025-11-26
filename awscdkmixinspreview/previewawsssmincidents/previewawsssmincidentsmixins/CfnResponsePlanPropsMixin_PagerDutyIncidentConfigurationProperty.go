package previewawsssmincidentsmixins


// Details about the PagerDuty service where the response plan creates an incident.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pagerDutyIncidentConfigurationProperty := &PagerDutyIncidentConfigurationProperty{
//   	ServiceId: jsii.String("serviceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-pagerdutyincidentconfiguration.html
//
type CfnResponsePlanPropsMixin_PagerDutyIncidentConfigurationProperty struct {
	// The ID of the PagerDuty service that the response plan associates with an incident when it launches.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-pagerdutyincidentconfiguration.html#cfn-ssmincidents-responseplan-pagerdutyincidentconfiguration-serviceid
	//
	ServiceId *string `field:"optional" json:"serviceId" yaml:"serviceId"`
}

