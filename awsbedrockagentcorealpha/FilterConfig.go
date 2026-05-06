package awsbedrockagentcorealpha


// Filter configuration for online evaluation.
//
// Filters determine which agent traces should be included in the evaluation
// based on trace properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var filterOperator FilterOperator
//   var filterValue FilterValue
//
//   filterConfig := &FilterConfig{
//   	Key: jsii.String("key"),
//   	Operator: filterOperator,
//   	Value: filterValue,
//   }
//
// Experimental.
type FilterConfig struct {
	// The key or field name to filter on within the agent trace data.
	//
	// Example:
	//   "user.region"
	//
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The comparison operator to use for filtering.
	// Experimental.
	Operator FilterOperator `field:"required" json:"operator" yaml:"operator"`
	// The value to compare against using the specified operator.
	//
	// Use `FilterValue.string()`, `FilterValue.number()`, or `FilterValue.boolean()`
	// to create typed filter values.
	// Experimental.
	Value FilterValue `field:"required" json:"value" yaml:"value"`
}

