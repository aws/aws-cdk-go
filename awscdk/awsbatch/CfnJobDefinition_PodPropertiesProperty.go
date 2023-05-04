package awsbatch


// The properties for the pod.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var labels interface{}
//   var limits interface{}
//   var requests interface{}
//
//   podPropertiesProperty := &PodPropertiesProperty{
//   	Containers: []interface{}{
//   		&EksContainerProperty{
//   			Image: jsii.String("image"),
//
//   			// the properties below are optional
//   			Args: []*string{
//   				jsii.String("args"),
//   			},
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			Env: []interface{}{
//   				&EksContainerEnvironmentVariableProperty{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			ImagePullPolicy: jsii.String("imagePullPolicy"),
//   			Name: jsii.String("name"),
//   			Resources: &ResourcesProperty{
//   				Limits: limits,
//   				Requests: requests,
//   			},
//   			SecurityContext: &SecurityContextProperty{
//   				Privileged: jsii.Boolean(false),
//   				ReadOnlyRootFilesystem: jsii.Boolean(false),
//   				RunAsGroup: jsii.Number(123),
//   				RunAsNonRoot: jsii.Boolean(false),
//   				RunAsUser: jsii.Number(123),
//   			},
//   			VolumeMounts: []interface{}{
//   				&EksContainerVolumeMountProperty{
//   					MountPath: jsii.String("mountPath"),
//   					Name: jsii.String("name"),
//   					ReadOnly: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   	DnsPolicy: jsii.String("dnsPolicy"),
//   	HostNetwork: jsii.Boolean(false),
//   	Metadata: &MetadataProperty{
//   		Labels: labels,
//   	},
//   	ServiceAccountName: jsii.String("serviceAccountName"),
//   	Volumes: []interface{}{
//   		&EksVolumeProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			EmptyDir: &EmptyDirProperty{
//   				Medium: jsii.String("medium"),
//   				SizeLimit: jsii.String("sizeLimit"),
//   			},
//   			HostPath: &HostPathProperty{
//   				Path: jsii.String("path"),
//   			},
//   			Secret: &EksSecretProperty{
//   				SecretName: jsii.String("secretName"),
//
//   				// the properties below are optional
//   				Optional: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   }
//
type CfnJobDefinition_PodPropertiesProperty struct {
	// The properties of the container that's used on the Amazon EKS pod.
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// The DNS policy for the pod.
	//
	// The default value is `ClusterFirst` . If the `hostNetwork` parameter is not specified, the default is `ClusterFirstWithHostNet` . `ClusterFirst` indicates that any DNS query that does not match the configured cluster domain suffix is forwarded to the upstream nameserver inherited from the node. If no value was specified for `dnsPolicy` in the [RegisterJobDefinition](https://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html) API operation, then no value will be returned for `dnsPolicy` by either of [DescribeJobDefinitions](https://docs.aws.amazon.com/batch/latest/APIReference/API_DescribeJobDefinitions.html) or [DescribeJobs](https://docs.aws.amazon.com/batch/latest/APIReference/API_DescribeJobs.html) API operations. The pod spec setting will contain either `ClusterFirst` or `ClusterFirstWithHostNet` , depending on the value of the `hostNetwork` parameter. For more information, see [Pod's DNS policy](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy) in the *Kubernetes documentation* .
	//
	// Valid values: `Default` | `ClusterFirst` | `ClusterFirstWithHostNet`.
	DnsPolicy *string `field:"optional" json:"dnsPolicy" yaml:"dnsPolicy"`
	// Indicates if the pod uses the hosts' network IP address.
	//
	// The default value is `true` . Setting this to `false` enables the Kubernetes pod networking model. Most AWS Batch workloads are egress-only and don't require the overhead of IP allocation for each pod for incoming connections. For more information, see [Host namespaces](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/security/pod-security-policy/#host-namespaces) and [Pod networking](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/workloads/pods/#pod-networking) in the *Kubernetes documentation* .
	HostNetwork interface{} `field:"optional" json:"hostNetwork" yaml:"hostNetwork"`
	// `CfnJobDefinition.PodPropertiesProperty.Metadata`.
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// The name of the service account that's used to run the pod.
	//
	// For more information, see [Kubernetes service accounts](https://docs.aws.amazon.com/eks/latest/userguide/service-accounts.html) and [Configure a Kubernetes service account to assume an IAM role](https://docs.aws.amazon.com/eks/latest/userguide/associate-service-account-role.html) in the *Amazon EKS User Guide* and [Configure service accounts for pods](https://docs.aws.amazon.com/https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/) in the *Kubernetes documentation* .
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// Specifies the volumes for a job definition that uses Amazon EKS resources.
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

