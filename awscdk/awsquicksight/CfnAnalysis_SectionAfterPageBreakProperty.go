package awsquicksight


// The configuration of a page break after a section.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sectionAfterPageBreakProperty := &SectionAfterPageBreakProperty{
//   	Status: jsii.String("status"),
//   }
//
type CfnAnalysis_SectionAfterPageBreakProperty struct {
	// The option that enables or disables a page break at the end of a section.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

