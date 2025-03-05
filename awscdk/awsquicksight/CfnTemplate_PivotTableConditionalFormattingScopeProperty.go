package awsquicksight


// The scope of the cell for conditional formatting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTableConditionalFormattingScopeProperty := &PivotTableConditionalFormattingScopeProperty{
//   	Role: jsii.String("role"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableconditionalformattingscope.html
//
type CfnTemplate_PivotTableConditionalFormattingScopeProperty struct {
	// The role (field, field total, grand total) of the cell for conditional formatting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableconditionalformattingscope.html#cfn-quicksight-template-pivottableconditionalformattingscope-role
	//
	Role *string `field:"optional" json:"role" yaml:"role"`
}

