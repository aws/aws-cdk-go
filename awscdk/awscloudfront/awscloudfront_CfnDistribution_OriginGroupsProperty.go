package awscloudfront


// A complex data type for the origin groups specified for a distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originGroupsProperty := &originGroupsProperty{
//   	quantity: jsii.Number(123),
//
//   	// the properties below are optional
//   	items: []interface{}{
//   		&originGroupProperty{
//   			failoverCriteria: &originGroupFailoverCriteriaProperty{
//   				statusCodes: &statusCodesProperty{
//   					items: []interface{}{
//   						jsii.Number(123),
//   					},
//   					quantity: jsii.Number(123),
//   				},
//   			},
//   			id: jsii.String("id"),
//   			members: &originGroupMembersProperty{
//   				items: []interface{}{
//   					&originGroupMemberProperty{
//   						originId: jsii.String("originId"),
//   					},
//   				},
//   				quantity: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnDistribution_OriginGroupsProperty struct {
	// The number of origin groups.
	Quantity *float64 `field:"required" json:"quantity" yaml:"quantity"`
	// The items (origin groups) in a distribution.
	Items interface{} `field:"optional" json:"items" yaml:"items"`
}

