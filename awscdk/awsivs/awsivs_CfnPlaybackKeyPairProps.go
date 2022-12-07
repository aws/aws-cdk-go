package awsivs

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPlaybackKeyPair`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPlaybackKeyPairProps := &cfnPlaybackKeyPairProps{
//   	publicKeyMaterial: jsii.String("publicKeyMaterial"),
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
type CfnPlaybackKeyPairProps struct {
	// The public portion of a customer-generated key pair.
	PublicKeyMaterial *string `field:"required" json:"publicKeyMaterial" yaml:"publicKeyMaterial"`
	// Playback-key-pair name.
	//
	// The value does not need to be unique.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

