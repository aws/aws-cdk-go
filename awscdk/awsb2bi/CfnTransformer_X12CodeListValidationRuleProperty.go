package awsb2bi


// Code list validation rule configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   x12CodeListValidationRuleProperty := &X12CodeListValidationRuleProperty{
//   	ElementId: jsii.String("elementId"),
//
//   	// the properties below are optional
//   	CodesToAdd: []*string{
//   		jsii.String("codesToAdd"),
//   	},
//   	CodesToRemove: []*string{
//   		jsii.String("codesToRemove"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12codelistvalidationrule.html
//
type CfnTransformer_X12CodeListValidationRuleProperty struct {
	// Specifies the four-digit element ID to which the code list modifications apply.
	//
	// This identifies which X12 element will have its allowed code values modified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12codelistvalidationrule.html#cfn-b2bi-transformer-x12codelistvalidationrule-elementid
	//
	ElementId *string `field:"required" json:"elementId" yaml:"elementId"`
	// Specifies a list of code values to add to the element's allowed values.
	//
	// These codes will be considered valid for the specified element in addition to the standard codes defined by the X12 specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12codelistvalidationrule.html#cfn-b2bi-transformer-x12codelistvalidationrule-codestoadd
	//
	CodesToAdd *[]*string `field:"optional" json:"codesToAdd" yaml:"codesToAdd"`
	// Specifies a list of code values to remove from the element's allowed values.
	//
	// These codes will be considered invalid for the specified element, even if they are part of the standard codes defined by the X12 specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12codelistvalidationrule.html#cfn-b2bi-transformer-x12codelistvalidationrule-codestoremove
	//
	CodesToRemove *[]*string `field:"optional" json:"codesToRemove" yaml:"codesToRemove"`
}

