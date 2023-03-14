package awswafv2


// Filtering that specifies which web requests are kept in the logs and which are dropped, defined for a web ACL's `LoggingConfiguration` .
//
// You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingFilterProperty := &LoggingFilterProperty{
//   	DefaultBehavior: jsii.String("defaultBehavior"),
//   	Filters: []interface{}{
//   		&FilterProperty{
//   			Behavior: jsii.String("behavior"),
//   			Conditions: []interface{}{
//   				&ConditionProperty{
//   					ActionCondition: &ActionConditionProperty{
//   						Action: jsii.String("action"),
//   					},
//   					LabelNameCondition: &LabelNameConditionProperty{
//   						LabelName: jsii.String("labelName"),
//   					},
//   				},
//   			},
//   			Requirement: jsii.String("requirement"),
//   		},
//   	},
//   }
//
type CfnLoggingConfiguration_LoggingFilterProperty struct {
	// Default handling for logs that don't match any of the specified filtering conditions.
	DefaultBehavior *string `field:"required" json:"defaultBehavior" yaml:"defaultBehavior"`
	// The filters that you want to apply to the logs.
	Filters interface{} `field:"required" json:"filters" yaml:"filters"`
}

