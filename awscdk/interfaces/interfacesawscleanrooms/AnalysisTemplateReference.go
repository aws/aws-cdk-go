package interfacesawscleanrooms


// A reference to a AnalysisTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisTemplateReference := &AnalysisTemplateReference{
//   	AnalysisTemplateArn: jsii.String("analysisTemplateArn"),
//   	AnalysisTemplateIdentifier: jsii.String("analysisTemplateIdentifier"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   }
//
type AnalysisTemplateReference struct {
	// The ARN of the AnalysisTemplate resource.
	AnalysisTemplateArn *string `field:"required" json:"analysisTemplateArn" yaml:"analysisTemplateArn"`
	// The AnalysisTemplateIdentifier of the AnalysisTemplate resource.
	AnalysisTemplateIdentifier *string `field:"required" json:"analysisTemplateIdentifier" yaml:"analysisTemplateIdentifier"`
	// The MembershipIdentifier of the AnalysisTemplate resource.
	MembershipIdentifier *string `field:"required" json:"membershipIdentifier" yaml:"membershipIdentifier"`
}

