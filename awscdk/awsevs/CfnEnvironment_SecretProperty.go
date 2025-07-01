package awsevs


// > Amazon EVS is in public preview release and is subject to change.
//
// A managed secret that contains the credentials for installing vCenter Server, NSX, and SDDC Manager. During environment creation, the Amazon EVS control plane uses AWS Secrets Manager to create, encrypt, validate, and store secrets. If you choose to delete your environment, Amazon EVS also deletes the secrets that are associated with your environment. Amazon EVS does not provide managed rotation of secrets. We recommend that you rotate secrets regularly to ensure that secrets are not long-lived.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretProperty := &SecretProperty{
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-secret.html
//
type CfnEnvironment_SecretProperty struct {
	// The Amazon Resource Name (ARN) of the secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-secret.html#cfn-evs-environment-secret-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

