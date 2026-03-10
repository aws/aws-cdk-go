package previewawssecurityhubevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_securityhub.events.SecurityHubFindingsCustomAction",
		reflect.TypeOf((*SecurityHubFindingsCustomAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_SecurityHubFindingsCustomAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securityhub.events.SecurityHubFindingsCustomAction.SecurityHubFindingsCustomActionProps",
		reflect.TypeOf((*SecurityHubFindingsCustomAction_SecurityHubFindingsCustomActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_securityhub.events.SecurityHubFindingsImported",
		reflect.TypeOf((*SecurityHubFindingsImported)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_SecurityHubFindingsImported{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securityhub.events.SecurityHubFindingsImported.SecurityHubFindingsImportedProps",
		reflect.TypeOf((*SecurityHubFindingsImported_SecurityHubFindingsImportedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_securityhub.events.SecurityHubInsightResults",
		reflect.TypeOf((*SecurityHubInsightResults)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_SecurityHubInsightResults{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securityhub.events.SecurityHubInsightResults.SecurityHubInsightResultsProps",
		reflect.TypeOf((*SecurityHubInsightResults_SecurityHubInsightResultsProps)(nil)).Elem(),
	)
}
