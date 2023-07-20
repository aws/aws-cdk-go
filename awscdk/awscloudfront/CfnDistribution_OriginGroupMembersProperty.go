package awscloudfront


// A complex data type for the origins included in an origin group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originGroupMembersProperty := &OriginGroupMembersProperty{
//   	Items: []interface{}{
//   		&OriginGroupMemberProperty{
//   			OriginId: jsii.String("originId"),
//   		},
//   	},
//   	Quantity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroupmembers.html
//
type CfnDistribution_OriginGroupMembersProperty struct {
	// Items (origins) in an origin group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroupmembers.html#cfn-cloudfront-distribution-origingroupmembers-items
	//
	Items interface{} `field:"required" json:"items" yaml:"items"`
	// The number of origins in an origin group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroupmembers.html#cfn-cloudfront-distribution-origingroupmembers-quantity
	//
	Quantity *float64 `field:"required" json:"quantity" yaml:"quantity"`
}

