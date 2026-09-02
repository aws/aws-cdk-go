package awsdataexchange


// Details of the operation to be performed by the job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoExportRevisionToS3RequestDetailsProperty := &AutoExportRevisionToS3RequestDetailsProperty{
//   	RevisionDestination: &AutoExportRevisionDestinationEntryProperty{
//   		Bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		KeyPattern: jsii.String("keyPattern"),
//   	},
//
//   	// the properties below are optional
//   	Encryption: &ExportServerSideEncryptionProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-autoexportrevisiontos3requestdetails.html
//
type CfnEventAction_AutoExportRevisionToS3RequestDetailsProperty struct {
	// A revision destination is the Amazon S3 bucket folder destination to where the export will be sent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-autoexportrevisiontos3requestdetails.html#cfn-dataexchange-eventaction-autoexportrevisiontos3requestdetails-revisiondestination
	//
	RevisionDestination interface{} `field:"required" json:"revisionDestination" yaml:"revisionDestination"`
	// Encryption configuration of the export job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-autoexportrevisiontos3requestdetails.html#cfn-dataexchange-eventaction-autoexportrevisiontos3requestdetails-encryption
	//
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
}

