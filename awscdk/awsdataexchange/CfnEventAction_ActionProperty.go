package awsdataexchange


// What occurs after a certain event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &ActionProperty{
//   	ExportRevisionToS3: &AutoExportRevisionToS3RequestDetailsProperty{
//   		RevisionDestination: &AutoExportRevisionDestinationEntryProperty{
//   			Bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			KeyPattern: jsii.String("keyPattern"),
//   		},
//
//   		// the properties below are optional
//   		Encryption: &ExportServerSideEncryptionProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-action.html
//
type CfnEventAction_ActionProperty struct {
	// Details of the operation to be performed by the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-action.html#cfn-dataexchange-eventaction-action-exportrevisiontos3
	//
	ExportRevisionToS3 interface{} `field:"optional" json:"exportRevisionToS3" yaml:"exportRevisionToS3"`
}

