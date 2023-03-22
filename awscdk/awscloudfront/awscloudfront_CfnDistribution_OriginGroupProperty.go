package awscloudfront


// An origin group includes two origins (a primary origin and a second origin to failover to) and a failover criteria that you specify.
//
// You create an origin group to support origin failover in CloudFront. When you create or update a distribution, you can specifiy the origin group instead of a single origin, and CloudFront will failover from the primary origin to the second origin under the failover conditions that you've chosen.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originGroupProperty := &originGroupProperty{
//   	failoverCriteria: &originGroupFailoverCriteriaProperty{
//   		statusCodes: &statusCodesProperty{
//   			items: []interface{}{
//   				jsii.Number(123),
//   			},
//   			quantity: jsii.Number(123),
//   		},
//   	},
//   	id: jsii.String("id"),
//   	members: &originGroupMembersProperty{
//   		items: []interface{}{
//   			&originGroupMemberProperty{
//   				originId: jsii.String("originId"),
//   			},
//   		},
//   		quantity: jsii.Number(123),
//   	},
//   }
//
type CfnDistribution_OriginGroupProperty struct {
	// A complex type that contains information about the failover criteria for an origin group.
	FailoverCriteria interface{} `field:"required" json:"failoverCriteria" yaml:"failoverCriteria"`
	// The origin group's ID.
	Id *string `field:"required" json:"id" yaml:"id"`
	// A complex type that contains information about the origins in an origin group.
	Members interface{} `field:"required" json:"members" yaml:"members"`
}

