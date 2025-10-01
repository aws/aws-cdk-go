package awspinpoint


// A reference to a Campaign resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   campaignReference := &CampaignReference{
//   	CampaignArn: jsii.String("campaignArn"),
//   	CampaignId: jsii.String("campaignId"),
//   }
//
type CampaignReference struct {
	// The ARN of the Campaign resource.
	CampaignArn *string `field:"required" json:"campaignArn" yaml:"campaignArn"`
	// The CampaignId of the Campaign resource.
	CampaignId *string `field:"required" json:"campaignId" yaml:"campaignId"`
}

