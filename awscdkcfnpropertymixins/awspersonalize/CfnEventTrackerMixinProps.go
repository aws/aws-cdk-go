package awspersonalize

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnEventTrackerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnEventTrackerMixinProps := &CfnEventTrackerMixinProps{
//   	DatasetGroupArn: jsii.String("datasetGroupArn"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-eventtracker.html
//
type CfnEventTrackerMixinProps struct {
	// The Amazon Resource Name (ARN) of the dataset group that receives the event data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-eventtracker.html#cfn-personalize-eventtracker-datasetgrouparn
	//
	DatasetGroupArn *string `field:"optional" json:"datasetGroupArn" yaml:"datasetGroupArn"`
	// The name for the event tracker.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-eventtracker.html#cfn-personalize-eventtracker-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of tags to apply to the event tracker.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-eventtracker.html#cfn-personalize-eventtracker-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

