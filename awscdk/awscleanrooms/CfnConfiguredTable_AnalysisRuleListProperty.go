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
//
//   	// the properties below are optional
//   	AllowedJoinOperators: []*string{
//   		jsii.String("allowedJoinOperators"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisrulelist.html
//
type CfnConfiguredTable_AnalysisRuleListProperty struct {
	// Columns that can be used to join a configured table with the table of the member who can query and other members' configured tables.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisrulelist.html#cfn-cleanrooms-configuredtable-analysisrulelist-joincolumns
	//
	JoinColumns *[]*string `field:"required" json:"joinColumns" yaml:"joinColumns"`
	// Columns that can be listed in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisrulelist.html#cfn-cleanrooms-configuredtable-analysisrulelist-listcolumns
	//
	ListColumns *[]*string `field:"required" json:"listColumns" yaml:"listColumns"`
	// Which logical operators (if any) are to be used in an INNER JOIN match condition.
	//
	// Default is `AND` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisrulelist.html#cfn-cleanrooms-configuredtable-analysisrulelist-allowedjoinoperators
	//
	AllowedJoinOperators *[]*string `field:"optional" json:"allowedJoinOperators" yaml:"allowedJoinOperators"`
}

