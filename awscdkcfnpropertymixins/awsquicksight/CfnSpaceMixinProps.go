package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSpacePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSpaceMixinProps := &CfnSpaceMixinProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Permissions: []interface{}{
//   		&ResourcePermissionProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			Principal: jsii.String("principal"),
//   		},
//   	},
//   	Resources: []interface{}{
//   		&SpaceResourceProperty{
//   			ResourceArn: jsii.String("resourceArn"),
//   			ResourceType: jsii.String("resourceType"),
//   		},
//   	},
//   	SpaceId: jsii.String("spaceId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-space.html
//
type CfnSpaceMixinProps struct {
	// The ID of the Amazon Web Services account where the space is being created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-space.html#cfn-quicksight-space-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// A description of the space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-space.html#cfn-quicksight-space-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-space.html#cfn-quicksight-space-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of permissions granted on the space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-space.html#cfn-quicksight-space-permissions
	//
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// A list of QuickSight resources attached to the space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-space.html#cfn-quicksight-space-resources
	//
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// The unique identifier for the space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-space.html#cfn-quicksight-space-spaceid
	//
	SpaceId *string `field:"optional" json:"spaceId" yaml:"spaceId"`
	// A list of key-value pairs to associate with the space resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-space.html#cfn-quicksight-space-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

