package awsopensearchserverless


// A reference to a SecurityConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityConfigReference := &SecurityConfigReference{
//   	SecurityConfigId: jsii.String("securityConfigId"),
//   }
//
type SecurityConfigReference struct {
	// The Id of the SecurityConfig resource.
	SecurityConfigId *string `field:"required" json:"securityConfigId" yaml:"securityConfigId"`
}

