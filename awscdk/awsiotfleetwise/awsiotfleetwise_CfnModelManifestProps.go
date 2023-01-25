package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnModelManifest`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnModelManifestProps := &cfnModelManifestProps{
//   	name: jsii.String("name"),
//   	signalCatalogArn: jsii.String("signalCatalogArn"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	nodes: []*string{
//   		jsii.String("nodes"),
//   	},
//   	status: jsii.String("status"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnModelManifestProps struct {
	// The name of the vehicle model.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the signal catalog associated with the vehicle model.
	SignalCatalogArn *string `field:"required" json:"signalCatalogArn" yaml:"signalCatalogArn"`
	// A brief description of the vehicle model.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::IoTFleetWise::ModelManifest.Nodes`.
	Nodes *[]*string `field:"optional" json:"nodes" yaml:"nodes"`
	// The state of the vehicle model.
	//
	// If the status is `ACTIVE` , the vehicle model can't be edited. If the status is `DRAFT` , you can edit the vehicle model.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// `AWS::IoTFleetWise::ModelManifest.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

