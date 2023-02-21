package awsroute53recoverycontrol


// Properties for defining a `CfnRoutingControl`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRoutingControlProps := &cfnRoutingControlProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	clusterArn: jsii.String("clusterArn"),
//   	controlPanelArn: jsii.String("controlPanelArn"),
//   }
//
type CfnRoutingControlProps struct {
	// The name of the routing control.
	//
	// You can use any non-white space character in the name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the cluster that includes the routing control.
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// The Amazon Resource Name (ARN) of the control panel that includes the routing control.
	ControlPanelArn *string `field:"optional" json:"controlPanelArn" yaml:"controlPanelArn"`
}

