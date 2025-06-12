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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sectionafterpagebreak.html
//
type CfnTemplate_SectionAfterPageBreakProperty struct {
	// The option that enables or disables a page break at the end of a section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sectionafterpagebreak.html#cfn-quicksight-template-sectionafterpagebreak-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

