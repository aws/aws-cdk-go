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
//   loggingFilterProperty := &loggingFilterProperty{
//   	defaultBehavior: jsii.String("defaultBehavior"),
//   	filters: []interface{}{
//   		&filterProperty{
//   			behavior: jsii.String("behavior"),
//   			conditions: []interface{}{
//   				&conditionProperty{
//   					actionCondition: &actionConditionProperty{
//   						action: jsii.String("action"),
//   					},
//   					labelNameCondition: &labelNameConditionProperty{
//   						labelName: jsii.String("labelName"),
//   					},
//   				},
//   			},
//   			requirement: jsii.String("requirement"),
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

