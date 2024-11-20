package awsconnectcampaignsv2


// The possible types of channel config parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceProperty := &SourceProperty{
//   	CustomerProfilesSegmentArn: jsii.String("customerProfilesSegmentArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-source.html
//
type CfnCampaign_SourceProperty struct {
	// Arn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-source.html#cfn-connectcampaignsv2-campaign-source-customerprofilessegmentarn
	//
	CustomerProfilesSegmentArn *string `field:"required" json:"customerProfilesSegmentArn" yaml:"customerProfilesSegmentArn"`
}

