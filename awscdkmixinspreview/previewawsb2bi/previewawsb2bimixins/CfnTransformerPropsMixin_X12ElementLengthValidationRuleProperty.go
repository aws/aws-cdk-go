package previewawsb2bimixins


// Defines a validation rule that specifies custom length constraints for a specific X12 element.
//
// This rule allows you to override the standard minimum and maximum length requirements for an element, enabling validation of trading partner-specific length requirements that may differ from the X12 specification. Both minimum and maximum length values must be specified.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   x12ElementLengthValidationRuleProperty := &X12ElementLengthValidationRuleProperty{
//   	ElementId: jsii.String("elementId"),
//   	MaxLength: jsii.Number(123),
//   	MinLength: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12elementlengthvalidationrule.html
//
type CfnTransformerPropsMixin_X12ElementLengthValidationRuleProperty struct {
	// Specifies the four-digit element ID to which the length constraints will be applied.
	//
	// This identifies which X12 element will have its length requirements modified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12elementlengthvalidationrule.html#cfn-b2bi-transformer-x12elementlengthvalidationrule-elementid
	//
	ElementId *string `field:"optional" json:"elementId" yaml:"elementId"`
	// Specifies the maximum allowed length for the identified element.
	//
	// This value defines the upper limit for the element's content length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12elementlengthvalidationrule.html#cfn-b2bi-transformer-x12elementlengthvalidationrule-maxlength
	//
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	// Specifies the minimum required length for the identified element.
	//
	// This value defines the lower limit for the element's content length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12elementlengthvalidationrule.html#cfn-b2bi-transformer-x12elementlengthvalidationrule-minlength
	//
	MinLength *float64 `field:"optional" json:"minLength" yaml:"minLength"`
}

