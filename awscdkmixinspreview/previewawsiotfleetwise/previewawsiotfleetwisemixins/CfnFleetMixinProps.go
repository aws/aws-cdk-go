package previewawsiotfleetwisemixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFleetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFleetMixinProps := &CfnFleetMixinProps{
//   	Description: jsii.String("description"),
//   	Id: jsii.String("id"),
//   	SignalCatalogArn: jsii.String("signalCatalogArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-fleet.html
//
type CfnFleetMixinProps struct {
	// A brief description of the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-fleet.html#cfn-iotfleetwise-fleet-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique ID of the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-fleet.html#cfn-iotfleetwise-fleet-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ARN of the signal catalog associated with the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-fleet.html#cfn-iotfleetwise-fleet-signalcatalogarn
	//
	SignalCatalogArn *string `field:"optional" json:"signalCatalogArn" yaml:"signalCatalogArn"`
	// Metadata that can be used to manage the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-fleet.html#cfn-iotfleetwise-fleet-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

