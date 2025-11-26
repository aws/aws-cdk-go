package previewawsfmsmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnResourceSetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourceSetMixinProps := &CfnResourceSetMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Resources: []*string{
//   		jsii.String("resources"),
//   	},
//   	ResourceTypeList: []*string{
//   		jsii.String("resourceTypeList"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-resourceset.html
//
type CfnResourceSetMixinProps struct {
	// A description of the resource set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-resourceset.html#cfn-fms-resourceset-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The descriptive name of the resource set.
	//
	// You can't change the name of a resource set after you create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-resourceset.html#cfn-fms-resourceset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-resourceset.html#cfn-fms-resourceset-resources
	//
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// Determines the resources that can be associated to the resource set.
	//
	// Depending on your setting for max results and the number of resource sets, a single call might not return the full list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-resourceset.html#cfn-fms-resourceset-resourcetypelist
	//
	ResourceTypeList *[]*string `field:"optional" json:"resourceTypeList" yaml:"resourceTypeList"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-resourceset.html#cfn-fms-resourceset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

