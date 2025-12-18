package awsresourcegroups


// One of the items in the service configuration assigned to a resource group.
//
// A service configuration can consist of one or more items. For details service configurations and how to construct them, see [Service configurations for resource groups](https://docs.aws.amazon.com//ARG/latest/APIReference/about-slg.html) in the *Resource Groups User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationItemProperty := &ConfigurationItemProperty{
//   	Parameters: []interface{}{
//   		&ConfigurationParameterProperty{
//   			Name: jsii.String("name"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourcegroups-group-configurationitem.html
//
type CfnGroup_ConfigurationItemProperty struct {
	// A collection of parameters for this configuration item.
	//
	// For the list of parameters that you can use with each configuration item `Type` , see [Supported resource types and parameters](https://docs.aws.amazon.com//ARG/latest/APIReference/about-slg.html#about-slg-types) in the *Resource Groups User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourcegroups-group-configurationitem.html#cfn-resourcegroups-group-configurationitem-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Specifies the type of configuration item.
	//
	// Each item must have a unique value for type. For the list of the types that you can specify for a configuration item, see [Supported resource types and parameters](https://docs.aws.amazon.com//ARG/latest/APIReference/about-slg.html#about-slg-types) in the *Resource Groups User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourcegroups-group-configurationitem.html#cfn-resourcegroups-group-configurationitem-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

