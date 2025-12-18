package previewawssecurityhubmixins


// Information about the configuration and status of a Jira Cloud integration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   jiraCloudProperty := &JiraCloudProperty{
//   	AuthStatus: jsii.String("authStatus"),
//   	AuthUrl: jsii.String("authUrl"),
//   	CloudId: jsii.String("cloudId"),
//   	Domain: jsii.String("domain"),
//   	ProjectKey: jsii.String("projectKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-jiracloud.html
//
type CfnConnectorV2PropsMixin_JiraCloudProperty struct {
	// The status of the authorization between Jira Cloud and the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-jiracloud.html#cfn-securityhub-connectorv2-jiracloud-authstatus
	//
	AuthStatus *string `field:"optional" json:"authStatus" yaml:"authStatus"`
	// The URL to provide to customers for OAuth auth code flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-jiracloud.html#cfn-securityhub-connectorv2-jiracloud-authurl
	//
	AuthUrl *string `field:"optional" json:"authUrl" yaml:"authUrl"`
	// The cloud id of the Jira Cloud.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-jiracloud.html#cfn-securityhub-connectorv2-jiracloud-cloudid
	//
	CloudId *string `field:"optional" json:"cloudId" yaml:"cloudId"`
	// The URL domain of your Jira Cloud instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-jiracloud.html#cfn-securityhub-connectorv2-jiracloud-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The projectKey of Jira Cloud.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-jiracloud.html#cfn-securityhub-connectorv2-jiracloud-projectkey
	//
	ProjectKey *string `field:"optional" json:"projectKey" yaml:"projectKey"`
}

