package previewawsroute53recoverycontrolmixins


// Properties for CfnRoutingControlPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRoutingControlMixinProps := &CfnRoutingControlMixinProps{
//   	ClusterArn: jsii.String("clusterArn"),
//   	ControlPanelArn: jsii.String("controlPanelArn"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html
//
type CfnRoutingControlMixinProps struct {
	// The Amazon Resource Name (ARN) of the cluster that hosts the routing control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html#cfn-route53recoverycontrol-routingcontrol-clusterarn
	//
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// The Amazon Resource Name (ARN) of the control panel that includes the routing control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html#cfn-route53recoverycontrol-routingcontrol-controlpanelarn
	//
	ControlPanelArn *string `field:"optional" json:"controlPanelArn" yaml:"controlPanelArn"`
	// The name of the routing control.
	//
	// You can use any non-white space character in the name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html#cfn-route53recoverycontrol-routingcontrol-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

