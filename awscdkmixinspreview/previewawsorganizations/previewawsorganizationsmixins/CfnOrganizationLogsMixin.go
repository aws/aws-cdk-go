package previewawsorganizationsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsorganizations/previewawsorganizationsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an AWS organization.
//
// The account whose user is calling the [`CreateOrganization`](https://docs.aws.amazon.com/organizations/latest/APIReference/API_CreateOrganization.html) operation automatically becomes the [management account](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account) of the new organization.
//
// This operation must be called using credentials from the account that is to become the new organization's management account. The principal must also have the [relevant IAM permissions](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_create.html) .
//
// > - If you delete an organization, you can't recover it. If you created any policies inside of the organization, they're also deleted and you can't recover them.
// > - You can delete an organization only after you remove all member accounts from the organization. If you created some of your member accounts using AWS Organizations , you might be blocked from removing those accounts. You can remove a member account only if it has all the information that's required to operate as a standalone AWS account. For more information about how to provide that information and then remove the account, see [Leave an organization from your member account](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_leave-as-member.html) in the *AWS Organizations User Guide* .
// > - If you closed a member account before you remove it from the organization, it enters a 'suspended' state for a period of time and you can't remove the account from the organization until it is finally closed. This can take up to 90 days and can prevent you from deleting the organization until all member accounts are completely closed.
// >
// > For more information, see [Deleting an organization](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_delete.html) in the *AWS Organizations User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnOrganizationLogsMixin := awscdkmixinspreview.Mixins.NewCfnOrganizationLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-organizations-organization.html
//
type CfnOrganizationLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOrganizationLogsMixin
type jsiiProxy_CfnOrganizationLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnOrganizationLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::Organizations::Organization`.
func NewCfnOrganizationLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnOrganizationLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnOrganizationLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOrganizationLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_organizations.mixins.CfnOrganizationLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::Organizations::Organization`.
func NewCfnOrganizationLogsMixin_Override(c CfnOrganizationLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_organizations.mixins.CfnOrganizationLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnOrganizationLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOrganizationLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_organizations.mixins.CfnOrganizationLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOrganizationLogsMixin_ACCESS_CONTROL_LOGS() CfnOrganizationAccessControlLogs {
	_init_.Initialize()
	var returns CfnOrganizationAccessControlLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_organizations.mixins.CfnOrganizationLogsMixin",
		"ACCESS_CONTROL_LOGS",
		&returns,
	)
	return returns
}

func CfnOrganizationLogsMixin_AUTHENTICATION_LOGS() CfnOrganizationAuthenticationLogs {
	_init_.Initialize()
	var returns CfnOrganizationAuthenticationLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_organizations.mixins.CfnOrganizationLogsMixin",
		"AUTHENTICATION_LOGS",
		&returns,
	)
	return returns
}

func CfnOrganizationLogsMixin_WORKMAIL_AVAILABILITY_PROVIDER_LOGS() CfnOrganizationWorkmailAvailabilityProviderLogs {
	_init_.Initialize()
	var returns CfnOrganizationWorkmailAvailabilityProviderLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_organizations.mixins.CfnOrganizationLogsMixin",
		"WORKMAIL_AVAILABILITY_PROVIDER_LOGS",
		&returns,
	)
	return returns
}

func CfnOrganizationLogsMixin_WORKMAIL_MAILBOX_ACCESS_LOGS() CfnOrganizationWorkmailMailboxAccessLogs {
	_init_.Initialize()
	var returns CfnOrganizationWorkmailMailboxAccessLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_organizations.mixins.CfnOrganizationLogsMixin",
		"WORKMAIL_MAILBOX_ACCESS_LOGS",
		&returns,
	)
	return returns
}

func CfnOrganizationLogsMixin_WORKMAIL_PERSONAL_ACCESS_TOKEN_LOGS() CfnOrganizationWorkmailPersonalAccessTokenLogs {
	_init_.Initialize()
	var returns CfnOrganizationWorkmailPersonalAccessTokenLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_organizations.mixins.CfnOrganizationLogsMixin",
		"WORKMAIL_PERSONAL_ACCESS_TOKEN_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOrganizationLogsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOrganizationLogsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

