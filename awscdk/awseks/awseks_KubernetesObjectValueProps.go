package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for KubernetesObjectValue.
//
// Example:
//   var cluster cluster
//
//   // query the load balancer address
//   myServiceAddress := eks.NewKubernetesObjectValue(this, jsii.String("LoadBalancerAttribute"), &kubernetesObjectValueProps{
//   	cluster: cluster,
//   	objectType: jsii.String("service"),
//   	objectName: jsii.String("my-service"),
//   	jsonPath: jsii.String(".status.loadBalancer.ingress[0].hostname"),
//   })
//
//   // pass the address to a lambda function
//   proxyFunction := lambda.NewFunction(this, jsii.String("ProxyFunction"), &functionProps{
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromInline(jsii.String("my-code")),
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	environment: map[string]*string{
//   		"myServiceAddress": myServiceAddress.value,
//   	},
//   })
//
type KubernetesObjectValueProps struct {
	// The EKS cluster to fetch attributes from.
	//
	// [disable-awslint:ref-via-interface].
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// JSONPath to the specific value.
	// See: https://kubernetes.io/docs/reference/kubectl/jsonpath/
	//
	JsonPath *string `field:"required" json:"jsonPath" yaml:"jsonPath"`
	// The name of the object to query.
	ObjectName *string `field:"required" json:"objectName" yaml:"objectName"`
	// The object type to query.
	//
	// (e.g 'service', 'pod'...)
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
	// The namespace the object belongs to.
	ObjectNamespace *string `field:"optional" json:"objectNamespace" yaml:"objectNamespace"`
	// Timeout for waiting on a value.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

