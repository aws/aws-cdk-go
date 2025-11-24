package mixinsawsnetworkmanager


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   segmentsProperty := &SegmentsProperty{
//   	SendTo: []*string{
//   		jsii.String("sendTo"),
//   	},
//   	SendVia: []*string{
//   		jsii.String("sendVia"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-segments.html
//
type CfnCoreNetworkPropsMixin_SegmentsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-segments.html#cfn-networkmanager-corenetwork-segments-sendto
	//
	SendTo *[]*string `field:"optional" json:"sendTo" yaml:"sendTo"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-segments.html#cfn-networkmanager-corenetwork-segments-sendvia
	//
	SendVia *[]*string `field:"optional" json:"sendVia" yaml:"sendVia"`
}

