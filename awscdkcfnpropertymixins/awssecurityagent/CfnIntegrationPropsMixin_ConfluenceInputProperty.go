package awssecurityagent


// The input required to create a Confluence integration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   confluenceInputProperty := &ConfluenceInputProperty{
//   	Code: jsii.String("code"),
//   	InstallationId: jsii.String("installationId"),
//   	SiteUrl: jsii.String("siteUrl"),
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-confluenceinput.html
//
type CfnIntegrationPropsMixin_ConfluenceInputProperty struct {
	// The OAuth 2.0 authorization code returned from the consent redirect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-confluenceinput.html#cfn-securityagent-integration-confluenceinput-code
	//
	Code *string `field:"optional" json:"code" yaml:"code"`
	// The Atlassian installation identifier, available from the Atlassian administration console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-confluenceinput.html#cfn-securityagent-integration-confluenceinput-installationid
	//
	InstallationId *string `field:"optional" json:"installationId" yaml:"installationId"`
	// The Confluence Cloud site URL, for example https://mysite.atlassian.net.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-confluenceinput.html#cfn-securityagent-integration-confluenceinput-siteurl
	//
	SiteUrl *string `field:"optional" json:"siteUrl" yaml:"siteUrl"`
	// The CSRF state token echoed back from the OAuth redirect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-confluenceinput.html#cfn-securityagent-integration-confluenceinput-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

