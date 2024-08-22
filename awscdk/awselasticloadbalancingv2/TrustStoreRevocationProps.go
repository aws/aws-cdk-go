package awselasticloadbalancingv2


// Properties for the trust store revocation.
//
// Example:
//   var trustStore trustStore
//   var bucket bucket
//
//
//   elbv2.NewTrustStoreRevocation(this, jsii.String("Revocation"), &TrustStoreRevocationProps{
//   	TrustStore: TrustStore,
//   	RevocationContents: []revocationContent{
//   		&revocationContent{
//   			RevocationType: elbv2.RevocationType_CRL,
//   			Bucket: *Bucket,
//   			Key: jsii.String("crl.pem"),
//   		},
//   	},
//   })
//
type TrustStoreRevocationProps struct {
	// The revocation file to add.
	RevocationContents *[]*RevocationContent `field:"required" json:"revocationContents" yaml:"revocationContents"`
	// The trust store.
	TrustStore ITrustStore `field:"required" json:"trustStore" yaml:"trustStore"`
}

