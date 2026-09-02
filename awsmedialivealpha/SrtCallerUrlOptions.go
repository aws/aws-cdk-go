package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Options for a URL-based SRT caller destination (`SrtDestination.callerUrl`).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secret Secret
//
//   srtCallerUrlOptions := &SrtCallerUrlOptions{
//   	EncryptionPassphraseSecret: secret,
//
//   	// the properties below are optional
//   	StreamId: jsii.String("streamId"),
//   }
//
// Experimental.
type SrtCallerUrlOptions struct {
	// The Secrets Manager secret containing the encryption passphrase.
	//
	// [disable-awslint:prefer-ref-interface].
	// Experimental.
	EncryptionPassphraseSecret awssecretsmanager.ISecret `field:"required" json:"encryptionPassphraseSecret" yaml:"encryptionPassphraseSecret"`
	// The stream ID for the SRT connection.
	// Default: - no stream ID.
	//
	// Experimental.
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

