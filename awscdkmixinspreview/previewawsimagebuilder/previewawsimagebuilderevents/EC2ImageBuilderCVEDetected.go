package previewawsimagebuilderevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.imagebuilder@EC2ImageBuilderCVEDetected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2ImageBuilderCVEDetected := awscdkmixinspreview.Events.NewEC2ImageBuilderCVEDetected()
//
// Experimental.
type EC2ImageBuilderCVEDetected interface {
}

// The jsii proxy struct for EC2ImageBuilderCVEDetected
type jsiiProxy_EC2ImageBuilderCVEDetected struct {
	_ byte // padding
}

// Experimental.
func NewEC2ImageBuilderCVEDetected() EC2ImageBuilderCVEDetected {
	_init_.Initialize()

	j := jsiiProxy_EC2ImageBuilderCVEDetected{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_imagebuilder.events.EC2ImageBuilderCVEDetected",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEC2ImageBuilderCVEDetected_Override(e EC2ImageBuilderCVEDetected) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_imagebuilder.events.EC2ImageBuilderCVEDetected",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EC2 Image Builder CVE Detected.
// Experimental.
func EC2ImageBuilderCVEDetected_Ec2ImageBuilderCVEDetectedPattern(options *EC2ImageBuilderCVEDetected_EC2ImageBuilderCVEDetectedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEC2ImageBuilderCVEDetected_Ec2ImageBuilderCVEDetectedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_imagebuilder.events.EC2ImageBuilderCVEDetected",
		"ec2ImageBuilderCVEDetectedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

