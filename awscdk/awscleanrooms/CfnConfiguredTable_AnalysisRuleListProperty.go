package awscleanrooms


// A type of analysis rule that enables row-level analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisRuleListProperty := &AnalysisRuleListProperty{
//   	JoinColumns: []*string{
//   		jsii.String("joinColumns"),
//   	},
//   	ListColumns: []*string{
//   		jsii.String("listColumns"),
//   	},
//   }
//
type CfnConfiguredTable_AnalysisRuleListProperty struct {
	// Columns that can be used to join a configured table with the table of the member who can query and other members' configured tables.
	JoinColumns *[]*string `field:"required" json:"joinColumns" yaml:"joinColumns"`
	// Columns that can be listed in the output.
	ListColumns *[]*string `field:"required" json:"listColumns" yaml:"listColumns"`
}

