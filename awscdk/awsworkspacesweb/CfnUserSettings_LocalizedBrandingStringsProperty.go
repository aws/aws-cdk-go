package awsworkspacesweb


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html#cfn-workspacesweb-usersettings-localizedbrandingstrings-browsertabtitle
	//
	BrowserTabTitle *string `field:"required" json:"browserTabTitle" yaml:"browserTabTitle"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html#cfn-workspacesweb-usersettings-localizedbrandingstrings-welcometext
	//
	WelcomeText *string `field:"required" json:"welcomeText" yaml:"welcomeText"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html#cfn-workspacesweb-usersettings-localizedbrandingstrings-contactbuttontext
	//
	ContactButtonText *string `field:"optional" json:"contactButtonText" yaml:"contactButtonText"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html#cfn-workspacesweb-usersettings-localizedbrandingstrings-contactlink
	//
	ContactLink *string `field:"optional" json:"contactLink" yaml:"contactLink"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html#cfn-workspacesweb-usersettings-localizedbrandingstrings-loadingtext
	//
	LoadingText *string `field:"optional" json:"loadingText" yaml:"loadingText"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html#cfn-workspacesweb-usersettings-localizedbrandingstrings-loginbuttontext
	//
	LoginButtonText *string `field:"optional" json:"loginButtonText" yaml:"loginButtonText"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html#cfn-workspacesweb-usersettings-localizedbrandingstrings-logindescription
	//
	LoginDescription *string `field:"optional" json:"loginDescription" yaml:"loginDescription"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html#cfn-workspacesweb-usersettings-localizedbrandingstrings-logintitle
	//
	LoginTitle *string `field:"optional" json:"loginTitle" yaml:"loginTitle"`
}

