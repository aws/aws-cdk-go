package awsresourcegroups


// One parameter for a group configuration item.
//
// For details about service configurations and how to construct them, see [Service configurations for resource groups](https://docs.aws.amazon.com//ARG/latest/APIReference/about-slg.html) in the *AWS Resource Groups User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationParameterProperty := &ConfigurationParameterProperty{
//   	Name: jsii.String("name"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourcegroups-group-configurationparameter.html
//
type CfnGroup_ConfigurationParameterProperty struct {
	// The name of the group configuration parameter.
	//
	// For the list of parameters that you can use with each configuration item type, see [Supported resource types and parameters](https://docs.aws.amazon.com//ARG/latest/APIReference/about-slg.html#about-slg-types) in the *AWS Resource Groups User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourcegroups-group-configurationparameter.html#cfn-resourcegroups-group-configurationparameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value or values to be used for the specified parameter.
	//
	// For the list of values you can use with each parameter, see [Supported resource types and parameters](https://docs.aws.amazon.com//ARG/latest/APIReference/about-slg.html#about-slg-types) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourcegroups-group-configurationparameter.html#cfn-resourcegroups-group-configurationparameter-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

