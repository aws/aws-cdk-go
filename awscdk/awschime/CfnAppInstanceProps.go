package awschime

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAppInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppInstanceProps := &CfnAppInstanceProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Metadata: jsii.String("metadata"),
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
type CfnAppInstanceProps struct {
	// The name of the AppInstance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstance.html#cfn-chime-appinstance-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The metadata of the AppInstance.
	//
	// Limited to a 1KB string in UTF-8.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstance.html#cfn-chime-appinstance-metadata
	//
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// Tags assigned to the AppInstance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-appinstance.html#cfn-chime-appinstance-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

