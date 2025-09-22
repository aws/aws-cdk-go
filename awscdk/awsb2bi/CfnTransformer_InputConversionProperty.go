package awsb2bi


// Contains the input formatting options for an inbound transformer (takes an X12-formatted EDI document as input and converts it to JSON or XML.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputConversionProperty := &InputConversionProperty{
//   	FromFormat: jsii.String("fromFormat"),
//
//   	// the properties below are optional
//   	AdvancedOptions: &AdvancedOptionsProperty{
//   		X12: &X12AdvancedOptionsProperty{
//   			SplitOptions: &X12SplitOptionsProperty{
//   				SplitBy: jsii.String("splitBy"),
//   			},
//   			ValidationOptions: &X12ValidationOptionsProperty{
//   				ValidationRules: []interface{}{
//   					&X12ValidationRuleProperty{
//   						CodeListValidationRule: &X12CodeListValidationRuleProperty{
//   							ElementId: jsii.String("elementId"),
//
//   							// the properties below are optional
//   							CodesToAdd: []*string{
//   								jsii.String("codesToAdd"),
//   							},
//   							CodesToRemove: []*string{
//   								jsii.String("codesToRemove"),
//   							},
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-inputconversion.html
//
type CfnTransformer_InputConversionProperty struct {
	// The format for the transformer input: currently on `X12` is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-inputconversion.html#cfn-b2bi-transformer-inputconversion-fromformat
	//
	FromFormat *string `field:"required" json:"fromFormat" yaml:"fromFormat"`
	// Specifies advanced options for the input conversion process.
	//
	// These options provide additional control over how EDI files are processed during transformation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-inputconversion.html#cfn-b2bi-transformer-inputconversion-advancedoptions
	//
	AdvancedOptions interface{} `field:"optional" json:"advancedOptions" yaml:"advancedOptions"`
	// A structure that contains the formatting options for an inbound transformer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-inputconversion.html#cfn-b2bi-transformer-inputconversion-formatoptions
	//
	FormatOptions interface{} `field:"optional" json:"formatOptions" yaml:"formatOptions"`
}

