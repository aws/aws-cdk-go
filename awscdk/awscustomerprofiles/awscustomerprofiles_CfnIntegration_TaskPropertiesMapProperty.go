package awscustomerprofiles


// A map used to store task-related information.
//
// The execution service looks for particular information based on the `TaskType` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskPropertiesMapProperty := &taskPropertiesMapProperty{
//   	operatorPropertyKey: jsii.String("operatorPropertyKey"),
//   	property: jsii.String("property"),
//   }
//
type CfnIntegration_TaskPropertiesMapProperty struct {
	// The task property key.
	OperatorPropertyKey *string `field:"required" json:"operatorPropertyKey" yaml:"operatorPropertyKey"`
	// The task property value.
	Property *string `field:"required" json:"property" yaml:"property"`
}

