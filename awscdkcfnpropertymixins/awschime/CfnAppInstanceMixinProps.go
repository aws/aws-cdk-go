package awschime

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAppInstancePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAppInstanceMixinProps := &CfnAppInstanceMixinProps{
//   	Metadata: jsii.String("metadata"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstance.html
//
type CfnAppInstanceMixinProps struct {
	// The metadata of the AppInstance.
	//
	// Limited to a 1KB string in UTF-8.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstance.html#cfn-chime-appinstance-metadata
	//
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// The name of the AppInstance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstance.html#cfn-chime-appinstance-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags assigned to the AppInstance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstance.html#cfn-chime-appinstance-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

