package awsbatch


// The properties for the pod.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var limits interface{}
//   var requests interface{}
//
//   podPropertiesProperty := &podPropertiesProperty{
//   	containers: []interface{}{
//   		&eksContainerProperty{
//   			image: jsii.String("image"),
//
//   			// the properties below are optional
//   			args: []*string{
//   				jsii.String("args"),
//   			},
//   			command: []*string{
//   				jsii.String("command"),
//   			},
//   			env: []interface{}{
//   				&eksContainerEnvironmentVariableProperty{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					value: jsii.String("value"),
//   				},
//   			},
//   			imagePullPolicy: jsii.String("imagePullPolicy"),
//   			name: jsii.String("name"),
//   			resources: &resourcesProperty{
//   				limits: limits,
//   				requests: requests,
//   			},
//   			securityContext: &securityContextProperty{
//   				privileged: jsii.Boolean(false),
//   				readOnlyRootFilesystem: jsii.Boolean(false),
//   				runAsGroup: jsii.Number(123),
//   				runAsNonRoot: jsii.Boolean(false),
//   				runAsUser: jsii.Number(123),
//   			},
//   			volumeMounts: []interface{}{
//   				&eksContainerVolumeMountProperty{
//   					mountPath: jsii.String("mountPath"),
//   					name: jsii.String("name"),
//   					readOnly: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   	dnsPolicy: jsii.String("dnsPolicy"),
//   	hostNetwork: jsii.Boolean(false),
//   	serviceAccountName: jsii.String("serviceAccountName"),
//   	volumes: []interface{}{
//   		&eksVolumeProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			emptyDir: &emptyDirProperty{
//   				medium: jsii.String("medium"),
//   				sizeLimit: jsii.String("sizeLimit"),
//   			},
//   			hostPath: &hostPathProperty{
//   				path: jsii.String("path"),
//   			},
//   			secret: &secretProperty{
//   				name: jsii.String("name"),
//   				valueFrom: jsii.String("valueFrom"),
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
	// The name of the service account that's used to run the pod.
	//
	// For more information, see [Kubernetes service accounts](https://docs.aws.amazon.com/eks/latest/userguide/service-accounts.html) and [Configure a Kubernetes service account to assume an IAM role](https://docs.aws.amazon.com/eks/latest/userguide/associate-service-account-role.html) in the *Amazon EKS User Guide* and [Configure service accounts for pods](https://docs.aws.amazon.com/https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/) in the *Kubernetes documentation* .
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// Specifies the volumes for a job definition that uses Amazon EKS resources.
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

