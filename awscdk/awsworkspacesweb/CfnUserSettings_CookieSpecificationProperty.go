package awsworkspacesweb


// Specifies a single cookie or set of cookies in an end user's browser.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cookieSpecificationProperty := &CookieSpecificationProperty{
//   	Domain: jsii.String("domain"),
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	Path: jsii.String("path"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-cookiespecification.html
//
type CfnUserSettings_CookieSpecificationProperty struct {
	// The domain of the cookie.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-cookiespecification.html#cfn-workspacesweb-usersettings-cookiespecification-domain
	//
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The name of the cookie.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-cookiespecification.html#cfn-workspacesweb-usersettings-cookiespecification-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The path of the cookie.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-cookiespecification.html#cfn-workspacesweb-usersettings-cookiespecification-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
}

