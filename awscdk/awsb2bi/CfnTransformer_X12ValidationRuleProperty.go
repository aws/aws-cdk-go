package awsb2bi


// Represents a single validation rule that can be applied during X12 EDI processing.
//
// This is a union type that can contain one of several specific validation rule types: code list validation rules for modifying allowed element codes, element length validation rules for enforcing custom length constraints, or element requirement validation rules for changing mandatory/optional status. Each validation rule targets specific aspects of EDI document validation to ensure compliance with trading partner requirements and business rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   x12ValidationRuleProperty := &X12ValidationRuleProperty{
//   	CodeListValidationRule: &X12CodeListValidationRuleProperty{
//   		ElementId: jsii.String("elementId"),
//
//   		// the properties below are optional
//   		CodesToAdd: []*string{
//   			jsii.String("codesToAdd"),
//   		},
//   		CodesToRemove: []*string{
//   			jsii.String("codesToRemove"),
//   		},
//   	},
//   	ElementLengthValidationRule: &X12ElementLengthValidationRuleProperty{
//   		ElementId: jsii.String("elementId"),
//   		MaxLength: jsii.Number(123),
//   		MinLength: jsii.Number(123),
//   	},
//   	ElementRequirementValidationRule: &X12ElementRequirementValidationRuleProperty{
//   		ElementPosition: jsii.String("elementPosition"),
//   		Requirement: jsii.String("requirement"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12validationrule.html
//
type CfnTransformer_X12ValidationRuleProperty struct {
	// Specifies a code list validation rule that modifies the allowed code values for a specific X12 element.
	//
	// This rule enables you to customize which codes are considered valid for an element, allowing for trading partner-specific code requirements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12validationrule.html#cfn-b2bi-transformer-x12validationrule-codelistvalidationrule
	//
	CodeListValidationRule interface{} `field:"optional" json:"codeListValidationRule" yaml:"codeListValidationRule"`
	// Specifies an element length validation rule that defines custom length constraints for a specific X12 element.
	//
	// This rule allows you to enforce minimum and maximum length requirements that may differ from the standard X12 specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12validationrule.html#cfn-b2bi-transformer-x12validationrule-elementlengthvalidationrule
	//
	ElementLengthValidationRule interface{} `field:"optional" json:"elementLengthValidationRule" yaml:"elementLengthValidationRule"`
	// Specifies an element requirement validation rule that modifies whether a specific X12 element is required or optional within a segment.
	//
	// This rule provides flexibility to accommodate different trading partner requirements for element presence.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12validationrule.html#cfn-b2bi-transformer-x12validationrule-elementrequirementvalidationrule
	//
	ElementRequirementValidationRule interface{} `field:"optional" json:"elementRequirementValidationRule" yaml:"elementRequirementValidationRule"`
}

