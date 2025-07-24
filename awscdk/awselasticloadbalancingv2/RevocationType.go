package awselasticloadbalancingv2


// The type of revocation file.
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
type RevocationType string

const (
	// A signed list of revoked certificates.
	RevocationType_CRL RevocationType = "CRL"
)

