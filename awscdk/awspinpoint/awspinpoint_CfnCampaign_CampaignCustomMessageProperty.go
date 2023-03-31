package awspinpoint


// Specifies the contents of a message that's sent through a custom channel to recipients of a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   campaignCustomMessageProperty := &campaignCustomMessageProperty{
//   	data: jsii.String("data"),
//   }
//
type CfnCampaign_CampaignCustomMessageProperty struct {
	// The raw, JSON-formatted string to use as the payload for the message.
	//
	// The maximum size is 5 KB.
	Data *string `field:"optional" json:"data" yaml:"data"`
}

