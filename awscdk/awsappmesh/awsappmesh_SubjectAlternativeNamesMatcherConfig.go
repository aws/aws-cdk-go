package awsappmesh


// All Properties for Subject Alternative Names Matcher for both Client Policy and Listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subjectAlternativeNamesMatcherConfig := &subjectAlternativeNamesMatcherConfig{
//   	subjectAlternativeNamesMatch: &subjectAlternativeNameMatchersProperty{
//   		exact: []*string{
//   			jsii.String("exact"),
//   		},
//   	},
//   }
//
type SubjectAlternativeNamesMatcherConfig struct {
	// VirtualNode CFN configuration for subject alternative names secured by the certificate.
	SubjectAlternativeNamesMatch *CfnVirtualNode_SubjectAlternativeNameMatchersProperty `field:"required" json:"subjectAlternativeNamesMatch" yaml:"subjectAlternativeNamesMatch"`
}

