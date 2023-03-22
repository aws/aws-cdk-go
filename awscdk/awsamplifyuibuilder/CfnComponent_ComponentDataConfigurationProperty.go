package awsamplifyuibuilder


// The `ComponentDataConfiguration` property specifies the configuration for binding a component's properties to data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var predicateProperty_ predicateProperty
//
//   componentDataConfigurationProperty := &ComponentDataConfigurationProperty{
//   	Model: jsii.String("model"),
//
//   	// the properties below are optional
//   	Identifiers: []*string{
//   		jsii.String("identifiers"),
//   	},
//   	Predicate: &predicateProperty{
//   		And: []interface{}{
//   			predicateProperty_,
//   		},
//   		Field: jsii.String("field"),
//   		Operand: jsii.String("operand"),
//   		Operator: jsii.String("operator"),
//   		Or: []interface{}{
//   			predicateProperty_,
//   		},
//   	},
//   	Sort: []interface{}{
//   		&SortPropertyProperty{
//   			Direction: jsii.String("direction"),
//   			Field: jsii.String("field"),
//   		},
//   	},
//   }
//
type CfnComponent_ComponentDataConfigurationProperty struct {
	// The name of the data model to use to bind data to a component.
	Model *string `field:"required" json:"model" yaml:"model"`
	// A list of IDs to use to bind data to a component.
	//
	// Use this property to bind specifically chosen data, rather than data retrieved from a query.
	Identifiers *[]*string `field:"optional" json:"identifiers" yaml:"identifiers"`
	// Represents the conditional logic to use when binding data to a component.
	//
	// Use this property to retrieve only a subset of the data in a collection.
	Predicate interface{} `field:"optional" json:"predicate" yaml:"predicate"`
	// Describes how to sort the component's properties.
	Sort interface{} `field:"optional" json:"sort" yaml:"sort"`
}

