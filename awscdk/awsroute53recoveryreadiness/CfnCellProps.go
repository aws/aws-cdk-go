package awsroute53recoveryreadiness

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCell`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCellProps := &CfnCellProps{
//   	CellName: jsii.String("cellName"),
//   	Cells: []*string{
//   		jsii.String("cells"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCellProps struct {
	// The name of the cell to create.
	CellName *string `field:"optional" json:"cellName" yaml:"cellName"`
	// A list of cell Amazon Resource Names (ARNs) contained within this cell, for use in nested cells.
	//
	// For example, Availability Zones within specific AWS Regions .
	Cells *[]*string `field:"optional" json:"cells" yaml:"cells"`
	// A collection of tags associated with a resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

