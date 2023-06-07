package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFleetProps := &CfnFleetProps{
//   	Id: jsii.String("id"),
//   	SignalCatalogArn: jsii.String("signalCatalogArn"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFleetProps struct {
	// The unique ID of the fleet.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The ARN of the signal catalog associated with the fleet.
	SignalCatalogArn *string `field:"required" json:"signalCatalogArn" yaml:"signalCatalogArn"`
	// (Optional) A brief description of the fleet.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// (Optional) Metadata that can be used to manage the fleet.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

