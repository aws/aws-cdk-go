package awsefs


// Describes a policy used by EFS lifecycle management and EFS Intelligent-Tiering that specifies when to transition files into and out of the file system's Infrequent Access (IA) storage class.
//
// For more information, see [EFS Intelligent‐Tiering and EFS Lifecycle Management](https://docs.aws.amazon.com/efs/latest/ug/lifecycle-management-efs.html) .
//
// > - Each `LifecyclePolicy` object can have only a single transition. This means that in a request body, `LifecyclePolicies` must be structured as an array of `LifecyclePolicy` objects, one object for each transition, `TransitionToIA` , `TransitionToPrimaryStorageClass` .
// > - See the AWS::EFS::FileSystem examples for the correct `LifecyclePolicy` structure. Do not use the syntax shown on this page.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecyclePolicyProperty := &lifecyclePolicyProperty{
//   	transitionToIa: jsii.String("transitionToIa"),
//   	transitionToPrimaryStorageClass: jsii.String("transitionToPrimaryStorageClass"),
//   }
//
type CfnFileSystem_LifecyclePolicyProperty struct {
	// Describes the period of time that a file is not accessed, after which it transitions to IA storage.
	//
	// Metadata operations such as listing the contents of a directory don't count as file access events.
	TransitionToIa *string `field:"optional" json:"transitionToIa" yaml:"transitionToIa"`
	// Describes when to transition a file from IA storage to primary storage.
	//
	// Metadata operations such as listing the contents of a directory don't count as file access events.
	TransitionToPrimaryStorageClass *string `field:"optional" json:"transitionToPrimaryStorageClass" yaml:"transitionToPrimaryStorageClass"`
}

