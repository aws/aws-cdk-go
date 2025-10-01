package awsshield


// A reference to a DRTAccess resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dRTAccessReference := &DRTAccessReference{
//   	AccountId: jsii.String("accountId"),
//   }
//
type DRTAccessReference struct {
	// The AccountId of the DRTAccess resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

