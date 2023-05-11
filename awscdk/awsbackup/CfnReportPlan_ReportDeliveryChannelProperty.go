package awsbackup


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reportDeliveryChannelProperty := &ReportDeliveryChannelProperty{
//   	S3BucketName: jsii.String("s3BucketName"),
//
//   	// the properties below are optional
//   	Formats: []*string{
//   		jsii.String("formats"),
//   	},
//   	S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   }
//
type CfnReportPlan_ReportDeliveryChannelProperty struct {
	// `CfnReportPlan.ReportDeliveryChannelProperty.S3BucketName`.
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// `CfnReportPlan.ReportDeliveryChannelProperty.Formats`.
	Formats *[]*string `field:"optional" json:"formats" yaml:"formats"`
	// `CfnReportPlan.ReportDeliveryChannelProperty.S3KeyPrefix`.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

