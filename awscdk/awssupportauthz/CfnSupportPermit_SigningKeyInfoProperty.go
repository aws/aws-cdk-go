package awssupportauthz


// The signing key used by the permit.
//
// Exactly one key type must be provided.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signingKeyInfoProperty := &SigningKeyInfoProperty{
//   	KmsKey: jsii.String("kmsKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-supportauthz-supportpermit-signingkeyinfo.html
//
type CfnSupportPermit_SigningKeyInfoProperty struct {
	// The ARN of the KMS key used to sign permit grants.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-supportauthz-supportpermit-signingkeyinfo.html#cfn-supportauthz-supportpermit-signingkeyinfo-kmskey
	//
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
}

