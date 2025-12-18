package awsresourcegroups

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGroupProps := &CfnGroupProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Configuration: []interface{}{
//   		&ConfigurationItemProperty{
//   			Parameters: []interface{}{
//   				&ConfigurationParameterProperty{
//   					Name: jsii.String("name"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	ResourceQuery: &ResourceQueryProperty{
//   		Query: &QueryProperty{
//   			ResourceTypeFilters: []*string{
//   				jsii.String("resourceTypeFilters"),
//   			},
//   			StackIdentifier: jsii.String("stackIdentifier"),
//   			TagFilters: []interface{}{
//   				&TagFilterProperty{
//   					Key: jsii.String("key"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	Resources: []*string{
//   		jsii.String("resources"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourcegroups-group.html
//
type CfnGroupProps struct {
	// The name of a resource group.
	//
	// The name must be unique within the AWS Region in which you create the resource. To create multiple resource groups based on the same CloudFormation stack, you must generate unique names for each.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourcegroups-group.html#cfn-resourcegroups-group-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The service configuration currently associated with the resource group and in effect for the members of the resource group.
	//
	// A `Configuration` consists of one or more `ConfigurationItem` entries. For information about service configurations for resource groups and how to construct them, see [Service configurations for resource groups](https://docs.aws.amazon.com//ARG/latest/APIReference/about-slg.html) in the *Resource Groups User Guide* .
	//
	// > You can include either a `Configuration` or a `ResourceQuery` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourcegroups-group.html#cfn-resourcegroups-group-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The description of the resource group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourcegroups-group.html#cfn-resourcegroups-group-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The resource query structure that is used to dynamically determine which AWS resources are members of the associated resource group.
	//
	// For more information about queries and how to construct them, see [Build queries and groups in Resource Groups](https://docs.aws.amazon.com//ARG/latest/userguide/gettingstarted-query.html) in the *Resource Groups User Guide*
	//
	// > - You can include either a `ResourceQuery` or a `Configuration` , but not both.
	// > - You can specify the group's membership either by using a `ResourceQuery` or by using a list of `Resources` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourcegroups-group.html#cfn-resourcegroups-group-resourcequery
	//
	ResourceQuery interface{} `field:"optional" json:"resourceQuery" yaml:"resourceQuery"`
	// A list of the Amazon Resource Names (ARNs) of AWS resources that you want to add to the specified group.
	//
	// > - You can specify the group membership either by using a list of `Resources` or by using a `ResourceQuery` , but not both.
	// > - You can include a `Resources` property only if you also specify a `Configuration` property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourcegroups-group.html#cfn-resourcegroups-group-resources
	//
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// The tag key and value pairs that are attached to the resource group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourcegroups-group.html#cfn-resourcegroups-group-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

