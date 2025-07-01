package awsservicediscovery

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnHttpNamespace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHttpNamespaceProps := &CfnHttpNamespaceProps{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-httpnamespace.html
//
type CfnHttpNamespaceProps struct {
	// The name that you want to assign to this namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-httpnamespace.html#cfn-servicediscovery-httpnamespace-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description for the namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-httpnamespace.html#cfn-servicediscovery-httpnamespace-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags for the namespace.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-httpnamespace.html#cfn-servicediscovery-httpnamespace-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

