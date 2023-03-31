package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDimension`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDimensionProps := &cfnDimensionProps{
//   	stringValues: []*string{
//   		jsii.String("stringValues"),
//   	},
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDimensionProps struct {
	// Specifies the value or list of values for the dimension.
	//
	// For `TOPIC_FILTER` dimensions, this is a pattern used to match the MQTT topic (for example, "admin/#").
	StringValues *[]*string `field:"required" json:"stringValues" yaml:"stringValues"`
	// Specifies the type of dimension.
	//
	// Supported types: `TOPIC_FILTER.`
	Type *string `field:"required" json:"type" yaml:"type"`
	// A unique identifier for the dimension.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Metadata that can be used to manage the dimension.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

