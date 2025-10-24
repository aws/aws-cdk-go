package awseks


// Options for `ServiceAccount`.
//
// Example:
//   var cluster Cluster
//
//   // add service account with annotations and labels
//   serviceAccount := cluster.addServiceAccount(jsii.String("MyServiceAccount"), &ServiceAccountOptions{
//   	Annotations: map[string]*string{
//   		"eks.amazonaws.com/sts-regional-endpoints": jsii.String("false"),
//   	},
//   	Labels: map[string]*string{
//   		"some-label": jsii.String("with-some-value"),
//   	},
//   })
//
type ServiceAccountOptions struct {
	// Additional annotations of the service account.
	// Default: - no additional annotations.
	//
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// The identity type to use for the service account.
	// Default: IdentityType.IRSA
	//
	IdentityType IdentityType `field:"optional" json:"identityType" yaml:"identityType"`
	// Additional labels of the service account.
	// Default: - no additional labels.
	//
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the service account.
	//
	// The name of a ServiceAccount object must be a valid DNS subdomain name.
	// https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	// Default: - If no name is given, it will use the id of the resource.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The namespace of the service account.
	//
	// All namespace names must be valid RFC 1123 DNS labels.
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/#namespaces-and-dns
	// Default: "default".
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

