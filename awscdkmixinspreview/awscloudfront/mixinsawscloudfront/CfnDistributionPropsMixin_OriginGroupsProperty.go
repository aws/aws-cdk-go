package mixinsawscloudfront


// A complex data type for the origin groups specified for a distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   originGroupsProperty := &OriginGroupsProperty{
//   	Items: []interface{}{
//   		&OriginGroupProperty{
//   			FailoverCriteria: &OriginGroupFailoverCriteriaProperty{
//   				StatusCodes: &StatusCodesProperty{
//   					Items: []interface{}{
//   						jsii.Number(123),
//   					},
//   					Quantity: jsii.Number(123),
//   				},
//   			},
//   			Id: jsii.String("id"),
//   			Members: &OriginGroupMembersProperty{
//   				Items: []interface{}{
//   					&OriginGroupMemberProperty{
//   						OriginId: jsii.String("originId"),
//   					},
//   				},
//   				Quantity: jsii.Number(123),
//   			},
//   			SelectionCriteria: jsii.String("selectionCriteria"),
//   		},
//   	},
//   	Quantity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroups.html
//
type CfnDistributionPropsMixin_OriginGroupsProperty struct {
	// The items (origin groups) in a distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroups.html#cfn-cloudfront-distribution-origingroups-items
	//
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// The number of origin groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroups.html#cfn-cloudfront-distribution-origingroups-quantity
	//
	Quantity *float64 `field:"optional" json:"quantity" yaml:"quantity"`
}

