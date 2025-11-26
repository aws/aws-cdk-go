package previewawswafmixins


// Properties for CfnSizeConstraintSetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSizeConstraintSetMixinProps := &CfnSizeConstraintSetMixinProps{
//   	Name: jsii.String("name"),
//   	SizeConstraints: []interface{}{
//   		&SizeConstraintProperty{
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			FieldToMatch: &FieldToMatchProperty{
//   				Data: jsii.String("data"),
//   				Type: jsii.String("type"),
//   			},
//   			Size: jsii.Number(123),
//   			TextTransformation: jsii.String("textTransformation"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html
//
type CfnSizeConstraintSetMixinProps struct {
	// The name, if any, of the `SizeConstraintSet` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html#cfn-waf-sizeconstraintset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The size constraint and the part of the web request to check.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html#cfn-waf-sizeconstraintset-sizeconstraints
	//
	SizeConstraints interface{} `field:"optional" json:"sizeConstraints" yaml:"sizeConstraints"`
}

