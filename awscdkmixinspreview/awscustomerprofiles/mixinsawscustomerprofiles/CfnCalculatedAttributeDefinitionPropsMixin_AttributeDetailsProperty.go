package mixinsawscustomerprofiles


// Mathematical expression and a list of attribute items specified in that expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   attributeDetailsProperty := &AttributeDetailsProperty{
//   	Attributes: []interface{}{
//   		&AttributeItemProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Expression: jsii.String("expression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-attributedetails.html
//
type CfnCalculatedAttributeDefinitionPropsMixin_AttributeDetailsProperty struct {
	// Mathematical expression and a list of attribute items specified in that expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-attributedetails.html#cfn-customerprofiles-calculatedattributedefinition-attributedetails-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// Mathematical expression that is performed on attribute items provided in the attribute list.
	//
	// Each element in the expression should follow the structure of \"{ObjectTypeName.AttributeName}\".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-attributedetails.html#cfn-customerprofiles-calculatedattributedefinition-attributedetails-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

