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
type CfnResourceSetProps struct {
	// The descriptive name of the resource set.
	//
	// You can't change the name of a resource set after you create it.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Determines the resources that can be associated to the resource set.
	//
	// Depending on your setting for max results and the number of resource sets, a single call might not return the full list.
	ResourceTypeList *[]*string `field:"required" json:"resourceTypeList" yaml:"resourceTypeList"`
	// A description of the resource set.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The resources included in the resource set.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// A collection of key:value pairs associated with a resource set.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

