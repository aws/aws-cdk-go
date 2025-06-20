package awscloudfront


// An origin group includes two origins (a primary origin and a secondary origin to failover to) and a failover criteria that you specify.
//
// You create an origin group to support origin failover in CloudFront. When you create or update a distribution, you can specify the origin group instead of a single origin, and CloudFront will failover from the primary origin to the secondary origin under the failover conditions that you've chosen.
//
// Optionally, you can choose selection criteria for your origin group to specify how your origins are selected when your distribution routes viewer requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originGroupProperty := &OriginGroupProperty{
//   	FailoverCriteria: &OriginGroupFailoverCriteriaProperty{
//   		StatusCodes: &StatusCodesProperty{
//   			Items: []interface{}{
//   				jsii.Number(123),
//   			},
//   			Quantity: jsii.Number(123),
//   		},
//   	},
//   	Id: jsii.String("id"),
//   	Members: &OriginGroupMembersProperty{
//   		Items: []interface{}{
//   			&OriginGroupMemberProperty{
//   				OriginId: jsii.String("originId"),
//   			},
//   		},
//   		Quantity: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	SelectionCriteria: jsii.String("selectionCriteria"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroup.html
//
type CfnDistribution_OriginGroupProperty struct {
	// A complex type that contains information about the failover criteria for an origin group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroup.html#cfn-cloudfront-distribution-origingroup-failovercriteria
	//
	FailoverCriteria interface{} `field:"required" json:"failoverCriteria" yaml:"failoverCriteria"`
	// The origin group's ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroup.html#cfn-cloudfront-distribution-origingroup-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// A complex type that contains information about the origins in an origin group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroup.html#cfn-cloudfront-distribution-origingroup-members
	//
	Members interface{} `field:"required" json:"members" yaml:"members"`
	// The selection criteria for the origin group.
	//
	// For more information, see [Create an origin group](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/high_availability_origin_failover.html#concept_origin_groups.creating) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroup.html#cfn-cloudfront-distribution-origingroup-selectioncriteria
	//
	SelectionCriteria *string `field:"optional" json:"selectionCriteria" yaml:"selectionCriteria"`
}

