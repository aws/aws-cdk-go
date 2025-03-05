package awsfms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnResourceSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceSetProps := &CfnResourceSetProps{
//   	Name: jsii.String("name"),
//   	ResourceTypeList: []*string{
//   		jsii.String("resourceTypeList"),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Resources: []*string{
//   		jsii.String("resources"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-resourceset.html
//
type CfnResourceSetProps struct {
	// The descriptive name of the resource set.
	//
	// You can't change the name of a resource set after you create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-resourceset.html#cfn-fms-resourceset-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Determines the resources that can be associated to the resource set.
	//
	// Depending on your setting for max results and the number of resource sets, a single call might not return the full list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-resourceset.html#cfn-fms-resourceset-resourcetypelist
	//
	ResourceTypeList *[]*string `field:"required" json:"resourceTypeList" yaml:"resourceTypeList"`
	// A description of the resource set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-resourceset.html#cfn-fms-resourceset-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-resourceset.html#cfn-fms-resourceset-resources
	//
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-resourceset.html#cfn-fms-resourceset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

