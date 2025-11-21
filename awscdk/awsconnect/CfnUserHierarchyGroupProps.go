package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnUserHierarchyGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserHierarchyGroupProps := &CfnUserHierarchyGroupProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ParentGroupArn: jsii.String("parentGroupArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-userhierarchygroup.html
//
type CfnUserHierarchyGroupProps struct {
	// The Amazon Resource Name (ARN) of the user hierarchy group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-userhierarchygroup.html#cfn-connect-userhierarchygroup-instancearn
	//
	InstanceArn interface{} `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the user hierarchy group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-userhierarchygroup.html#cfn-connect-userhierarchygroup-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the parent group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-userhierarchygroup.html#cfn-connect-userhierarchygroup-parentgrouparn
	//
	ParentGroupArn *string `field:"optional" json:"parentGroupArn" yaml:"parentGroupArn"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-userhierarchygroup.html#cfn-connect-userhierarchygroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

