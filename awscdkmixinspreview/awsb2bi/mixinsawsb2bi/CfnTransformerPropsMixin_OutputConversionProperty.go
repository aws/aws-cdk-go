package mixinsawsb2bi


// Contains the formatting options for an outbound transformer (takes JSON or XML as input and converts it to an EDI document (currently only X12 format is supported).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outputConversionProperty := &OutputConversionProperty{
//   	AdvancedOptions: &AdvancedOptionsProperty{
//   		X12: &X12AdvancedOptionsProperty{
//   			SplitOptions: &X12SplitOptionsProperty{
//   				SplitBy: jsii.String("splitBy"),
//   			},
//   			ValidationOptions: &X12ValidationOptionsProperty{
//   				ValidationRules: []interface{}{
//   					&X12ValidationRuleProperty{
//   						CodeListValidationRule: &X12CodeListValidationRuleProperty{
//   							CodesToAdd: []*string{
//   								jsii.String("codesToAdd"),
//   							},
//   							CodesToRemove: []*string{
//   								jsii.String("codesToRemove"),
//   							},
//   							ElementId: jsii.String("elementId"),
//   						},
//   						ElementLengthValidationRule: &X12ElementLengthValidationRuleProperty{
//   							ElementId: jsii.String("elementId"),
//   							MaxLength: jsii.Number(123),
//   							MinLength: jsii.Number(123),
//   						},
//   						ElementRequirementValidationRule: &X12ElementRequirementValidationRuleProperty{
//   							ElementPosition: jsii.String("elementPosition"),
//   							Requirement: jsii.String("requirement"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	FormatOptions: &FormatOptionsProperty{
//   		X12: &X12DetailsProperty{
//   			TransactionSet: jsii.String("transactionSet"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   	ToFormat: jsii.String("toFormat"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-outputconversion.html
//
type CfnTransformerPropsMixin_OutputConversionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-outputconversion.html#cfn-b2bi-transformer-outputconversion-advancedoptions
	//
	AdvancedOptions interface{} `field:"optional" json:"advancedOptions" yaml:"advancedOptions"`
	// A structure that contains the X12 transaction set and version for the transformer output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-outputconversion.html#cfn-b2bi-transformer-outputconversion-formatoptions
	//
	FormatOptions interface{} `field:"optional" json:"formatOptions" yaml:"formatOptions"`
	// The format for the output from an outbound transformer: only X12 is currently supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-outputconversion.html#cfn-b2bi-transformer-outputconversion-toformat
	//
	ToFormat *string `field:"optional" json:"toFormat" yaml:"toFormat"`
}

