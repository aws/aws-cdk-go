package previewawslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// Properties for S3 logs destination configuration.
// Experimental.
type IS3LogsDestinationProps interface {
	// KMS key to use for encrypting logs in the S3 bucket.
	// Default: - No encryption key is configured.
	//
	// Experimental.
	EncryptionKey() interfacesawskms.IKeyRef
}

// The jsii proxy for IS3LogsDestinationProps
type jsiiProxy_IS3LogsDestinationProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IS3LogsDestinationProps) EncryptionKey() interfacesawskms.IKeyRef {
	var returns interfacesawskms.IKeyRef
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

