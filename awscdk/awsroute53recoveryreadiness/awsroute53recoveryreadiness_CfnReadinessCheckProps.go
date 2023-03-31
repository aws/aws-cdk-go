package awsroute53recoveryreadiness

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnReadinessCheck`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReadinessCheckProps := &cfnReadinessCheckProps{
//   	readinessCheckName: jsii.String("readinessCheckName"),
//   	resourceSetName: jsii.String("resourceSetName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

