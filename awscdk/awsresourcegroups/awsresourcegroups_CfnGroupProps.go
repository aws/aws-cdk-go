package awsresourcegroups

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGroupProps := &cfnGroupProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	configuration: []interface{}{
//   		&configurationItemProperty{
//   			parameters: []interface{}{
//   				&configurationParameterProperty{
//   					name: jsii.String("name"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			type: jsii.String("type"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	resourceQuery: &resourceQueryProperty{
//   		query: &queryProperty{
//   			resourceTypeFilters: []*string{
//   				jsii.String("resourceTypeFilters"),
//   			},
//   			stackIdentifier: jsii.String("stackIdentifier"),
//   			tagFilters: []interface{}{
//   				&tagFilterProperty{
//   					key: jsii.String("key"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   		},
//   		type: jsii.String("type"),
//   	},
//   	resources: []*string{
//   		jsii.String("resources"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnGroupProps struct {
	// The name of a resource group.
	//
	// The name must be unique within the AWS Region in which you create the resource. To create multiple resource groups based on the same CloudFormation stack, you must generate unique names for each.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The service configuration currently associated with the resource group and in effect for the members of the resource group.
	//
	// A `Configuration` consists of one or more `ConfigurationItem` entries. For information about service configurations for resource groups and how to construct them, see [Service configurations for resource groups](https://docs.aws.amazon.com//ARG/latest/APIReference/about-slg.html) in the *AWS Resource Groups User Guide* .
	//
	// > You can include either a `Configuration` or a `ResourceQuery` , but not both.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The description of the resource group.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The resource query structure that is used to dynamically determine which AWS resources are members of the associated resource group.
	//
	// For more information about queries and how to construct them, see [Build queries and groups in AWS Resource Groups](https://docs.aws.amazon.com//ARG/latest/userguide/gettingstarted-query.html) in the *AWS Resource Groups User Guide*
	//
	// > - You can include either a `ResourceQuery` or a `Configuration` , but not both.
	// > - You can specify the group's membership either by using a `ResourceQuery` or by using a list of `Resources` , but not both.
	ResourceQuery interface{} `field:"optional" json:"resourceQuery" yaml:"resourceQuery"`
	// A list of the Amazon Resource Names (ARNs) of AWS resources that you want to add to the specified group.
	//
	// > - You can specify the group membership either by using a list of `Resources` or by using a `ResourceQuery` , but not both.
	// > - You can include a `Resources` property only if you also specify a `Configuration` property.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// The tag key and value pairs that are attached to the resource group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

