package previewawssecurityhubmixins


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
	// The auth status of the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-jiracloud.html#cfn-securityhub-connectorv2-jiracloud-authstatus
	//
	AuthStatus *string `field:"optional" json:"authStatus" yaml:"authStatus"`
	// The authUrl of the JiraCloud connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-jiracloud.html#cfn-securityhub-connectorv2-jiracloud-authurl
	//
	AuthUrl *string `field:"optional" json:"authUrl" yaml:"authUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-jiracloud.html#cfn-securityhub-connectorv2-jiracloud-cloudid
	//
	CloudId *string `field:"optional" json:"cloudId" yaml:"cloudId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-jiracloud.html#cfn-securityhub-connectorv2-jiracloud-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-jiracloud.html#cfn-securityhub-connectorv2-jiracloud-projectkey
	//
	ProjectKey *string `field:"optional" json:"projectKey" yaml:"projectKey"`
}

