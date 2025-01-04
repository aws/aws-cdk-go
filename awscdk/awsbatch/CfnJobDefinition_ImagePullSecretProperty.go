package awsbatch


// References a Kubernetes secret resource.
//
// This name of the secret must start and end with an alphanumeric character, is required to be lowercase, can include periods (.) and hyphens (-), and can't contain more than 253 characters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imagePullSecretProperty := &ImagePullSecretProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-imagepullsecret.html
//
type CfnJobDefinition_ImagePullSecretProperty struct {
	// Provides a unique identifier for the `ImagePullSecret` .
	//
	// This object is required when `EksPodProperties$imagePullSecrets` is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-imagepullsecret.html#cfn-batch-jobdefinition-imagepullsecret-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

