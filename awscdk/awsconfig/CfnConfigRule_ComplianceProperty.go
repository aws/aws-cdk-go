package awsconfig


// Compliance details of the Config rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   complianceProperty := &ComplianceProperty{
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-compliance.html
//
type CfnConfigRule_ComplianceProperty struct {
	// Compliance type determined by the Config rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-compliance.html#cfn-config-configrule-compliance-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

