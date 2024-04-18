package awsentityresolution


// Properties for defining a `CfnPolicyStatement`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPolicyStatementProps := &CfnPolicyStatementProps{
//   	Arn: jsii.String("arn"),
//   	StatementId: jsii.String("statementId"),
//
//   	// the properties below are optional
//   	Action: []*string{
//   		jsii.String("action"),
//   	},
//   	Condition: jsii.String("condition"),
//   	Effect: jsii.String("effect"),
//   	Principal: []*string{
//   		jsii.String("principal"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-policystatement.html
//
type CfnPolicyStatementProps struct {
	// The Amazon Resource Name (ARN) of the resource that will be accessed by the principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-policystatement.html#cfn-entityresolution-policystatement-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// A statement identifier that differentiates the statement from others in the same policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-policystatement.html#cfn-entityresolution-policystatement-statementid
	//
	StatementId *string `field:"required" json:"statementId" yaml:"statementId"`
	// The action that the principal can use on the resource.
	//
	// For example, `entityresolution:GetIdMappingJob` , `entityresolution:GetMatchingJob` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-policystatement.html#cfn-entityresolution-policystatement-action
	//
	Action *[]*string `field:"optional" json:"action" yaml:"action"`
	// A set of condition keys that you can use in key policies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-policystatement.html#cfn-entityresolution-policystatement-condition
	//
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Determines whether the permissions specified in the policy are to be allowed ( `Allow` ) or denied ( `Deny` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-policystatement.html#cfn-entityresolution-policystatement-effect
	//
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// The AWS service or AWS account that can access the resource defined as ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-policystatement.html#cfn-entityresolution-policystatement-principal
	//
	Principal *[]*string `field:"optional" json:"principal" yaml:"principal"`
}

