package awsroute53recoveryreadiness

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnRecoveryGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecoveryGroupProps := &cfnRecoveryGroupProps{
//   	cells: []*string{
//   		jsii.String("cells"),
//   	},
//   	recoveryGroupName: jsii.String("recoveryGroupName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRecoveryGroupProps struct {
	// A list of the cell Amazon Resource Names (ARNs) in the recovery group.
	Cells *[]*string `field:"optional" json:"cells" yaml:"cells"`
	// The name of the recovery group to create.
	RecoveryGroupName *string `field:"optional" json:"recoveryGroupName" yaml:"recoveryGroupName"`
	// A collection of tags associated with a resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

