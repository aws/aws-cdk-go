package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksSecretProperty := &EksSecretProperty{
//   	SecretName: jsii.String("secretName"),
//
//   	// the properties below are optional
//   	Optional: jsii.Boolean(false),
//   }
//
type CfnJobDefinition_EksSecretProperty struct {
	// `CfnJobDefinition.EksSecretProperty.SecretName`.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// `CfnJobDefinition.EksSecretProperty.Optional`.
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

