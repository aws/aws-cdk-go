package awseks


// Properties for KubernetesPatch.
//
// Example:
//   var cluster cluster
//
//   eks.NewKubernetesPatch(this, jsii.String("hello-kub-deployment-label"), &kubernetesPatchProps{
//   	cluster: cluster,
//   	resourceName: jsii.String("deployment/hello-kubernetes"),
//   	applyPatch: map[string]interface{}{
//   		"spec": map[string]*f64{
//   			"replicas": jsii.Number(5),
//   		},
//   	},
//   	restorePatch: map[string]interface{}{
//   		"spec": map[string]*f64{
//   			"replicas": jsii.Number(3),
//   		},
//   	},
//   })
//
// Experimental.
type KubernetesPatchProps struct {
	// The JSON object to pass to `kubectl patch` when the resource is created/updated.
	// Experimental.
	ApplyPatch *map[string]interface{} `field:"required" json:"applyPatch" yaml:"applyPatch"`
	// The cluster to apply the patch to.
	//
	// [disable-awslint:ref-via-interface].
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The full name of the resource to patch (e.g. `deployment/coredns`).
	// Experimental.
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
	// The JSON object to pass to `kubectl patch` when the resource is removed.
	// Experimental.
	RestorePatch *map[string]interface{} `field:"required" json:"restorePatch" yaml:"restorePatch"`
	// The patch type to pass to `kubectl patch`.
	//
	// The default type used by `kubectl patch` is "strategic".
	// Experimental.
	PatchType PatchType `field:"optional" json:"patchType" yaml:"patchType"`
	// The kubernetes API namespace.
	// Experimental.
	ResourceNamespace *string `field:"optional" json:"resourceNamespace" yaml:"resourceNamespace"`
}

