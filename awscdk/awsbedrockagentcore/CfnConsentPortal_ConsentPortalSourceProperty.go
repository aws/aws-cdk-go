package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   consentPortalSourceProperty := &ConsentPortalSourceProperty{
//   	Identifier: jsii.String("identifier"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-consentportal-consentportalsource.html
//
type CfnConsentPortal_ConsentPortalSourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-consentportal-consentportalsource.html#cfn-bedrockagentcore-consentportal-consentportalsource-identifier
	//
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-consentportal-consentportalsource.html#cfn-bedrockagentcore-consentportal-consentportalsource-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

