package awscdkapprunneralpha


// Attributes for the App Runner Auto Scaling Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   autoScalingConfigurationAttributes := &AutoScalingConfigurationAttributes{
//   	AutoScalingConfigurationName: jsii.String("autoScalingConfigurationName"),
//   	AutoScalingConfigurationRevision: jsii.Number(123),
//   }
//
// Experimental.
type AutoScalingConfigurationAttributes struct {
	// The name of the Auto Scaling Configuration.
	// Experimental.
	AutoScalingConfigurationName *string `field:"required" json:"autoScalingConfigurationName" yaml:"autoScalingConfigurationName"`
	// The revision of the Auto Scaling Configuration.
	// Experimental.
	AutoScalingConfigurationRevision *float64 `field:"required" json:"autoScalingConfigurationRevision" yaml:"autoScalingConfigurationRevision"`
}

