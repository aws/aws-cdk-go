package interfacesawssecurityagent


// A reference to a TargetDomain resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetDomainReference := &TargetDomainReference{
//   	TargetDomainId: jsii.String("targetDomainId"),
//   }
//
type TargetDomainReference struct {
	// The TargetDomainId of the TargetDomain resource.
	TargetDomainId *string `field:"required" json:"targetDomainId" yaml:"targetDomainId"`
}

