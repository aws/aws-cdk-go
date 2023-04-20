package awstransfer

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProfileProps := &CfnProfileProps{
//   	As2Id: jsii.String("as2Id"),
//   	ProfileType: jsii.String("profileType"),
//
//   	// the properties below are optional
//   	CertificateIds: []*string{
//   		jsii.String("certificateIds"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnProfileProps struct {
	// The `As2Id` is the *AS2-name* , as defined in the [RFC 4130](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc4130) . For inbound transfers, this is the `AS2-From` header for the AS2 messages sent from the partner. For outbound connectors, this is the `AS2-To` header for the AS2 messages sent to the partner using the `StartFileTransfer` API operation. This ID cannot include spaces.
	As2Id *string `field:"required" json:"as2Id" yaml:"as2Id"`
	// Indicates whether to list only `LOCAL` type profiles or only `PARTNER` type profiles.
	//
	// If not supplied in the request, the command lists all types of profiles.
	ProfileType *string `field:"required" json:"profileType" yaml:"profileType"`
	// An array of identifiers for the imported certificates.
	//
	// You use this identifier for working with profiles and partner profiles.
	CertificateIds *[]*string `field:"optional" json:"certificateIds" yaml:"certificateIds"`
	// Key-value pairs that can be used to group and search for profiles.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

