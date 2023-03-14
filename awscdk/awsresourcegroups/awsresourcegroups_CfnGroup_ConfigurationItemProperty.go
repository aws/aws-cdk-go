package awsresourcegroups


// One of the items in the service configuration assigned to a resource group.
//
// A service configuration can consist of one or more items. For details service configurations and how to construct them, see [Service configurations for resource groups](https://docs.aws.amazon.com//ARG/latest/APIReference/about-slg.html) in the *AWS Resource Groups User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationItemProperty := &configurationItemProperty{
//   	parameters: []interface{}{
//   		&configurationParameterProperty{
//   			name: jsii.String("name"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnGroup_ConfigurationItemProperty struct {
	// A collection of parameters for this configuration item.
	//
	// For the list of parameters that you can use with each configuration item `Type` , see [Supported resource types and parameters](https://docs.aws.amazon.com/ARG/latest/APIReference/about-slg.html#about-slg-types) in the *AWS Resource Groups User Guide* .
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Specifies the type of configuration item.
	//
	// Each item must have a unique value for type. For the list of the types that you can specify for a configuration item, see [Supported resource types and parameters](https://docs.aws.amazon.com/ARG/latest/APIReference/about-slg.html#about-slg-types) in the *AWS Resource Groups User Guide* .
	Type *string `field:"optional" json:"type" yaml:"type"`
}

