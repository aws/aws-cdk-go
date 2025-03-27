package awsqbusiness


// The container for browser extension configuration for an Amazon Q Business web experience.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   browserExtensionConfigurationProperty := &BrowserExtensionConfigurationProperty{
//   	EnabledBrowserExtensions: []*string{
//   		jsii.String("enabledBrowserExtensions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-webexperience-browserextensionconfiguration.html
//
type CfnWebExperience_BrowserExtensionConfigurationProperty struct {
	// Specify the browser extensions allowed for your Amazon Q web experience.
	//
	// - `CHROME` — Enables the extension for Chromium-based browsers (Google Chrome, Microsoft Edge, Opera, etc.).
	// - `FIREFOX` — Enables the extension for Mozilla Firefox.
	// - `CHROME` and `FIREFOX` — Enable the extension for Chromium-based browsers and Mozilla Firefox.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-webexperience-browserextensionconfiguration.html#cfn-qbusiness-webexperience-browserextensionconfiguration-enabledbrowserextensions
	//
	EnabledBrowserExtensions *[]*string `field:"required" json:"enabledBrowserExtensions" yaml:"enabledBrowserExtensions"`
}

