package previewawsservicediscoverymixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnPublicDnsNamespacePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPublicDnsNamespaceMixinProps := &CfnPublicDnsNamespaceMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Properties: &PropertiesProperty{
//   		DnsProperties: &PublicDnsPropertiesMutableProperty{
//   			Soa: &SOAProperty{
//   				Ttl: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-publicdnsnamespace.html
//
type CfnPublicDnsNamespaceMixinProps struct {
	// A description for the namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-publicdnsnamespace.html#cfn-servicediscovery-publicdnsnamespace-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name that you want to assign to this namespace.
	//
	// > Do not include sensitive information in the name. The name is publicly available using DNS queries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-publicdnsnamespace.html#cfn-servicediscovery-publicdnsnamespace-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Properties for the public DNS namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-publicdnsnamespace.html#cfn-servicediscovery-publicdnsnamespace-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// The tags for the namespace.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-publicdnsnamespace.html#cfn-servicediscovery-publicdnsnamespace-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

