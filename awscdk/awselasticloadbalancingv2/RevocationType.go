package awselasticloadbalancingv2


// The type of revocation file.
//
// Example:
//   var trustStore TrustStore
//   var bucket Bucket
//
//
//   elbv2.NewTrustStoreRevocation(this, jsii.String("Revocation"), &TrustStoreRevocationProps{
//   	TrustStore: TrustStore,
//   	RevocationContents: []RevocationContent{
//   		&RevocationContent{
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

