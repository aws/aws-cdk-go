package mixinsawsworkspacesweb


// The policy that specifies which URLs end users are allowed to access or which URLs or domain categories they are restricted from accessing for enhanced security.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   webContentFilteringPolicyProperty := &WebContentFilteringPolicyProperty{
//   	AllowedUrls: []*string{
//   		jsii.String("allowedUrls"),
//   	},
//   	BlockedCategories: []*string{
//   		jsii.String("blockedCategories"),
//   	},
//   	BlockedUrls: []*string{
//   		jsii.String("blockedUrls"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-browsersettings-webcontentfilteringpolicy.html
//
type CfnBrowserSettingsPropsMixin_WebContentFilteringPolicyProperty struct {
	// URLs and domains that are always accessible to end users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-browsersettings-webcontentfilteringpolicy.html#cfn-workspacesweb-browsersettings-webcontentfilteringpolicy-allowedurls
	//
	AllowedUrls *[]*string `field:"optional" json:"allowedUrls" yaml:"allowedUrls"`
	// Categories of websites that are blocked on the end user's browsers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-browsersettings-webcontentfilteringpolicy.html#cfn-workspacesweb-browsersettings-webcontentfilteringpolicy-blockedcategories
	//
	BlockedCategories *[]*string `field:"optional" json:"blockedCategories" yaml:"blockedCategories"`
	// URLs and domains that end users cannot access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-browsersettings-webcontentfilteringpolicy.html#cfn-workspacesweb-browsersettings-webcontentfilteringpolicy-blockedurls
	//
	BlockedUrls *[]*string `field:"optional" json:"blockedUrls" yaml:"blockedUrls"`
}

