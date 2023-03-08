package awswaf


// Properties for defining a `CfnSizeConstraintSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSizeConstraintSetProps := &cfnSizeConstraintSetProps{
//   	name: jsii.String("name"),
//   	sizeConstraints: []interface{}{
//   		&sizeConstraintProperty{
//   			comparisonOperator: jsii.String("comparisonOperator"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				data: jsii.String("data"),
//   			},
//   			size: jsii.Number(123),
//   			textTransformation: jsii.String("textTransformation"),
//   		},
//   	},
//   }
//
type CfnSizeConstraintSetProps struct {
	// The name, if any, of the `SizeConstraintSet` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// The size constraint and the part of the web request to check.
	SizeConstraints interface{} `field:"required" json:"sizeConstraints" yaml:"sizeConstraints"`
}

