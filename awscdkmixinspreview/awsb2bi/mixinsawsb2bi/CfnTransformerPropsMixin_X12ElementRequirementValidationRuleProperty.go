package mixinsawsb2bi


// Defines a validation rule that modifies the requirement status of a specific X12 element within a segment.
//
// This rule allows you to make optional elements mandatory or mandatory elements optional, providing flexibility to accommodate different trading partner requirements and business rules. The rule targets a specific element position within a segment and sets its requirement status to either OPTIONAL or MANDATORY.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   x12ElementRequirementValidationRuleProperty := &X12ElementRequirementValidationRuleProperty{
//   	ElementPosition: jsii.String("elementPosition"),
//   	Requirement: jsii.String("requirement"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12elementrequirementvalidationrule.html
//
type CfnTransformerPropsMixin_X12ElementRequirementValidationRuleProperty struct {
	// Specifies the position of the element within an X12 segment for which the requirement status will be modified.
	//
	// The format follows the pattern of segment identifier followed by element position (e.g., "ST-01" for the first element of the ST segment).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12elementrequirementvalidationrule.html#cfn-b2bi-transformer-x12elementrequirementvalidationrule-elementposition
	//
	ElementPosition *string `field:"optional" json:"elementPosition" yaml:"elementPosition"`
	// Specifies the requirement status for the element at the specified position.
	//
	// Valid values are OPTIONAL (the element may be omitted) or MANDATORY (the element must be present).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12elementrequirementvalidationrule.html#cfn-b2bi-transformer-x12elementrequirementvalidationrule-requirement
	//
	Requirement *string `field:"optional" json:"requirement" yaml:"requirement"`
}

