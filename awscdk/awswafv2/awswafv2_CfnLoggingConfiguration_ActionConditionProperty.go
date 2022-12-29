package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionConditionProperty := &actionConditionProperty{
//   	action: jsii.String("action"),
//   }
//
type CfnLoggingConfiguration_ActionConditionProperty struct {
	// `CfnLoggingConfiguration.ActionConditionProperty.Action`.
	Action *string `field:"required" json:"action" yaml:"action"`
}

