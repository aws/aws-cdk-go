package interfacesawsbatch


// A reference to a QuotaShare resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   quotaShareReference := &QuotaShareReference{
//   	QuotaShareArn: jsii.String("quotaShareArn"),
//   }
//
type QuotaShareReference struct {
	// The QuotaShareArn of the QuotaShare resource.
	QuotaShareArn *string `field:"required" json:"quotaShareArn" yaml:"quotaShareArn"`
}

