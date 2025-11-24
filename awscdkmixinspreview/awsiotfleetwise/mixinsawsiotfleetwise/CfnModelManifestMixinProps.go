package mixinsawsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnModelManifestPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnModelManifestMixinProps := &CfnModelManifestMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Nodes: []*string{
//   		jsii.String("nodes"),
//   	},
//   	SignalCatalogArn: jsii.String("signalCatalogArn"),
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-modelmanifest.html
//
type CfnModelManifestMixinProps struct {
	// A brief description of the vehicle model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-modelmanifest.html#cfn-iotfleetwise-modelmanifest-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the vehicle model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-modelmanifest.html#cfn-iotfleetwise-modelmanifest-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of nodes, which are a general abstraction of signals.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-modelmanifest.html#cfn-iotfleetwise-modelmanifest-nodes
	//
	Nodes *[]*string `field:"optional" json:"nodes" yaml:"nodes"`
	// The Amazon Resource Name (ARN) of the signal catalog associated with the vehicle model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-modelmanifest.html#cfn-iotfleetwise-modelmanifest-signalcatalogarn
	//
	SignalCatalogArn *string `field:"optional" json:"signalCatalogArn" yaml:"signalCatalogArn"`
	// The state of the vehicle model.
	//
	// If the status is `ACTIVE` , the vehicle model can't be edited. If the status is `DRAFT` , you can edit the vehicle model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-modelmanifest.html#cfn-iotfleetwise-modelmanifest-status
	//
	// Default: - "DRAFT".
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Metadata that can be used to manage the vehicle model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-modelmanifest.html#cfn-iotfleetwise-modelmanifest-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

