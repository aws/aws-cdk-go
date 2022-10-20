package awseks


// Options for `ServiceAccount`.
//
// Example:
//   var cluster cluster
//
//   // add service account with annotations and labels
//   serviceAccount := cluster.addServiceAccount(jsii.String("MyServiceAccount"), &serviceAccountOptions{
//   	annotations: map[string]*string{
//   		"eks.amazonaws.com/sts-regional-endpoints": jsii.String("false"),
//   	},
//   	labels: map[string]*string{
//   		"some-label": jsii.String("with-some-value"),
//   	},
//   })
//
type ServiceAccountOptions struct {
	// Additional annotations of the service account.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Additional labels of the service account.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the service account.
	//
	// The name of a ServiceAccount object must be a valid DNS subdomain name.
	// https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The namespace of the service account.
	//
	// All namespace names must be valid RFC 1123 DNS labels.
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/#namespaces-and-dns
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

