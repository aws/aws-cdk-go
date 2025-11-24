package mixinsawsb2bi


// Contains advanced options specific to X12 EDI processing, such as splitting large X12 files into smaller units.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   x12AdvancedOptionsProperty := &X12AdvancedOptionsProperty{
//   	SplitOptions: &X12SplitOptionsProperty{
//   		SplitBy: jsii.String("splitBy"),
//   	},
//   	ValidationOptions: &X12ValidationOptionsProperty{
//   		ValidationRules: []interface{}{
//   			&X12ValidationRuleProperty{
//   				CodeListValidationRule: &X12CodeListValidationRuleProperty{
//   					CodesToAdd: []*string{
//   						jsii.String("codesToAdd"),
//   					},
//   					CodesToRemove: []*string{
//   						jsii.String("codesToRemove"),
//   					},
//   					ElementId: jsii.String("elementId"),
//   				},
//   				ElementLengthValidationRule: &X12ElementLengthValidationRuleProperty{
//   					ElementId: jsii.String("elementId"),
//   					MaxLength: jsii.Number(123),
//   					MinLength: jsii.Number(123),
//   				},
//   				ElementRequirementValidationRule: &X12ElementRequirementValidationRuleProperty{
//   					ElementPosition: jsii.String("elementPosition"),
//   					Requirement: jsii.String("requirement"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12advancedoptions.html
//
type CfnTransformerPropsMixin_X12AdvancedOptionsProperty struct {
	// Specifies options for splitting X12 EDI files.
	//
	// These options control how large X12 files are divided into smaller, more manageable units.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12advancedoptions.html#cfn-b2bi-transformer-x12advancedoptions-splitoptions
	//
	SplitOptions interface{} `field:"optional" json:"splitOptions" yaml:"splitOptions"`
	// Specifies validation options for X12 EDI processing.
	//
	// These options control how validation rules are applied during EDI document processing, including custom validation rules for element length constraints, code list validations, and element requirement checks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12advancedoptions.html#cfn-b2bi-transformer-x12advancedoptions-validationoptions
	//
	ValidationOptions interface{} `field:"optional" json:"validationOptions" yaml:"validationOptions"`
}

