package awsconnectcampaignsv2


// Time range in 24 hour format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeRangeProperty := &TimeRangeProperty{
//   	EndTime: jsii.String("endTime"),
//   	StartTime: jsii.String("startTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-timerange.html
//
type CfnCampaign_TimeRangeProperty struct {
	// Time in ISO 8601 format, e.g. T23:11.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-timerange.html#cfn-connectcampaignsv2-campaign-timerange-endtime
	//
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// Time in ISO 8601 format, e.g. T23:11.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-timerange.html#cfn-connectcampaignsv2-campaign-timerange-starttime
	//
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

