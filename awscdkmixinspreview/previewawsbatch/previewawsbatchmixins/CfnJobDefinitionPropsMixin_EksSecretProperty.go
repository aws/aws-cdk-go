package previewawsbatchmixins


// Specifies the configuration of a Kubernetes `secret` volume.
//
// For more information, see [secret](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/storage/volumes/#secret) in the *Kubernetes documentation* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eksSecretProperty := &EksSecretProperty{
//   	Optional: jsii.Boolean(false),
//   	SecretName: jsii.String("secretName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekssecret.html
//
type CfnJobDefinitionPropsMixin_EksSecretProperty struct {
	// Specifies whether the secret or the secret's keys must be defined.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekssecret.html#cfn-batch-jobdefinition-ekssecret-optional
	//
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
	// The name of the secret.
	//
	// The name must be allowed as a DNS subdomain name. For more information, see [DNS subdomain names](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-subdomain-names) in the *Kubernetes documentation* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekssecret.html#cfn-batch-jobdefinition-ekssecret-secretname
	//
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
}

