package previewawsbackupmixins


// Contains information from your report plan about where to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   reportDeliveryChannelProperty := &ReportDeliveryChannelProperty{
//   	Formats: []*string{
//   		jsii.String("formats"),
//   	},
//   	S3BucketName: jsii.String("s3BucketName"),
//   	S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-reportplan-reportdeliverychannel.html
//
type CfnReportPlanPropsMixin_ReportDeliveryChannelProperty struct {
	// The format of your reports: `CSV` , `JSON` , or both.
	//
	// If not specified, the default format is `CSV` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-reportplan-reportdeliverychannel.html#cfn-backup-reportplan-reportdeliverychannel-formats
	//
	Formats *[]*string `field:"optional" json:"formats" yaml:"formats"`
	// The unique name of the S3 bucket that receives your reports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-reportplan-reportdeliverychannel.html#cfn-backup-reportplan-reportdeliverychannel-s3bucketname
	//
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// The prefix for where AWS Backup Audit Manager delivers your reports to Amazon S3.
	//
	// The prefix is this part of the following path: s3://your-bucket-name/ `prefix` /Backup/us-west-2/year/month/day/report-name. If not specified, there is no prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-reportplan-reportdeliverychannel.html#cfn-backup-reportplan-reportdeliverychannel-s3keyprefix
	//
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

