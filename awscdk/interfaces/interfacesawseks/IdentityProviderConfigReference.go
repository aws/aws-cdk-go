package interfacesawseks


// A reference to a IdentityProviderConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityProviderConfigReference := &IdentityProviderConfigReference{
//   	ClusterName: jsii.String("clusterName"),
//   	IdentityProviderConfigArn: jsii.String("identityProviderConfigArn"),
//   	IdentityProviderConfigName: jsii.String("identityProviderConfigName"),
//   	Type: jsii.String("type"),
//   }
//
type IdentityProviderConfigReference struct {
	// The ClusterName of the IdentityProviderConfig resource.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The ARN of the IdentityProviderConfig resource.
	IdentityProviderConfigArn *string `field:"required" json:"identityProviderConfigArn" yaml:"identityProviderConfigArn"`
	// The IdentityProviderConfigName of the IdentityProviderConfig resource.
	IdentityProviderConfigName *string `field:"required" json:"identityProviderConfigName" yaml:"identityProviderConfigName"`
	// The Type of the IdentityProviderConfig resource.
	Type *string `field:"required" json:"type" yaml:"type"`
}

