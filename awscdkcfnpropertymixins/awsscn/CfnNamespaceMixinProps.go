package awsscn

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnNamespacePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnNamespaceMixinProps := &CfnNamespaceMixinProps{
//   	Description: jsii.String("description"),
//   	InstanceId: jsii.String("instanceId"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-namespace.html
//
type CfnNamespaceMixinProps struct {
	// The description of the namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-namespace.html#cfn-scn-namespace-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Web Services Supply Chain instance identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-namespace.html#cfn-scn-namespace-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The name of the namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-namespace.html#cfn-scn-namespace-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags for the namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-namespace.html#cfn-scn-namespace-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

