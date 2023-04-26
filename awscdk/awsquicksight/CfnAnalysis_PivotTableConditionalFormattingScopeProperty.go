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
type CfnAnalysis_PivotTableConditionalFormattingScopeProperty struct {
	// The role (field, field total, grand total) of the cell for conditional formatting.
	Role *string `field:"optional" json:"role" yaml:"role"`
}

