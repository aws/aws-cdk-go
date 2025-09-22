package awsroute53recoverycontrol


// A reference to a ControlPanel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   controlPanelReference := &ControlPanelReference{
//   	ControlPanelArn: jsii.String("controlPanelArn"),
//   }
//
type ControlPanelReference struct {
	// The ControlPanelArn of the ControlPanel resource.
	ControlPanelArn *string `field:"required" json:"controlPanelArn" yaml:"controlPanelArn"`
}

