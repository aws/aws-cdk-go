package interfacesawsquicksight


// A reference to a LimitsProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   limitsProfileReference := &LimitsProfileReference{
//   	AccountId: jsii.String("accountId"),
//   	LimitsProfileArn: jsii.String("limitsProfileArn"),
//   	ProfileId: jsii.String("profileId"),
//   }
//
type LimitsProfileReference struct {
	// The AccountId of the LimitsProfile resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The ARN of the LimitsProfile resource.
	LimitsProfileArn *string `field:"required" json:"limitsProfileArn" yaml:"limitsProfileArn"`
	// The ProfileId of the LimitsProfile resource.
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
}

