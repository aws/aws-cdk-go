package interfacesawssecretsmanager


// A reference to a ResourcePolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourcePolicyReference := &ResourcePolicyReference{
//   	SecretId: jsii.String("secretId"),
//   }
//
type ResourcePolicyReference struct {
	// The SecretId of the ResourcePolicy resource.
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
}

