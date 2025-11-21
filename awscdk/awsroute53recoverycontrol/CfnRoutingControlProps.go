package awsroute53recoverycontrol


// Properties for defining a `CfnRoutingControl`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRoutingControlProps := &CfnRoutingControlProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ClusterArn: jsii.String("clusterArn"),
//   	ControlPanelArn: jsii.String("controlPanelArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html
//
type CfnRoutingControlProps struct {
	// The name of the routing control.
	//
	// You can use any non-white space character in the name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html#cfn-route53recoverycontrol-routingcontrol-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the cluster that hosts the routing control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html#cfn-route53recoverycontrol-routingcontrol-clusterarn
	//
	ClusterArn interface{} `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// The Amazon Resource Name (ARN) of the control panel that includes the routing control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html#cfn-route53recoverycontrol-routingcontrol-controlpanelarn
	//
	ControlPanelArn interface{} `field:"optional" json:"controlPanelArn" yaml:"controlPanelArn"`
}

