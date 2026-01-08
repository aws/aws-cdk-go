package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticloadbalancingv2"
)

// Properties for the trust store revocation.
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
type TrustStoreRevocationProps struct {
	// The revocation file to add.
	RevocationContents *[]*RevocationContent `field:"required" json:"revocationContents" yaml:"revocationContents"`
	// The trust store.
	TrustStore interfacesawselasticloadbalancingv2.ITrustStoreRef `field:"required" json:"trustStore" yaml:"trustStore"`
}

