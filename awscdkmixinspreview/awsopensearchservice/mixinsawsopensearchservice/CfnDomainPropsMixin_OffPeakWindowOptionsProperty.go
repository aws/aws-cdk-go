package mixinsawsopensearchservice


// Off-peak window settings for the domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   offPeakWindowOptionsProperty := &OffPeakWindowOptionsProperty{
//   	Enabled: jsii.Boolean(false),
//   	OffPeakWindow: &OffPeakWindowProperty{
//   		WindowStartTime: &WindowStartTimeProperty{
//   			Hours: jsii.Number(123),
//   			Minutes: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-offpeakwindowoptions.html
//
type CfnDomainPropsMixin_OffPeakWindowOptionsProperty struct {
	// Specifies whether off-peak window settings are enabled for the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-offpeakwindowoptions.html#cfn-opensearchservice-domain-offpeakwindowoptions-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Off-peak window settings for the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-offpeakwindowoptions.html#cfn-opensearchservice-domain-offpeakwindowoptions-offpeakwindow
	//
	OffPeakWindow interface{} `field:"optional" json:"offPeakWindow" yaml:"offPeakWindow"`
}

