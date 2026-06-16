package awsappstream


// The configuration for URL redirection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   urlRedirectionConfigProperty := &UrlRedirectionConfigProperty{
//   	AllowedUrls: []*string{
//   		jsii.String("allowedUrls"),
//   	},
//   	DeniedUrls: []*string{
//   		jsii.String("deniedUrls"),
//   	},
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-urlredirectionconfig.html
//
type CfnStackPropsMixin_UrlRedirectionConfigProperty struct {
	// The URLs that are allowed for redirection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-urlredirectionconfig.html#cfn-appstream-stack-urlredirectionconfig-allowedurls
	//
	AllowedUrls *[]*string `field:"optional" json:"allowedUrls" yaml:"allowedUrls"`
	// The URLs that are denied for redirection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-urlredirectionconfig.html#cfn-appstream-stack-urlredirectionconfig-deniedurls
	//
	DeniedUrls *[]*string `field:"optional" json:"deniedUrls" yaml:"deniedUrls"`
	// Specifies whether URL redirection is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-urlredirectionconfig.html#cfn-appstream-stack-urlredirectionconfig-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

