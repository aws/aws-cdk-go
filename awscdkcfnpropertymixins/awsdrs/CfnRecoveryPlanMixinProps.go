package awsdrs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRecoveryPlanPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnRecoveryPlanMixinProps := &CfnRecoveryPlanMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
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
type CfnRecoveryPlanMixinProps struct {
	// The description of the Recovery Plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-recoveryplan.html#cfn-drs-recoveryplan-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the Recovery Plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-recoveryplan.html#cfn-drs-recoveryplan-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags associated with the Recovery Plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-recoveryplan.html#cfn-drs-recoveryplan-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

