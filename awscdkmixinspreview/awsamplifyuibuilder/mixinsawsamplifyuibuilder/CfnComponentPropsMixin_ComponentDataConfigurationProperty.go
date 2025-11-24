package mixinsawsamplifyuibuilder


// The `ComponentDataConfiguration` property specifies the configuration for binding a component's properties to data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var predicateProperty_ PredicateProperty
//
//   componentDataConfigurationProperty := &ComponentDataConfigurationProperty{
//   	Identifiers: []*string{
//   		jsii.String("identifiers"),
//   	},
//   	Model: jsii.String("model"),
//   	Predicate: &PredicateProperty{
//   		And: []interface{}{
//   			predicateProperty_,
//   		},
//   		Field: jsii.String("field"),
//   		Operand: jsii.String("operand"),
//   		OperandType: jsii.String("operandType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentdataconfiguration.html
//
type CfnComponentPropsMixin_ComponentDataConfigurationProperty struct {
	// A list of IDs to use to bind data to a component.
	//
	// Use this property to bind specifically chosen data, rather than data retrieved from a query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentdataconfiguration.html#cfn-amplifyuibuilder-component-componentdataconfiguration-identifiers
	//
	Identifiers *[]*string `field:"optional" json:"identifiers" yaml:"identifiers"`
	// The name of the data model to use to bind data to a component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentdataconfiguration.html#cfn-amplifyuibuilder-component-componentdataconfiguration-model
	//
	Model *string `field:"optional" json:"model" yaml:"model"`
	// Represents the conditional logic to use when binding data to a component.
	//
	// Use this property to retrieve only a subset of the data in a collection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentdataconfiguration.html#cfn-amplifyuibuilder-component-componentdataconfiguration-predicate
	//
	Predicate interface{} `field:"optional" json:"predicate" yaml:"predicate"`
	// Describes how to sort the component's properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentdataconfiguration.html#cfn-amplifyuibuilder-component-componentdataconfiguration-sort
	//
	Sort interface{} `field:"optional" json:"sort" yaml:"sort"`
}

