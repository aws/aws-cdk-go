package awsroute53recoverycontrol

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &CfnClusterProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnClusterProps struct {
	// Name of the cluster.
	//
	// You can use any non-white space character in the name except the following: & > < ' (single quote) " (double quote) ; (semicolon).
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value for a tag.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

