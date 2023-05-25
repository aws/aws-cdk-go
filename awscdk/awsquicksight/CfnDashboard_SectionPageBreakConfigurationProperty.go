package awsquicksight


// The configuration of a page break for a section.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sectionPageBreakConfigurationProperty := &SectionPageBreakConfigurationProperty{
//   	After: &SectionAfterPageBreakProperty{
//   		Status: jsii.String("status"),
//   	},
//   }
//
type CfnDashboard_SectionPageBreakConfigurationProperty struct {
	// The configuration of a page break after a section.
	After interface{} `field:"optional" json:"after" yaml:"after"`
}

