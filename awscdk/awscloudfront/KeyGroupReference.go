package awscloudfront


// A reference to a KeyGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyGroupReference := &KeyGroupReference{
//   	KeyGroupId: jsii.String("keyGroupId"),
//   }
//
type KeyGroupReference struct {
	// The Id of the KeyGroup resource.
	KeyGroupId *string `field:"required" json:"keyGroupId" yaml:"keyGroupId"`
}

