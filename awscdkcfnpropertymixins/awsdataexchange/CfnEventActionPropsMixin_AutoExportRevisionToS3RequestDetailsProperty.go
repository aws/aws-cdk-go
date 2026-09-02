package awsdataexchange


// Details of the operation to be performed by the job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   autoExportRevisionToS3RequestDetailsProperty := &AutoExportRevisionToS3RequestDetailsProperty{
//   	Encryption: &ExportServerSideEncryptionProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		Type: jsii.String("type"),
//   	},
//   	RevisionDestination: &AutoExportRevisionDestinationEntryProperty{
//   		Bucket: jsii.String("bucket"),
//   		KeyPattern: jsii.String("keyPattern"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-autoexportrevisiontos3requestdetails.html
//
type CfnEventActionPropsMixin_AutoExportRevisionToS3RequestDetailsProperty struct {
	// Encryption configuration of the export job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-autoexportrevisiontos3requestdetails.html#cfn-dataexchange-eventaction-autoexportrevisiontos3requestdetails-encryption
	//
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// A revision destination is the Amazon S3 bucket folder destination to where the export will be sent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-autoexportrevisiontos3requestdetails.html#cfn-dataexchange-eventaction-autoexportrevisiontos3requestdetails-revisiondestination
	//
	RevisionDestination interface{} `field:"optional" json:"revisionDestination" yaml:"revisionDestination"`
}

