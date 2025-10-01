package awselasticache


// A reference to a SecurityGroupIngress resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityGroupIngressReference := &SecurityGroupIngressReference{
//   	SecurityGroupIngressId: jsii.String("securityGroupIngressId"),
//   }
//
type SecurityGroupIngressReference struct {
	// The Id of the SecurityGroupIngress resource.
	SecurityGroupIngressId *string `field:"required" json:"securityGroupIngressId" yaml:"securityGroupIngressId"`
}

