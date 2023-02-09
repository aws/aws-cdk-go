package awsdetective

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnGraph`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGraphProps := &cfnGraphProps{
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnGraphProps struct {
	// The tag values to assign to the new behavior graph.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

