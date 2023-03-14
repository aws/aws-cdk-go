package awsbatch


// The security context for a job.
//
// For more information, see [Configure a security context for a pod or container](https://docs.aws.amazon.com/https://kubernetes.io/docs/tasks/configure-pod-container/security-context/) in the *Kubernetes documentation* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityContextProperty := &SecurityContextProperty{
//   	Privileged: jsii.Boolean(false),
//   	ReadOnlyRootFilesystem: jsii.Boolean(false),
//   	RunAsGroup: jsii.Number(123),
//   	RunAsNonRoot: jsii.Boolean(false),
//   	RunAsUser: jsii.Number(123),
//   }
//
type CfnJobDefinition_SecurityContextProperty struct {
	// When this parameter is `true` , the container is given elevated permissions on the host container instance.
	//
	// The level of permissions are similar to the `root` user permissions. The default value is `false` . This parameter maps to `privileged` policy in the [Privileged pod security policies](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/security/pod-security-policy/#privileged) in the *Kubernetes documentation* .
	Privileged interface{} `field:"optional" json:"privileged" yaml:"privileged"`
	// When this parameter is `true` , the container is given read-only access to its root file system.
	//
	// The default value is `false` . This parameter maps to `ReadOnlyRootFilesystem` policy in the [Volumes and file systems pod security policies](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/security/pod-security-policy/#volumes-and-file-systems) in the *Kubernetes documentation* .
	ReadOnlyRootFilesystem interface{} `field:"optional" json:"readOnlyRootFilesystem" yaml:"readOnlyRootFilesystem"`
	// When this parameter is specified, the container is run as the specified group ID ( `gid` ).
	//
	// If this parameter isn't specified, the default is the group that's specified in the image metadata. This parameter maps to `RunAsGroup` and `MustRunAs` policy in the [Users and groups pod security policies](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/security/pod-security-policy/#users-and-groups) in the *Kubernetes documentation* .
	RunAsGroup *float64 `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	// When this parameter is specified, the container is run as a user with a `uid` other than 0.
	//
	// If this parameter isn't specified, so such rule is enforced. This parameter maps to `RunAsUser` and `MustRunAsNonRoot` policy in the [Users and groups pod security policies](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/security/pod-security-policy/#users-and-groups) in the *Kubernetes documentation* .
	RunAsNonRoot interface{} `field:"optional" json:"runAsNonRoot" yaml:"runAsNonRoot"`
	// When this parameter is specified, the container is run as the specified user ID ( `uid` ).
	//
	// If this parameter isn't specified, the default is the user that's specified in the image metadata. This parameter maps to `RunAsUser` and `MustRanAs` policy in the [Users and groups pod security policies](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/security/pod-security-policy/#users-and-groups) in the *Kubernetes documentation* .
	RunAsUser *float64 `field:"optional" json:"runAsUser" yaml:"runAsUser"`
}

