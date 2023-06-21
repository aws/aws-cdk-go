package awssigner

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSigningProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSigningProfileProps := &CfnSigningProfileProps{
//   	PlatformId: jsii.String("platformId"),
//
//   	// the properties below are optional
//   	SignatureValidityPeriod: &SignatureValidityPeriodProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnSigningProfileProps struct {
	// The ID of a platform that is available for use by a signing profile.
	PlatformId *string `field:"required" json:"platformId" yaml:"platformId"`
	// The validity period override for any signature generated using this signing profile.
	//
	// If unspecified, the default is 135 months.
	SignatureValidityPeriod interface{} `field:"optional" json:"signatureValidityPeriod" yaml:"signatureValidityPeriod"`
	// A list of tags associated with the signing profile.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

