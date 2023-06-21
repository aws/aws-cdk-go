package awsservicediscovery

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPrivateDnsNamespace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPrivateDnsNamespaceProps := &CfnPrivateDnsNamespaceProps{
//   	Name: jsii.String("name"),
//   	Vpc: jsii.String("vpc"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Properties: &PropertiesProperty{
//   		DnsProperties: &PrivateDnsPropertiesMutableProperty{
//   			Soa: &SOAProperty{
//   				Ttl: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPrivateDnsNamespaceProps struct {
	// The name that you want to assign to this namespace.
	//
	// When you create a private DNS namespace, AWS Cloud Map automatically creates an Amazon Route 53 private hosted zone that has the same name as the namespace.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the Amazon VPC that you want to associate the namespace with.
	Vpc *string `field:"required" json:"vpc" yaml:"vpc"`
	// A description for the namespace.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Properties for the private DNS namespace.
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// The tags for the namespace.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

