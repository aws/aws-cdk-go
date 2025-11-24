package mixinsawsssmincidents


// Details about the PagerDuty configuration for a response plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pagerDutyConfigurationProperty := &PagerDutyConfigurationProperty{
//   	Name: jsii.String("name"),
//   	PagerDutyIncidentConfiguration: &PagerDutyIncidentConfigurationProperty{
//   		ServiceId: jsii.String("serviceId"),
//   	},
//   	SecretId: jsii.String("secretId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-pagerdutyconfiguration.html
//
type CfnResponsePlanPropsMixin_PagerDutyConfigurationProperty struct {
	// The name of the PagerDuty configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-pagerdutyconfiguration.html#cfn-ssmincidents-responseplan-pagerdutyconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Details about the PagerDuty service associated with the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-pagerdutyconfiguration.html#cfn-ssmincidents-responseplan-pagerdutyconfiguration-pagerdutyincidentconfiguration
	//
	PagerDutyIncidentConfiguration interface{} `field:"optional" json:"pagerDutyIncidentConfiguration" yaml:"pagerDutyIncidentConfiguration"`
	// The ID of the AWS Secrets Manager secret that stores your PagerDuty key, either a General Access REST API Key or User Token REST API Key, and other user credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-pagerdutyconfiguration.html#cfn-ssmincidents-responseplan-pagerdutyconfiguration-secretid
	//
	SecretId *string `field:"optional" json:"secretId" yaml:"secretId"`
}

