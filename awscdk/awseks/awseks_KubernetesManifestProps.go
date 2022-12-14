package awseks


// Properties for KubernetesManifest.
//
// Example:
//   var cluster cluster
//
//   appLabel := map[string]*string{
//   	"app": jsii.String("hello-kubernetes"),
//   }
//
//   deployment := map[string]interface{}{
//   	"apiVersion": jsii.String("apps/v1"),
//   	"kind": jsii.String("Deployment"),
//   	"metadata": map[string]*string{
//   		"name": jsii.String("hello-kubernetes"),
//   	},
//   	"spec": map[string]interface{}{
//   		"replicas": jsii.Number(3),
//   		"selector": map[string]map[string]*string{
//   			"matchLabels": appLabel,
//   		},
//   		"template": map[string]map[string]map[string]*string{
//   			"metadata": map[string]map[string]*string{
//   				"labels": appLabel,
//   			},
//   			"spec": map[string][]map[string]interface{}{
//   				"containers": []map[string]interface{}{
//   					map[string]interface{}{
//   						"name": jsii.String("hello-kubernetes"),
//   						"image": jsii.String("paulbouwer/hello-kubernetes:1.5"),
//   						"ports": []map[string]*f64{
//   							map[string]*f64{
//   								"containerPort": jsii.Number(8080),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
//   service := map[string]interface{}{
//   	"apiVersion": jsii.String("v1"),
//   	"kind": jsii.String("Service"),
//   	"metadata": map[string]*string{
//   		"name": jsii.String("hello-kubernetes"),
//   	},
//   	"spec": map[string]interface{}{
//   		"type": jsii.String("LoadBalancer"),
//   		"ports": []map[string]*f64{
//   			map[string]*f64{
//   				"port": jsii.Number(80),
//   				"targetPort": jsii.Number(8080),
//   			},
//   		},
//   		"selector": appLabel,
//   	},
//   }
//
//   // option 1: use a construct
//   // option 1: use a construct
//   eks.NewKubernetesManifest(this, jsii.String("hello-kub"), &kubernetesManifestProps{
//   	cluster: cluster,
//   	manifest: []map[string]interface{}{
//   		deployment,
//   		service,
//   	},
//   })
//
//   // or, option2: use `addManifest`
//   cluster.addManifest(jsii.String("hello-kub"), service, deployment)
//
type KubernetesManifestProps struct {
	// Automatically detect `Ingress` resources in the manifest and annotate them so they are picked up by an ALB Ingress Controller.
	IngressAlb *bool `field:"optional" json:"ingressAlb" yaml:"ingressAlb"`
	// Specify the ALB scheme that should be applied to `Ingress` resources.
	//
	// Only applicable if `ingressAlb` is set to `true`.
	IngressAlbScheme AlbScheme `field:"optional" json:"ingressAlbScheme" yaml:"ingressAlbScheme"`
	// When a resource is removed from a Kubernetes manifest, it no longer appears in the manifest, and there is no way to know that this resource needs to be deleted.
	//
	// To address this, `kubectl apply` has a `--prune` option which will
	// query the cluster for all resources with a specific label and will remove
	// all the labeld resources that are not part of the applied manifest. If this
	// option is disabled and a resource is removed, it will become "orphaned" and
	// will not be deleted from the cluster.
	//
	// When this option is enabled (default), the construct will inject a label to
	// all Kubernetes resources included in this manifest which will be used to
	// prune resources when the manifest changes via `kubectl apply --prune`.
	//
	// The label name will be `aws.cdk.eks/prune-<ADDR>` where `<ADDR>` is the
	// 42-char unique address of this construct in the construct tree. Value is
	// empty.
	// See: https://kubernetes.io/docs/tasks/manage-kubernetes-objects/declarative-config/#alternative-kubectl-apply-f-directory-prune-l-your-label
	//
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// A flag to signify if the manifest validation should be skipped.
	SkipValidation *bool `field:"optional" json:"skipValidation" yaml:"skipValidation"`
	// The EKS cluster to apply this manifest to.
	//
	// [disable-awslint:ref-via-interface].
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The manifest to apply.
	//
	// Consists of any number of child resources.
	//
	// When the resources are created/updated, this manifest will be applied to the
	// cluster through `kubectl apply` and when the resources or the stack is
	// deleted, the resources in the manifest will be deleted through `kubectl delete`.
	//
	// Example:
	//   []map[string]interface{}{
	//   	map[string]interface{}{
	//   		"apiVersion": jsii.String("v1"),
	//   		"kind": jsii.String("Pod"),
	//   		"metadata": map[string]*string{
	//   			"name": jsii.String("mypod"),
	//   		},
	//   		"spec": map[string][]map[string]interface{}{
	//   			"containers": []map[string]interface{}{
	//   				map[string]interface{}{
	//   					"name": jsii.String("hello"),
	//   					"image": jsii.String("paulbouwer/hello-kubernetes:1.5"),
	//   					"ports": []map[string]*f64{
	//   						map[string]*f64{
	//   							"containerPort": jsii.Number(8080),
	//   						},
	//   					},
	//   				},
	//   			},
	//   		},
	//   	},
	//   }
	//
	Manifest *[]*map[string]interface{} `field:"required" json:"manifest" yaml:"manifest"`
	// Overwrite any existing resources.
	//
	// If this is set, we will use `kubectl apply` instead of `kubectl create`
	// when the resource is created. Otherwise, if there is already a resource
	// in the cluster with the same name, the operation will fail.
	Overwrite *bool `field:"optional" json:"overwrite" yaml:"overwrite"`
}

