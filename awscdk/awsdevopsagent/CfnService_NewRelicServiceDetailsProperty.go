package awsdevopsagent


// New Relic service configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   newRelicServiceDetailsProperty := &NewRelicServiceDetailsProperty{
//   	AuthorizationConfig: &NewRelicAuthorizationConfigProperty{
//   		ApiKey: &NewRelicApiKeyConfigProperty{
//   			AccountId: jsii.String("accountId"),
//   			ApiKey: jsii.String("apiKey"),
//   			Region: jsii.String("region"),
//
//   			// the properties below are optional
//   			AlertPolicyIds: []*string{
//   				jsii.String("alertPolicyIds"),
//   			},
//   			ApplicationIds: []*string{
//   				jsii.String("applicationIds"),
//   			},
//   			EntityGuids: []*string{
//   				jsii.String("entityGuids"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicservicedetails.html
//
type CfnService_NewRelicServiceDetailsProperty struct {
	// New Relic authorization configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-newrelicservicedetails.html#cfn-devopsagent-service-newrelicservicedetails-authorizationconfig
	//
	AuthorizationConfig interface{} `field:"required" json:"authorizationConfig" yaml:"authorizationConfig"`
}

