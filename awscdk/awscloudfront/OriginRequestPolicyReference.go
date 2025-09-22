package awscloudfront


// A reference to a OriginRequestPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originRequestPolicyReference := &OriginRequestPolicyReference{
//   	OriginRequestPolicyId: jsii.String("originRequestPolicyId"),
//   }
//
type OriginRequestPolicyReference struct {
	// The Id of the OriginRequestPolicy resource.
	OriginRequestPolicyId *string `field:"required" json:"originRequestPolicyId" yaml:"originRequestPolicyId"`
}

