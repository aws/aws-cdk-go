package mixinsawsrtbfabric


// Describes a bid action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   actionProperty := &ActionProperty{
//   	HeaderTag: &HeaderTagActionProperty{
//   		Name: jsii.String("name"),
//   		Value: jsii.String("value"),
//   	},
//   	NoBid: &NoBidActionProperty{
//   		NoBidReasonCode: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-action.html
//
type CfnLinkPropsMixin_ActionProperty struct {
	// Describes the header tag for a bid action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-action.html#cfn-rtbfabric-link-action-headertag
	//
	HeaderTag interface{} `field:"optional" json:"headerTag" yaml:"headerTag"`
	// Describes the parameters of a no bid module.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-action.html#cfn-rtbfabric-link-action-nobid
	//
	NoBid interface{} `field:"optional" json:"noBid" yaml:"noBid"`
}

