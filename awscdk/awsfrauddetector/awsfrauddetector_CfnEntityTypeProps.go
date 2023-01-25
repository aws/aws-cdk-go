package awsfrauddetector

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnEntityType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEntityTypeProps := &cfnEntityTypeProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEntityTypeProps struct {
	// The entity type name.
	//
	// Pattern: `^[0-9a-z_-]+$`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The entity type description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A key and value pair.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

