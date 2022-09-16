package awscontroltower


// Properties for defining a `CfnEnabledControl`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnabledControlProps := &cfnEnabledControlProps{
//   	controlIdentifier: jsii.String("controlIdentifier"),
//   	targetIdentifier: jsii.String("targetIdentifier"),
//   }
//
type CfnEnabledControlProps struct {
	// `AWS::ControlTower::EnabledControl.ControlIdentifier`.
	ControlIdentifier *string `field:"required" json:"controlIdentifier" yaml:"controlIdentifier"`
	// `AWS::ControlTower::EnabledControl.TargetIdentifier`.
	TargetIdentifier *string `field:"required" json:"targetIdentifier" yaml:"targetIdentifier"`
}

