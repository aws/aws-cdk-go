package awscdkeksv2alpha


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
	// Overwrite existing service account.
	//
	// If this is set, we will use `kubectl apply` instead of `kubectl create`
	// when the service account is created. Otherwise, if there is already a service account
	// in the cluster with the same name, the operation will fail.
	// Default: false.
	//
	// Experimental.
	OverwriteServiceAccount *bool `field:"optional" json:"overwriteServiceAccount" yaml:"overwriteServiceAccount"`
}

