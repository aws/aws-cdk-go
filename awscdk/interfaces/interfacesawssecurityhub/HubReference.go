package interfacesawssecurityhub


// A reference to a Hub resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hubReference := &HubReference{
//   	HubArn: jsii.String("hubArn"),
//   }
//
type HubReference struct {
	// The ARN of the Hub resource.
	HubArn *string `field:"required" json:"hubArn" yaml:"hubArn"`
}

