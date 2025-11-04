package awscdkeksv2alpha


// Options for `ServiceAccount`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeksv2alpha"
//
//   serviceAccountOptions := &ServiceAccountOptions{
//   	Annotations: map[string]*string{
//   		"annotationsKey": jsii.String("annotations"),
//   	},
//   	IdentityType: eks_v2_alpha.IdentityType_IRSA,
//   	Labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	Name: jsii.String("name"),
//   	Namespace: jsii.String("namespace"),
//   }
//
// Experimental.
type ServiceAccountOptions struct {
	// Additional annotations of the service account.
	// Default: - no additional annotations.
	//
	// Experimental.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// The identity type to use for the service account.
	// Default: IdentityType.IRSA
	//
	// Experimental.
	IdentityType IdentityType `field:"optional" json:"identityType" yaml:"identityType"`
	// Additional labels of the service account.
	// Default: - no additional labels.
	//
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the service account.
	//
	// The name of a ServiceAccount object must be a valid DNS subdomain name.
	// https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	// Default: - If no name is given, it will use the id of the resource.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The namespace of the service account.
	//
	// All namespace names must be valid RFC 1123 DNS labels.
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/#namespaces-and-dns
	// Default: "default".
	//
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

