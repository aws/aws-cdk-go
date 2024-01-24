package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The tag specifications of an Amazon EBS volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eBSTagSpecificationProperty := &EBSTagSpecificationProperty{
//   	ResourceType: jsii.String("resourceType"),
//
//   	// the properties below are optional
//   	PropagateTags: jsii.String("propagateTags"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-ebstagspecification.html
//
type CfnService_EBSTagSpecificationProperty struct {
	// The type of volume resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-ebstagspecification.html#cfn-ecs-service-ebstagspecification-resourcetype
	//
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Determines whether to propagate the tags from the task definition to the Amazon EBS volume.
	//
	// Tags can only propagate to a `SERVICE` specified in `ServiceVolumeConfiguration` . If no value is specified, the tags aren't propagated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-ebstagspecification.html#cfn-ecs-service-ebstagspecification-propagatetags
	//
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The tags applied to this Amazon EBS volume.
	//
	// `AmazonECSCreated` and `AmazonECSManaged` are reserved tags that can't be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-ebstagspecification.html#cfn-ecs-service-ebstagspecification-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

