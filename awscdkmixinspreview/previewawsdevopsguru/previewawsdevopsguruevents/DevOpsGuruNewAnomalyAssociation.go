package previewawsdevopsguruevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.devopsguru@DevOpsGuruNewAnomalyAssociation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   devOpsGuruNewAnomalyAssociation := awscdkmixinspreview.Events.NewDevOpsGuruNewAnomalyAssociation()
//
// Experimental.
type DevOpsGuruNewAnomalyAssociation interface {
}

// The jsii proxy struct for DevOpsGuruNewAnomalyAssociation
type jsiiProxy_DevOpsGuruNewAnomalyAssociation struct {
	_ byte // padding
}

// Experimental.
func NewDevOpsGuruNewAnomalyAssociation() DevOpsGuruNewAnomalyAssociation {
	_init_.Initialize()

	j := jsiiProxy_DevOpsGuruNewAnomalyAssociation{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devopsguru.events.DevOpsGuruNewAnomalyAssociation",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDevOpsGuruNewAnomalyAssociation_Override(d DevOpsGuruNewAnomalyAssociation) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devopsguru.events.DevOpsGuruNewAnomalyAssociation",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for DevOps Guru New Anomaly Association.
// Experimental.
func DevOpsGuruNewAnomalyAssociation_DevOpsGuruNewAnomalyAssociationPattern(options *DevOpsGuruNewAnomalyAssociation_DevOpsGuruNewAnomalyAssociationProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDevOpsGuruNewAnomalyAssociation_DevOpsGuruNewAnomalyAssociationPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_devopsguru.events.DevOpsGuruNewAnomalyAssociation",
		"devOpsGuruNewAnomalyAssociationPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

