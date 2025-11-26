package previewawsecrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.ecr@ECRImageScan event types for Repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCRImageScan := #error#.NewECRImageScan()
//
// Experimental.
type RepositoryEvents_ECRImageScan interface {
}

// The jsii proxy struct for RepositoryEvents_ECRImageScan
type jsiiProxy_RepositoryEvents_ECRImageScan struct {
	_ byte // padding
}

// Experimental.
func NewRepositoryEvents_ECRImageScan() RepositoryEvents_ECRImageScan {
	_init_.Initialize()

	j := jsiiProxy_RepositoryEvents_ECRImageScan{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRImageScan",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRepositoryEvents_ECRImageScan_Override(r RepositoryEvents_ECRImageScan) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRImageScan",
		nil, // no parameters
		r,
	)
}

