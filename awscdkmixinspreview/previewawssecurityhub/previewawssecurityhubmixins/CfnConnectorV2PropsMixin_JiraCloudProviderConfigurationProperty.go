package previewawssecurityhubmixins


// The initial configuration settings required to establish an integration between Security Hub CSPM and Jira Cloud.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   jiraCloudProviderConfigurationProperty := &JiraCloudProviderConfigurationProperty{
//   	ProjectKey: jsii.String("projectKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-jiracloudproviderconfiguration.html
//
type CfnConnectorV2PropsMixin_JiraCloudProviderConfigurationProperty struct {
	// The project key for a JiraCloud instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-jiracloudproviderconfiguration.html#cfn-securityhub-connectorv2-jiracloudproviderconfiguration-projectkey
	//
	ProjectKey *string `field:"optional" json:"projectKey" yaml:"projectKey"`
}

