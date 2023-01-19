package awsservicediscovery

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnHttpNamespace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHttpNamespaceProps := &cfnHttpNamespaceProps{
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
type CfnHttpNamespaceProps struct {
	// The name that you want to assign to this namespace.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description for the namespace.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags for the namespace.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

