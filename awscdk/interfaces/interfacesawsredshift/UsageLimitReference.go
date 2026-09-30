package interfacesawsredshift


// A reference to a UsageLimit resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   usageLimitReference := &UsageLimitReference{
//   	UsageLimitArn: jsii.String("usageLimitArn"),
//   }
//
type UsageLimitReference struct {
	// The Arn of the UsageLimit resource.
	UsageLimitArn *string `field:"required" json:"usageLimitArn" yaml:"usageLimitArn"`
}

