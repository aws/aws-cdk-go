package awssecurityhub


// A reference to a HubV2 resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hubV2Reference := &HubV2Reference{
//   	HubV2Arn: jsii.String("hubV2Arn"),
//   }
//
type HubV2Reference struct {
	// The HubV2Arn of the HubV2 resource.
	HubV2Arn *string `field:"required" json:"hubV2Arn" yaml:"hubV2Arn"`
}

