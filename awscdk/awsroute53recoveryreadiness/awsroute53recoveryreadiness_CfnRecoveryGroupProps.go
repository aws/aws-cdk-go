package awsroute53recoveryreadiness

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRecoveryGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecoveryGroupProps := &cfnRecoveryGroupProps{
//   	recoveryGroupName: jsii.String("recoveryGroupName"),
//
//   	// the properties below are optional
//   	cells: []*string{
//   		jsii.String("cells"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRecoveryGroupProps struct {
	// The name of the recovery group to create.
	RecoveryGroupName *string `field:"required" json:"recoveryGroupName" yaml:"recoveryGroupName"`
	// A list of the cell Amazon Resource Names (ARNs) in the recovery group.
	Cells *[]*string `field:"optional" json:"cells" yaml:"cells"`
	// A collection of tags associated with a resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

