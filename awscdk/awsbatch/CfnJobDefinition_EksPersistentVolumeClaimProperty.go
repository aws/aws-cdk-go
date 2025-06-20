package awsbatch


// A `persistentVolumeClaim` volume is used to mount a [PersistentVolume](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/storage/persistent-volumes/) into a Pod. PersistentVolumeClaims are a way for users to "claim" durable storage without knowing the details of the particular cloud environment. See the information about [PersistentVolumes](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/storage/persistent-volumes/) in the *Kubernetes documentation* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksPersistentVolumeClaimProperty := &EksPersistentVolumeClaimProperty{
//   	ClaimName: jsii.String("claimName"),
//
//   	// the properties below are optional
//   	ReadOnly: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekspersistentvolumeclaim.html
//
type CfnJobDefinition_EksPersistentVolumeClaimProperty struct {
	// The name of the `persistentVolumeClaim` bounded to a `persistentVolume` .
	//
	// For more information, see [Persistent Volume Claims](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/storage/persistent-volumes/#persistentvolumeclaims) in the *Kubernetes documentation* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekspersistentvolumeclaim.html#cfn-batch-jobdefinition-ekspersistentvolumeclaim-claimname
	//
	ClaimName *string `field:"required" json:"claimName" yaml:"claimName"`
	// An optional boolean value indicating if the mount is read only.
	//
	// Default is false. For more information, see [Read Only Mounts](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/storage/volumes/#read-only-mounts) in the *Kubernetes documentation* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekspersistentvolumeclaim.html#cfn-batch-jobdefinition-ekspersistentvolumeclaim-readonly
	//
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

