package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dkimAttributesProperty := &dkimAttributesProperty{
//   	signingEnabled: jsii.Boolean(false),
//   }
//
type CfnEmailIdentity_DkimAttributesProperty struct {
	// `CfnEmailIdentity.DkimAttributesProperty.SigningEnabled`.
	SigningEnabled interface{} `field:"optional" json:"signingEnabled" yaml:"signingEnabled"`
}

