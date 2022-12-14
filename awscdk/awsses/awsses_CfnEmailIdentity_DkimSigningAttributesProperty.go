package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dkimSigningAttributesProperty := &dkimSigningAttributesProperty{
//   	domainSigningPrivateKey: jsii.String("domainSigningPrivateKey"),
//   	domainSigningSelector: jsii.String("domainSigningSelector"),
//   	nextSigningKeyLength: jsii.String("nextSigningKeyLength"),
//   }
//
type CfnEmailIdentity_DkimSigningAttributesProperty struct {
	// `CfnEmailIdentity.DkimSigningAttributesProperty.DomainSigningPrivateKey`.
	DomainSigningPrivateKey *string `field:"optional" json:"domainSigningPrivateKey" yaml:"domainSigningPrivateKey"`
	// `CfnEmailIdentity.DkimSigningAttributesProperty.DomainSigningSelector`.
	DomainSigningSelector *string `field:"optional" json:"domainSigningSelector" yaml:"domainSigningSelector"`
	// `CfnEmailIdentity.DkimSigningAttributesProperty.NextSigningKeyLength`.
	NextSigningKeyLength *string `field:"optional" json:"nextSigningKeyLength" yaml:"nextSigningKeyLength"`
}

