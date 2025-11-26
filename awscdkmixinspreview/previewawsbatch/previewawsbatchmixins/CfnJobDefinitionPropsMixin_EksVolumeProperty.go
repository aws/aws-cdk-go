package previewawsbatchmixins


// Specifies an Amazon EKS volume for a job definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eksVolumeProperty := &EksVolumeProperty{
//   	EmptyDir: &EmptyDirProperty{
//   		Medium: jsii.String("medium"),
//   		SizeLimit: jsii.String("sizeLimit"),
//   	},
//   	HostPath: &HostPathProperty{
//   		Path: jsii.String("path"),
//   	},
//   	Name: jsii.String("name"),
//   	PersistentVolumeClaim: &EksPersistentVolumeClaimProperty{
//   		ClaimName: jsii.String("claimName"),
//   		ReadOnly: jsii.Boolean(false),
//   	},
//   	Secret: &EksSecretProperty{
//   		Optional: jsii.Boolean(false),
//   		SecretName: jsii.String("secretName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksvolume.html
//
type CfnJobDefinitionPropsMixin_EksVolumeProperty struct {
	// Specifies the configuration of a Kubernetes `emptyDir` volume.
	//
	// For more information, see [emptyDir](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/storage/volumes/#emptydir) in the *Kubernetes documentation* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksvolume.html#cfn-batch-jobdefinition-eksvolume-emptydir
	//
	EmptyDir interface{} `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// Specifies the configuration of a Kubernetes `hostPath` volume.
	//
	// For more information, see [hostPath](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/storage/volumes/#hostpath) in the *Kubernetes documentation* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksvolume.html#cfn-batch-jobdefinition-eksvolume-hostpath
	//
	HostPath interface{} `field:"optional" json:"hostPath" yaml:"hostPath"`
	// The name of the volume.
	//
	// The name must be allowed as a DNS subdomain name. For more information, see [DNS subdomain names](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-subdomain-names) in the *Kubernetes documentation* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksvolume.html#cfn-batch-jobdefinition-eksvolume-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the configuration of a Kubernetes `persistentVolumeClaim` bounded to a `persistentVolume` .
	//
	// For more information, see [Persistent Volume Claims](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/storage/persistent-volumes/#persistentvolumeclaims) in the *Kubernetes documentation* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksvolume.html#cfn-batch-jobdefinition-eksvolume-persistentvolumeclaim
	//
	PersistentVolumeClaim interface{} `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	// Specifies the configuration of a Kubernetes `secret` volume.
	//
	// For more information, see [secret](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/storage/volumes/#secret) in the *Kubernetes documentation* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksvolume.html#cfn-batch-jobdefinition-eksvolume-secret
	//
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
}

