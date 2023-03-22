package awsroute53recoverycontrol

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnControlPanel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnControlPanelProps := &cfnControlPanelProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	clusterArn: jsii.String("clusterArn"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnControlPanelProps struct {
	// The name of the control panel.
	//
	// You can use any non-white space character in the name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the cluster for the control panel.
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// The value for a tag.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

