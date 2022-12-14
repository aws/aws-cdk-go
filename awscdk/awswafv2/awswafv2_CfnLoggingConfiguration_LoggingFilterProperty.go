package awswafv2


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
	// `CfnLoggingConfiguration.LoggingFilterProperty.DefaultBehavior`.
	DefaultBehavior *string `field:"required" json:"defaultBehavior" yaml:"defaultBehavior"`
	// `CfnLoggingConfiguration.LoggingFilterProperty.Filters`.
	Filters interface{} `field:"required" json:"filters" yaml:"filters"`
}

