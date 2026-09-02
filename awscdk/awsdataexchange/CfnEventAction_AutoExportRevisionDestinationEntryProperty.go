package awsdataexchange


// A revision destination is the Amazon S3 bucket folder destination to where the export will be sent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoExportRevisionDestinationEntryProperty := &AutoExportRevisionDestinationEntryProperty{
//   	Bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	KeyPattern: jsii.String("keyPattern"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-autoexportrevisiondestinationentry.html
//
type CfnEventAction_AutoExportRevisionDestinationEntryProperty struct {
	// The Amazon S3 bucket that is the destination for the event action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-autoexportrevisiondestinationentry.html#cfn-dataexchange-eventaction-autoexportrevisiondestinationentry-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// A string representing the pattern for generated names of the individual assets in the revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dataexchange-eventaction-autoexportrevisiondestinationentry.html#cfn-dataexchange-eventaction-autoexportrevisiondestinationentry-keypattern
	//
	KeyPattern *string `field:"optional" json:"keyPattern" yaml:"keyPattern"`
}

