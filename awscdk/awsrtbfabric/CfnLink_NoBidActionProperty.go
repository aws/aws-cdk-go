package awsrtbfabric


// Describes a no bid action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   noBidActionProperty := &NoBidActionProperty{
//   	NoBidReasonCode: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-nobidaction.html
//
type CfnLink_NoBidActionProperty struct {
	// The reason code for the no bid action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-nobidaction.html#cfn-rtbfabric-link-nobidaction-nobidreasoncode
	//
	NoBidReasonCode *float64 `field:"optional" json:"noBidReasonCode" yaml:"noBidReasonCode"`
}

