package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Specifies the tags to apply to a resource when the resource is being created for an EC2 Fleet.
//
// `TagSpecification` is a property of the [AWS::EC2::EC2Fleet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html) resource.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-tagspecification.html
//
type CfnEC2Fleet_TagSpecificationProperty struct {
	// The type of resource to tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-tagspecification.html#cfn-ec2-ec2fleet-tagspecification-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The tags to apply to the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-tagspecification.html#cfn-ec2-ec2fleet-tagspecification-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

