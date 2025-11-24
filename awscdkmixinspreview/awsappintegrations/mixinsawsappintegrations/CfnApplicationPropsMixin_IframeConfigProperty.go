package mixinsawsappintegrations


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   iframeConfigProperty := &IframeConfigProperty{
//   	Allow: []*string{
//   		jsii.String("allow"),
//   	},
//   	Sandbox: []*string{
//   		jsii.String("sandbox"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-iframeconfig.html
//
type CfnApplicationPropsMixin_IframeConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-iframeconfig.html#cfn-appintegrations-application-iframeconfig-allow
	//
	Allow *[]*string `field:"optional" json:"allow" yaml:"allow"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-iframeconfig.html#cfn-appintegrations-application-iframeconfig-sandbox
	//
	Sandbox *[]*string `field:"optional" json:"sandbox" yaml:"sandbox"`
}

