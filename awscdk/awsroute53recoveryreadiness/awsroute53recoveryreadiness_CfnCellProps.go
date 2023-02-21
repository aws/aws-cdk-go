package awsroute53recoveryreadiness

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCell`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCellProps := &cfnCellProps{
//   	cellName: jsii.String("cellName"),
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

