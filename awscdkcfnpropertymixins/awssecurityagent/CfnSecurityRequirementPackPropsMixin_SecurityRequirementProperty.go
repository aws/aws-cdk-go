package awssecurityagent


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   securityRequirementProperty := &SecurityRequirementProperty{
//   	Description: jsii.String("description"),
//   	Domain: jsii.String("domain"),
//   	Evaluation: jsii.String("evaluation"),
//   	Name: jsii.String("name"),
//   	Remediation: jsii.String("remediation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-securityrequirementpack-securityrequirement.html
//
type CfnSecurityRequirementPackPropsMixin_SecurityRequirementProperty struct {
	// Description of the security requirement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-securityrequirementpack-securityrequirement.html#cfn-securityagent-securityrequirementpack-securityrequirement-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Security domain this requirement belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-securityrequirementpack-securityrequirement.html#cfn-securityagent-securityrequirementpack-securityrequirement-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// How to evaluate compliance with this requirement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-securityrequirementpack-securityrequirement.html#cfn-securityagent-securityrequirementpack-securityrequirement-evaluation
	//
	Evaluation *string `field:"optional" json:"evaluation" yaml:"evaluation"`
	// Name of the security requirement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-securityrequirementpack-securityrequirement.html#cfn-securityagent-securityrequirementpack-securityrequirement-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// How to remediate non-compliance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-securityrequirementpack-securityrequirement.html#cfn-securityagent-securityrequirementpack-securityrequirement-remediation
	//
	Remediation *string `field:"optional" json:"remediation" yaml:"remediation"`
}

