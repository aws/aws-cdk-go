package awsdataexchange


// What occurs after a certain event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   actionProperty := &ActionProperty{
//   	ExportRevisionToS3: &AutoExportRevisionToS3RequestDetailsProperty{
//   		Encryption: &ExportServerSideEncryptionProperty{
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   			Type: jsii.String("type"),
//   		},
//   		RevisionDestination: &AutoExportRevisionDestinationEntryProperty{
//   			Bucket: jsii.String("bucket"),
//   			KeyPattern: jsii.String("keyPattern"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-action.html
//
type CfnEventActionPropsMixin_ActionProperty struct {
	// Details of the operation to be performed by the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-action.html#cfn-dataexchange-eventaction-action-exportrevisiontos3
	//
	ExportRevisionToS3 interface{} `field:"optional" json:"exportRevisionToS3" yaml:"exportRevisionToS3"`
}

