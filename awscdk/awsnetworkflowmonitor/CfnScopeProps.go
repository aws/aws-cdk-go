package awsnetworkflowmonitor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnScope`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScopeProps := &CfnScopeProps{
//   	Targets: []interface{}{
//   		&TargetResourceProperty{
//   			Region: jsii.String("region"),
//   			TargetIdentifier: &TargetIdentifierProperty{
//   				TargetId: &TargetIdProperty{
//   					AccountId: jsii.String("accountId"),
//   				},
//   				TargetType: jsii.String("targetType"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkflowmonitor-scope.html
//
type CfnScopeProps struct {
	// The targets for the scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkflowmonitor-scope.html#cfn-networkflowmonitor-scope-targets
	//
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// The tags for the scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkflowmonitor-scope.html#cfn-networkflowmonitor-scope-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

