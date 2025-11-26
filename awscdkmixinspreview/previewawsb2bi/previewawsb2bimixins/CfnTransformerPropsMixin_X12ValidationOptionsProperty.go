package previewawsb2bimixins


// Contains configuration options for X12 EDI validation.
//
// This structure allows you to specify custom validation rules that will be applied during EDI document processing, including element length constraints, code list modifications, and element requirement changes. These validation options provide flexibility to accommodate trading partner-specific requirements while maintaining EDI compliance. The validation rules are applied in addition to standard X12 validation to ensure documents meet both standard and custom requirements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   x12ValidationOptionsProperty := &X12ValidationOptionsProperty{
//   	ValidationRules: []interface{}{
//   		&X12ValidationRuleProperty{
//   			CodeListValidationRule: &X12CodeListValidationRuleProperty{
//   				CodesToAdd: []*string{
//   					jsii.String("codesToAdd"),
//   				},
//   				CodesToRemove: []*string{
//   					jsii.String("codesToRemove"),
//   				},
//   				ElementId: jsii.String("elementId"),
//   			},
//   			ElementLengthValidationRule: &X12ElementLengthValidationRuleProperty{
//   				ElementId: jsii.String("elementId"),
//   				MaxLength: jsii.Number(123),
//   				MinLength: jsii.Number(123),
//   			},
//   			ElementRequirementValidationRule: &X12ElementRequirementValidationRuleProperty{
//   				ElementPosition: jsii.String("elementPosition"),
//   				Requirement: jsii.String("requirement"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12validationoptions.html
//
type CfnTransformerPropsMixin_X12ValidationOptionsProperty struct {
	// Specifies a list of validation rules to apply during EDI document processing.
	//
	// These rules can include code list modifications, element length constraints, and element requirement changes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12validationoptions.html#cfn-b2bi-transformer-x12validationoptions-validationrules
	//
	ValidationRules interface{} `field:"optional" json:"validationRules" yaml:"validationRules"`
}

