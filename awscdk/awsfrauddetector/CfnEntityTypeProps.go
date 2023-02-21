package awsfrauddetector

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEntityType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEntityTypeProps := &CfnEntityTypeProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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

