package awsconnectcampaignsv2


// A reference to a Campaign resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   campaignReference := &CampaignReference{
//   	CampaignArn: jsii.String("campaignArn"),
//   }
//
type CampaignReference struct {
	// The Arn of the Campaign resource.
	CampaignArn *string `field:"required" json:"campaignArn" yaml:"campaignArn"`
}

