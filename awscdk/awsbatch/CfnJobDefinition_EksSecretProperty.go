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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekssecret.html
//
type CfnJobDefinition_EksSecretProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekssecret.html#cfn-batch-jobdefinition-ekssecret-secretname
	//
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekssecret.html#cfn-batch-jobdefinition-ekssecret-optional
	//
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

