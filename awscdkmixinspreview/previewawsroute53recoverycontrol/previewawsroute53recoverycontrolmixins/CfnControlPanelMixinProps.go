package previewawsroute53recoverycontrolmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnControlPanelPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnControlPanelMixinProps := &CfnControlPanelMixinProps{
//   	ClusterArn: jsii.String("clusterArn"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-controlpanel.html
//
type CfnControlPanelMixinProps struct {
	// The Amazon Resource Name (ARN) of the cluster for the control panel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-controlpanel.html#cfn-route53recoverycontrol-controlpanel-clusterarn
	//
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// The name of the control panel.
	//
	// You can use any non-white space character in the name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-controlpanel.html#cfn-route53recoverycontrol-controlpanel-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags associated with the control panel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-controlpanel.html#cfn-route53recoverycontrol-controlpanel-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

