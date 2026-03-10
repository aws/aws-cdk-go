package previewawssecurityhubevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.securityhub@SecurityHubFindingsImported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   securityHubFindingsImported := awscdkmixinspreview.Events.NewSecurityHubFindingsImported()
//
// Experimental.
type SecurityHubFindingsImported interface {
}

// The jsii proxy struct for SecurityHubFindingsImported
type jsiiProxy_SecurityHubFindingsImported struct {
	_ byte // padding
}

// Experimental.
func NewSecurityHubFindingsImported() SecurityHubFindingsImported {
	_init_.Initialize()

	j := jsiiProxy_SecurityHubFindingsImported{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.events.SecurityHubFindingsImported",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSecurityHubFindingsImported_Override(s SecurityHubFindingsImported) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.events.SecurityHubFindingsImported",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for Security Hub Findings - Imported.
// Experimental.
func SecurityHubFindingsImported_SecurityHubFindingsImportedPattern(options *SecurityHubFindingsImported_SecurityHubFindingsImportedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSecurityHubFindingsImported_SecurityHubFindingsImportedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_securityhub.events.SecurityHubFindingsImported",
		"securityHubFindingsImportedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

