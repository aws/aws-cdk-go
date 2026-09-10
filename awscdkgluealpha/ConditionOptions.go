package awscdkgluealpha


// Options shared by all trigger conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   conditionOptions := &ConditionOptions{
//   	LogicalOperator: glue_alpha.ConditionLogicalOperator_EQUALS,
//   }
//
// Experimental.
type ConditionOptions struct {
	// The logical operator for the condition.
	// Default: ConditionLogicalOperator.EQUALS
	//
	// Experimental.
	LogicalOperator ConditionLogicalOperator `field:"optional" json:"logicalOperator" yaml:"logicalOperator"`
}

