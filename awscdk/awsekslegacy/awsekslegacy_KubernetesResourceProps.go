package awsekslegacy


// Example:
//   var cluster cluster
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
//   // option 1: use a construct
//   // option 1: use a construct
//   eks.NewKubernetesResource(this, jsii.String("hello-kub"), &KubernetesResourceProps{
//   	Cluster: Cluster,
//   	Manifest: []interface{}{
//   		deployment,
//   		service,
//   	},
//   })
//
//   // or, option2: use `addResource`
//   cluster.AddResource(jsii.String("hello-kub"), service, deployment)
//
// Experimental.
type KubernetesResourceProps struct {
	// The EKS cluster to apply this configuration to.
	//
	// [disable-awslint:ref-via-interface].
	// Experimental.
	Cluster Cluster `field:"required" json:"cluster" yaml:"cluster"`
	// The resource manifest.
	//
	// Consists of any number of child resources.
	//
	// When the resource is created/updated, this manifest will be applied to the
	// cluster through `kubectl apply` and when the resource or the stack is
	// deleted, the manifest will be deleted through `kubectl delete`.
	//
	// ```
	// const manifest = {
	//    apiVersion: 'v1',
	//    kind: 'Pod',
	//    metadata: { name: 'mypod' },
	//    spec: {
	//      containers: [ { name: 'hello', image: 'paulbouwer/hello-kubernetes:1.5', ports: [ { containerPort: 8080 } ] } ]
	//    }
	// }
	// ```.
	// Experimental.
	Manifest *[]interface{} `field:"required" json:"manifest" yaml:"manifest"`
}

