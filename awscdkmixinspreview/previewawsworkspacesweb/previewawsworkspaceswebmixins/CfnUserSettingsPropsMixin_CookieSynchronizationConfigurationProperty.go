package previewawsworkspaceswebmixins


// The configuration that specifies which cookies should be synchronized from the end user's local browser to the remote browser.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cookieSynchronizationConfigurationProperty := &CookieSynchronizationConfigurationProperty{
//   	Allowlist: []interface{}{
//   		&CookieSpecificationProperty{
//   			Domain: jsii.String("domain"),
//   			Name: jsii.String("name"),
//   			Path: jsii.String("path"),
//   		},
//   	},
//   	Blocklist: []interface{}{
//   		&CookieSpecificationProperty{
//   			Domain: jsii.String("domain"),
//   			Name: jsii.String("name"),
//   			Path: jsii.String("path"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-cookiesynchronizationconfiguration.html
//
type CfnUserSettingsPropsMixin_CookieSynchronizationConfigurationProperty struct {
	// The list of cookie specifications that are allowed to be synchronized to the remote browser.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-cookiesynchronizationconfiguration.html#cfn-workspacesweb-usersettings-cookiesynchronizationconfiguration-allowlist
	//
	Allowlist interface{} `field:"optional" json:"allowlist" yaml:"allowlist"`
	// The list of cookie specifications that are blocked from being synchronized to the remote browser.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-cookiesynchronizationconfiguration.html#cfn-workspacesweb-usersettings-cookiesynchronizationconfiguration-blocklist
	//
	Blocklist interface{} `field:"optional" json:"blocklist" yaml:"blocklist"`
}

