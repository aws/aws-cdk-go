package previewawsssmincidentsmixins


// Information about third-party services integrated into a response plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-integration.html
//
type CfnResponsePlanPropsMixin_IntegrationProperty struct {
	// Information about the PagerDuty service where the response plan creates an incident.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-integration.html#cfn-ssmincidents-responseplan-integration-pagerdutyconfiguration
	//
	PagerDutyConfiguration interface{} `field:"optional" json:"pagerDutyConfiguration" yaml:"pagerDutyConfiguration"`
}

