package awsappstream


// The configuration for URL redirection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   urlRedirectionConfigProperty := &UrlRedirectionConfigProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	AllowedUrls: []*string{
//   		jsii.String("allowedUrls"),
//   	},
//   	DeniedUrls: []*string{
//   		jsii.String("deniedUrls"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-urlredirectionconfig.html
//
type CfnStack_UrlRedirectionConfigProperty struct {
	// Specifies whether URL redirection is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-urlredirectionconfig.html#cfn-appstream-stack-urlredirectionconfig-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The URLs that are allowed for redirection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-urlredirectionconfig.html#cfn-appstream-stack-urlredirectionconfig-allowedurls
	//
	AllowedUrls *[]*string `field:"optional" json:"allowedUrls" yaml:"allowedUrls"`
	// The URLs that are denied for redirection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-urlredirectionconfig.html#cfn-appstream-stack-urlredirectionconfig-deniedurls
	//
	DeniedUrls *[]*string `field:"optional" json:"deniedUrls" yaml:"deniedUrls"`
}

