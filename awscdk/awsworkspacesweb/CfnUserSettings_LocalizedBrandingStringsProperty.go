package awsworkspacesweb


// Localized text strings for a specific language that customize the web portal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localizedBrandingStringsProperty := &LocalizedBrandingStringsProperty{
//   	BrowserTabTitle: jsii.String("browserTabTitle"),
//   	WelcomeText: jsii.String("welcomeText"),
//
//   	// the properties below are optional
//   	ContactButtonText: jsii.String("contactButtonText"),
//   	ContactLink: jsii.String("contactLink"),
//   	LoadingText: jsii.String("loadingText"),
//   	LoginButtonText: jsii.String("loginButtonText"),
//   	LoginDescription: jsii.String("loginDescription"),
//   	LoginTitle: jsii.String("loginTitle"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html
//
type CfnUserSettings_LocalizedBrandingStringsProperty struct {
	// The text displayed in the browser tab title.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html#cfn-workspacesweb-usersettings-localizedbrandingstrings-browsertabtitle
	//
	BrowserTabTitle *string `field:"required" json:"browserTabTitle" yaml:"browserTabTitle"`
	// The welcome text displayed on the sign-in page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html#cfn-workspacesweb-usersettings-localizedbrandingstrings-welcometext
	//
	WelcomeText *string `field:"required" json:"welcomeText" yaml:"welcomeText"`
	// The text displayed on the contact button.
	//
	// This field is optional and defaults to "Contact us".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html#cfn-workspacesweb-usersettings-localizedbrandingstrings-contactbuttontext
	//
	ContactButtonText *string `field:"optional" json:"contactButtonText" yaml:"contactButtonText"`
	// A contact link URL.
	//
	// The URL must start with `https://` or `mailto:` . If not provided, the contact button will be hidden from the web portal screen.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html#cfn-workspacesweb-usersettings-localizedbrandingstrings-contactlink
	//
	ContactLink *string `field:"optional" json:"contactLink" yaml:"contactLink"`
	// The text displayed during session loading.
	//
	// This field is optional and defaults to "Loading your session".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html#cfn-workspacesweb-usersettings-localizedbrandingstrings-loadingtext
	//
	LoadingText *string `field:"optional" json:"loadingText" yaml:"loadingText"`
	// The text displayed on the login button.
	//
	// This field is optional and defaults to "Sign In".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html#cfn-workspacesweb-usersettings-localizedbrandingstrings-loginbuttontext
	//
	LoginButtonText *string `field:"optional" json:"loginButtonText" yaml:"loginButtonText"`
	// The description text for the login section.
	//
	// This field is optional and defaults to "Sign in to your session".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html#cfn-workspacesweb-usersettings-localizedbrandingstrings-logindescription
	//
	LoginDescription *string `field:"optional" json:"loginDescription" yaml:"loginDescription"`
	// The title text for the login section.
	//
	// This field is optional and defaults to "Sign In".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html#cfn-workspacesweb-usersettings-localizedbrandingstrings-logintitle
	//
	LoginTitle *string `field:"optional" json:"loginTitle" yaml:"loginTitle"`
}

