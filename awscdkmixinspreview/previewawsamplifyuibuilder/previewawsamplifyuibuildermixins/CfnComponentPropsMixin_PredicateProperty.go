package previewawsamplifyuibuildermixins


// The `Predicate` property specifies information for generating Amplify DataStore queries.
//
// Use `Predicate` to retrieve a subset of the data in a collection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var predicateProperty_ PredicateProperty
//
//   predicateProperty := &PredicateProperty{
//   	And: []interface{}{
//   		predicateProperty_,
//   	},
//   	Field: jsii.String("field"),
//   	Operand: jsii.String("operand"),
//   	OperandType: jsii.String("operandType"),
//   	Operator: jsii.String("operator"),
//   	Or: []interface{}{
//   		predicateProperty_,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-predicate.html
//
type CfnComponentPropsMixin_PredicateProperty struct {
	// A list of predicates to combine logically.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-predicate.html#cfn-amplifyuibuilder-component-predicate-and
	//
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// The field to query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-predicate.html#cfn-amplifyuibuilder-component-predicate-field
	//
	Field *string `field:"optional" json:"field" yaml:"field"`
	// The value to use when performing the evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-predicate.html#cfn-amplifyuibuilder-component-predicate-operand
	//
	Operand *string `field:"optional" json:"operand" yaml:"operand"`
	// The type of value to use when performing the evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-predicate.html#cfn-amplifyuibuilder-component-predicate-operandtype
	//
	OperandType *string `field:"optional" json:"operandType" yaml:"operandType"`
	// The operator to use to perform the evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-predicate.html#cfn-amplifyuibuilder-component-predicate-operator
	//
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// A list of predicates to combine logically.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-predicate.html#cfn-amplifyuibuilder-component-predicate-or
	//
	Or interface{} `field:"optional" json:"or" yaml:"or"`
}

