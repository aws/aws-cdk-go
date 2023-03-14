package awsroute53


// Properties for defining a `CfnKeySigningKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnKeySigningKeyProps := &cfnKeySigningKeyProps{
//   	hostedZoneId: jsii.String("hostedZoneId"),
//   	keyManagementServiceArn: jsii.String("keyManagementServiceArn"),
//   	name: jsii.String("name"),
//   	status: jsii.String("status"),
//   }
//
type CfnKeySigningKeyProps struct {
	// The unique string (ID) that is used to identify a hosted zone.
	//
	// For example: `Z00001111A1ABCaaABC11` .
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
	// The Amazon resource name (ARN) for a customer managed customer master key (CMK) in AWS Key Management Service ( AWS KMS ).
	//
	// The `KeyManagementServiceArn` must be unique for each key-signing key (KSK) in a single hosted zone. For example: `arn:aws:kms:us-east-1:111122223333:key/111a2222-a11b-1ab1-2ab2-1ab21a2b3a111` .
	KeyManagementServiceArn *string `field:"required" json:"keyManagementServiceArn" yaml:"keyManagementServiceArn"`
	// A string used to identify a key-signing key (KSK).
	//
	// `Name` can include numbers, letters, and underscores (_). `Name` must be unique for each key-signing key in the same hosted zone.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A string that represents the current key-signing key (KSK) status.
	//
	// Status can have one of the following values:
	//
	// - **ACTIVE** - The KSK is being used for signing.
	// - **INACTIVE** - The KSK is not being used for signing.
	// - **DELETING** - The KSK is in the process of being deleted.
	// - **ACTION_NEEDED** - There is a problem with the KSK that requires you to take action to resolve. For example, the customer managed key might have been deleted, or the permissions for the customer managed key might have been changed.
	// - **INTERNAL_FAILURE** - There was an error during a request. Before you can continue to work with DNSSEC signing, including actions that involve this KSK, you must correct the problem. For example, you may need to activate or deactivate the KSK.
	Status *string `field:"required" json:"status" yaml:"status"`
}

