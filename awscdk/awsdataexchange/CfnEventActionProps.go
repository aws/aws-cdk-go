package awsdataexchange

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEventAction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventActionProps := &CfnEventActionProps{
//   	Action: &ActionProperty{
//   		ExportRevisionToS3: &AutoExportRevisionToS3RequestDetailsProperty{
//   			RevisionDestination: &AutoExportRevisionDestinationEntryProperty{
//   				Bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				KeyPattern: jsii.String("keyPattern"),
//   			},
//
//   			// the properties below are optional
//   			Encryption: &ExportServerSideEncryptionProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				KmsKeyArn: jsii.String("kmsKeyArn"),
//   			},
//   		},
//   	},
//   	Event: &EventProperty{
//   		RevisionPublished: &RevisionPublishedProperty{
//   			DataSetId: jsii.String("dataSetId"),
//   		},
//   	},
//
//   	// the properties below are optional
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
type CfnEventActionProps struct {
	// What occurs after a certain event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-eventaction.html#cfn-dataexchange-eventaction-action
	//
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// What occurs to start an action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-eventaction.html#cfn-dataexchange-eventaction-event
	//
	Event interface{} `field:"required" json:"event" yaml:"event"`
	// The tags for the event action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-eventaction.html#cfn-dataexchange-eventaction-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

