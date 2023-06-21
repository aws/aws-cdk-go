package awsroute53recoveryreadiness

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnReadinessCheck`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReadinessCheckProps := &CfnReadinessCheckProps{
//   	ReadinessCheckName: jsii.String("readinessCheckName"),
//   	ResourceSetName: jsii.String("resourceSetName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnReadinessCheckProps struct {
	// The name of the readiness check to create.
	ReadinessCheckName *string `field:"optional" json:"readinessCheckName" yaml:"readinessCheckName"`
	// The name of the resource set to check.
	ResourceSetName *string `field:"optional" json:"resourceSetName" yaml:"resourceSetName"`
	// A collection of tags associated with a resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

