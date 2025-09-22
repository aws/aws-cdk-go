package awsdatazone


// A reference to a GroupProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupProfileReference := &GroupProfileReference{
//   	DomainId: jsii.String("domainId"),
//   	GroupProfileId: jsii.String("groupProfileId"),
//   }
//
type GroupProfileReference struct {
	// The DomainId of the GroupProfile resource.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The Id of the GroupProfile resource.
	GroupProfileId *string `field:"required" json:"groupProfileId" yaml:"groupProfileId"`
}

