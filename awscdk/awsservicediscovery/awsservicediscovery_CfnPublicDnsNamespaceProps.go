package awsservicediscovery

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPublicDnsNamespace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPublicDnsNamespaceProps := &CfnPublicDnsNamespaceProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Properties: &PropertiesProperty{
//   		DnsProperties: &PublicDnsPropertiesMutableProperty{
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
type CfnPublicDnsNamespaceProps struct {
	// The name that you want to assign to this namespace.
	//
	// > Do not include sensitive information in the name. The name is publicly available using DNS queries.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description for the namespace.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Properties for the public DNS namespace.
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// The tags for the namespace.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

