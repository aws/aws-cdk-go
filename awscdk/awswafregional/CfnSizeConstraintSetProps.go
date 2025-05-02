package awswafregional


// Properties for defining a `CfnSizeConstraintSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSizeConstraintSetProps := &CfnSizeConstraintSetProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	SizeConstraints: []interface{}{
//   		&SizeConstraintProperty{
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			FieldToMatch: &FieldToMatchProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Data: jsii.String("data"),
//   			},
//   			Size: jsii.Number(123),
//   			TextTransformation: jsii.String("textTransformation"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html
//
type CfnSizeConstraintSetProps struct {
	// The name, if any, of the `SizeConstraintSet` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html#cfn-wafregional-sizeconstraintset-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The size constraint and the part of the web request to check.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html#cfn-wafregional-sizeconstraintset-sizeconstraints
	//
	SizeConstraints interface{} `field:"optional" json:"sizeConstraints" yaml:"sizeConstraints"`
}

