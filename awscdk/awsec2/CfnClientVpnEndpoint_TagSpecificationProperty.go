package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The tags to apply to a resource when the resource is being created.
//
// When you specify a tag, you must specify the resource type to tag, otherwise the request will fail.
//
// > The `Valid Values` lists all the resource types that can be tagged. However, the action you're using might not support tagging all of these resource types. If you try to tag a resource type that is unsupported for the action you're using, you'll get an error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagSpecificationProperty := &TagSpecificationProperty{
//   	ResourceType: jsii.String("resourceType"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-tagspecification.html
//
type CfnClientVpnEndpoint_TagSpecificationProperty struct {
	// The type of resource to tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-tagspecification.html#cfn-ec2-clientvpnendpoint-tagspecification-resourcetype
	//
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// The tags to apply to the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-tagspecification.html#cfn-ec2-clientvpnendpoint-tagspecification-tags
	//
	Tags *[]*awscdk.CfnTag `field:"required" json:"tags" yaml:"tags"`
}

