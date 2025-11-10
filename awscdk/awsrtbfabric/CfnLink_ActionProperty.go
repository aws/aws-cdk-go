package awsrtbfabric


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnLink_ActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-action.html#cfn-rtbfabric-link-action-headertag
	//
	HeaderTag interface{} `field:"required" json:"headerTag" yaml:"headerTag"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-action.html#cfn-rtbfabric-link-action-nobid
	//
	NoBid interface{} `field:"required" json:"noBid" yaml:"noBid"`
}

