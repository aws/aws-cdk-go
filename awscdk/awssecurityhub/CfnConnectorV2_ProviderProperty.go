package awssecurityhub


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   providerProperty := &ProviderProperty{
//   	JiraCloud: &JiraCloudProperty{
//   		ProjectKey: jsii.String("projectKey"),
//
//   		// the properties below are optional
//   		AuthStatus: jsii.String("authStatus"),
//   		AuthUrl: jsii.String("authUrl"),
//   		CloudId: jsii.String("cloudId"),
//   		Domain: jsii.String("domain"),
//   	},
//   	ServiceNow: &ServiceNowProperty{
//   		InstanceName: jsii.String("instanceName"),
//   		SecretArn: jsii.String("secretArn"),
//
//   		// the properties below are optional
//   		AuthStatus: jsii.String("authStatus"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-provider.html
//
type CfnConnectorV2_ProviderProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-provider.html#cfn-securityhub-connectorv2-provider-jiracloud
	//
	JiraCloud interface{} `field:"optional" json:"jiraCloud" yaml:"jiraCloud"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-provider.html#cfn-securityhub-connectorv2-provider-servicenow
	//
	ServiceNow interface{} `field:"optional" json:"serviceNow" yaml:"serviceNow"`
}

