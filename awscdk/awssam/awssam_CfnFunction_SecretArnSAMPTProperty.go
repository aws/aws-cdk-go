package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretArnSAMPTProperty := &secretArnSAMPTProperty{
//   	secretArn: jsii.String("secretArn"),
//   }
//
type CfnFunction_SecretArnSAMPTProperty struct {
	// `CfnFunction.SecretArnSAMPTProperty.SecretArn`.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
}

