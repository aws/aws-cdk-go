package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customActionURLOperationProperty := &CustomActionURLOperationProperty{
//   	UrlTarget: jsii.String("urlTarget"),
//   	UrlTemplate: jsii.String("urlTemplate"),
//   }
//
type CfnAnalysis_CustomActionURLOperationProperty struct {
	// `CfnAnalysis.CustomActionURLOperationProperty.URLTarget`.
	UrlTarget *string `field:"required" json:"urlTarget" yaml:"urlTarget"`
	// `CfnAnalysis.CustomActionURLOperationProperty.URLTemplate`.
	UrlTemplate *string `field:"required" json:"urlTemplate" yaml:"urlTemplate"`
}

