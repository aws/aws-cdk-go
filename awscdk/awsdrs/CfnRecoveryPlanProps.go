package awsdrs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRecoveryPlan`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecoveryPlanProps := &CfnRecoveryPlanProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-recoveryplan.html
//
type CfnRecoveryPlanProps struct {
	// The name of the Recovery Plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-recoveryplan.html#cfn-drs-recoveryplan-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the Recovery Plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-recoveryplan.html#cfn-drs-recoveryplan-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags associated with the Recovery Plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-recoveryplan.html#cfn-drs-recoveryplan-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

