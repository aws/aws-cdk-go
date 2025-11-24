package mixinsawsb2bi


// A structure that contains advanced options for EDI processing.
//
// Currently, only X12 advanced options are supported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   advancedOptionsProperty := &AdvancedOptionsProperty{
//   	X12: &X12AdvancedOptionsProperty{
//   		SplitOptions: &X12SplitOptionsProperty{
//   			SplitBy: jsii.String("splitBy"),
//   		},
//   		ValidationOptions: &X12ValidationOptionsProperty{
//   			ValidationRules: []interface{}{
//   				&X12ValidationRuleProperty{
//   					CodeListValidationRule: &X12CodeListValidationRuleProperty{
//   						CodesToAdd: []*string{
//   							jsii.String("codesToAdd"),
//   						},
//   						CodesToRemove: []*string{
//   							jsii.String("codesToRemove"),
//   						},
//   						ElementId: jsii.String("elementId"),
//   					},
//   					ElementLengthValidationRule: &X12ElementLengthValidationRuleProperty{
//   						ElementId: jsii.String("elementId"),
//   						MaxLength: jsii.Number(123),
//   						MinLength: jsii.Number(123),
//   					},
//   					ElementRequirementValidationRule: &X12ElementRequirementValidationRuleProperty{
//   						ElementPosition: jsii.String("elementPosition"),
//   						Requirement: jsii.String("requirement"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-advancedoptions.html
//
type CfnTransformerPropsMixin_AdvancedOptionsProperty struct {
	// A structure that contains X12-specific advanced options, such as split options for processing X12 EDI files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-advancedoptions.html#cfn-b2bi-transformer-advancedoptions-x12
	//
	X12 interface{} `field:"optional" json:"x12" yaml:"x12"`
}

