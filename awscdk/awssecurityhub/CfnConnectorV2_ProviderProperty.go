package awssecurityhub


// The third-party provider detail for a service configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   providerProperty := &ProviderProperty{
//   	JiraCloud: &JiraCloudProviderConfigurationProperty{
//   		ProjectKey: jsii.String("projectKey"),
//   	},
//   	ServiceNow: &ServiceNowProviderConfigurationProperty{
//   		InstanceName: jsii.String("instanceName"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-provider.html
//
type CfnConnectorV2_ProviderProperty struct {
	// Details about a Jira Cloud integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-provider.html#cfn-securityhub-connectorv2-provider-jiracloud
	//
	JiraCloud interface{} `field:"optional" json:"jiraCloud" yaml:"jiraCloud"`
	// Details about a ServiceNow ITSM integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-provider.html#cfn-securityhub-connectorv2-provider-servicenow
	//
	ServiceNow interface{} `field:"optional" json:"serviceNow" yaml:"serviceNow"`
}

