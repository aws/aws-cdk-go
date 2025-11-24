package mixinsawsverifiedpermissions


// A structure that defines a static policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   staticPolicyDefinitionProperty := &StaticPolicyDefinitionProperty{
//   	Description: jsii.String("description"),
//   	Statement: jsii.String("statement"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-staticpolicydefinition.html
//
type CfnPolicyPropsMixin_StaticPolicyDefinitionProperty struct {
	// The description of the static policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-staticpolicydefinition.html#cfn-verifiedpermissions-policy-staticpolicydefinition-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The policy content of the static policy, written in the Cedar policy language.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-staticpolicydefinition.html#cfn-verifiedpermissions-policy-staticpolicydefinition-statement
	//
	Statement *string `field:"optional" json:"statement" yaml:"statement"`
}

