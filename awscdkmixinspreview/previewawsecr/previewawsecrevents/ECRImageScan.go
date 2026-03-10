package previewawsecrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ecr@ECRImageScan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCRImageScan := awscdkmixinspreview.Events.NewECRImageScan()
//
// Experimental.
type ECRImageScan interface {
}

// The jsii proxy struct for ECRImageScan
type jsiiProxy_ECRImageScan struct {
	_ byte // padding
}

// Experimental.
func NewECRImageScan() ECRImageScan {
	_init_.Initialize()

	j := jsiiProxy_ECRImageScan{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRImageScan",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewECRImageScan_Override(e ECRImageScan) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRImageScan",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for ECR Image Scan.
// Experimental.
func ECRImageScan_EcrImageScanPattern(options *ECRImageScan_ECRImageScanProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateECRImageScan_EcrImageScanPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRImageScan",
		"ecrImageScanPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

