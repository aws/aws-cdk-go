package awsec2


// A reference to a SecurityGroupEgress resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityGroupEgressReference := &SecurityGroupEgressReference{
//   	SecurityGroupEgressId: jsii.String("securityGroupEgressId"),
//   }
//
type SecurityGroupEgressReference struct {
	// The Id of the SecurityGroupEgress resource.
	SecurityGroupEgressId *string `field:"required" json:"securityGroupEgressId" yaml:"securityGroupEgressId"`
}

