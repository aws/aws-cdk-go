package awsquicksight


// The configuration for a custom label on a `ReferenceLine` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceLineCustomLabelConfigurationProperty := &ReferenceLineCustomLabelConfigurationProperty{
//   	CustomLabel: jsii.String("customLabel"),
//   }
//
type CfnDashboard_ReferenceLineCustomLabelConfigurationProperty struct {
	// The string text of the custom label.
	CustomLabel *string `field:"required" json:"customLabel" yaml:"customLabel"`
}

