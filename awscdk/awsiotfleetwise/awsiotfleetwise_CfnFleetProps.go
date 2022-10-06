package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnFleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFleetProps := &cfnFleetProps{
//   	id: jsii.String("id"),
//   	signalCatalogArn: jsii.String("signalCatalogArn"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFleetProps struct {
	// `AWS::IoTFleetWise::Fleet.Id`.
	Id *string `field:"required" json:"id" yaml:"id"`
	// `AWS::IoTFleetWise::Fleet.SignalCatalogArn`.
	SignalCatalogArn *string `field:"required" json:"signalCatalogArn" yaml:"signalCatalogArn"`
	// `AWS::IoTFleetWise::Fleet.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::IoTFleetWise::Fleet.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

