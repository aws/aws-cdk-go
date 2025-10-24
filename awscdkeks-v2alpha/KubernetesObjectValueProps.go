package awscdkeks-v2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for KubernetesObjectValue.
//
// Example:
//   var cluster Cluster
//
//   // query the load balancer address
//   myServiceAddress := eks.NewKubernetesObjectValue(this, jsii.String("LoadBalancerAttribute"), &KubernetesObjectValueProps{
//   	Cluster: cluster,
//   	ObjectType: jsii.String("service"),
//   	ObjectName: jsii.String("my-service"),
//   	JsonPath: jsii.String(".status.loadBalancer.ingress[0].hostname"),
//   })
//
//   // pass the address to a lambda function
//   proxyFunction := lambda.NewFunction(this, jsii.String("ProxyFunction"), &FunctionProps{
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String("my-code")),
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Environment: map[string]*string{
//   		"myServiceAddress": myServiceAddress.value,
//   	},
//   })
//
// Experimental.
type KubernetesObjectValueProps struct {
	// The EKS cluster to fetch attributes from.
	//
	// [disable-awslint:ref-via-interface].
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// JSONPath to the specific value.
	// See: https://kubernetes.io/docs/reference/kubectl/jsonpath/
	//
	// Experimental.
	JsonPath *string `field:"required" json:"jsonPath" yaml:"jsonPath"`
	// The name of the object to query.
	// Experimental.
	ObjectName *string `field:"required" json:"objectName" yaml:"objectName"`
	// The object type to query.
	//
	// (e.g 'service', 'pod'...)
	// Experimental.
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
	// The namespace the object belongs to.
	// Default: 'default'.
	//
	// Experimental.
	ObjectNamespace *string `field:"optional" json:"objectNamespace" yaml:"objectNamespace"`
	// Timeout for waiting on a value.
	// Default: Duration.minutes(5)
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

