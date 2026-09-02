package awsdataexchange

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnEventActionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnEventActionMixinProps := &CfnEventActionMixinProps{
//   	Action: &ActionProperty{
//   		ExportRevisionToS3: &AutoExportRevisionToS3RequestDetailsProperty{
//   			Encryption: &ExportServerSideEncryptionProperty{
//   				KmsKeyArn: jsii.String("kmsKeyArn"),
//   				Type: jsii.String("type"),
//   			},
//   			RevisionDestination: &AutoExportRevisionDestinationEntryProperty{
//   				Bucket: jsii.String("bucket"),
//   				KeyPattern: jsii.String("keyPattern"),
//   			},
//   		},
//   	},
//   	Event: &EventProperty{
//   		RevisionPublished: &RevisionPublishedProperty{
//   			DataSetId: jsii.String("dataSetId"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-eventaction.html
//
type CfnEventActionMixinProps struct {
	// What occurs after a certain event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-eventaction.html#cfn-dataexchange-eventaction-action
	//
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// What occurs to start an action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-eventaction.html#cfn-dataexchange-eventaction-event
	//
	Event interface{} `field:"optional" json:"event" yaml:"event"`
	// The tags for the event action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-eventaction.html#cfn-dataexchange-eventaction-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

