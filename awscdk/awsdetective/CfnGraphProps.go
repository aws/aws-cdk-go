package awsdetective

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnGraph`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGraphProps := &CfnGraphProps{
//   	AutoEnableMembers: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-graph.html
//
type CfnGraphProps struct {
	// Indicates whether to automatically enable new organization accounts as member accounts in the organization behavior graph.
	//
	// By default, this property is set to `false` . If you want to change the value of this property, you must be the Detective administrator for the organization. For more information on setting a Detective administrator account, see [AWS::Detective::OrganizationAdmin](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-organizationadmin.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-graph.html#cfn-detective-graph-autoenablemembers
	//
	AutoEnableMembers interface{} `field:"optional" json:"autoEnableMembers" yaml:"autoEnableMembers"`
	// The tag values to assign to the new behavior graph.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-graph.html#cfn-detective-graph-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

