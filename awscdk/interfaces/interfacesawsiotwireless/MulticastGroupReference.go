package interfacesawsiotwireless


// A reference to a MulticastGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multicastGroupReference := &MulticastGroupReference{
//   	MulticastGroupArn: jsii.String("multicastGroupArn"),
//   	MulticastGroupId: jsii.String("multicastGroupId"),
//   }
//
type MulticastGroupReference struct {
	// The ARN of the MulticastGroup resource.
	MulticastGroupArn *string `field:"required" json:"multicastGroupArn" yaml:"multicastGroupArn"`
	// The Id of the MulticastGroup resource.
	MulticastGroupId *string `field:"required" json:"multicastGroupId" yaml:"multicastGroupId"`
}

