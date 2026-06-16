package awsappstream


// The content redirection settings for the stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   contentRedirectionProperty := &ContentRedirectionProperty{
//   	HostToClient: &UrlRedirectionConfigProperty{
//   		AllowedUrls: []*string{
//   			jsii.String("allowedUrls"),
//   		},
//   		DeniedUrls: []*string{
//   			jsii.String("deniedUrls"),
//   		},
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-contentredirection.html
//
type CfnStackPropsMixin_ContentRedirectionProperty struct {
	// The configuration for URL redirection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-contentredirection.html#cfn-appstream-stack-contentredirection-hosttoclient
	//
	HostToClient interface{} `field:"optional" json:"hostToClient" yaml:"hostToClient"`
}

