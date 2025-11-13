package interfacesawsiotfleetwise


// A reference to a Campaign resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   campaignReference := &CampaignReference{
//   	CampaignArn: jsii.String("campaignArn"),
//   	CampaignName: jsii.String("campaignName"),
//   }
//
type CampaignReference struct {
	// The ARN of the Campaign resource.
	CampaignArn *string `field:"required" json:"campaignArn" yaml:"campaignArn"`
	// The Name of the Campaign resource.
	CampaignName *string `field:"required" json:"campaignName" yaml:"campaignName"`
}

