package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Validation plugin that uses the CloudFormation validation engine to evaluate templates against built-in rules.
//
// Example:
//   // Rules text, read from disk perhaps
//   var myRules string
//   app := awscdk.NewApp()
//
//   awscdk.Validations_Of(app).AddPlugins(awscdk.NewCloudFormationValidatePlugin(&CloudFormationValidatePluginProps{
//   	GuardRules: []ValidationRuleSource{
//   		&ValidationRuleSource{
//   			Name: jsii.String("My rules"),
//   			Content: myRules,
//   		},
//   	},
//   }))
//
type CloudFormationValidatePlugin interface {
	IPolicyValidationPlugin
	// The name of the plugin that will be displayed in the validation report.
	Name() *string
	// The list of rule IDs that the plugin will evaluate.
	//
	// Used for analytics
	// purposes.
	RuleIds() *[]*string
	// The version of the plugin, following the Semantic Versioning specification (see https://semver.org/). This version is used for analytics purposes, to measure the usage of different plugins and different versions. The value of this property should be kept in sync with the actual version of the software package. If the version is not provided or is not a valid semantic version, it will be reported as `0.0.0`.
	Version() *string
	// The method that will be called by the CDK framework to perform validations.
	//
	// This is where the plugin will evaluate the CloudFormation
	// templates for compliance and report and violations.
	Validate(context IPolicyValidationContext) *PolicyValidationPluginReport
}

// The jsii proxy struct for CloudFormationValidatePlugin
type jsiiProxy_CloudFormationValidatePlugin struct {
	jsiiProxy_IPolicyValidationPlugin
}

func (j *jsiiProxy_CloudFormationValidatePlugin) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationValidatePlugin) RuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ruleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationValidatePlugin) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewCloudFormationValidatePlugin(props *CloudFormationValidatePluginProps) CloudFormationValidatePlugin {
	_init_.Initialize()

	if err := validateNewCloudFormationValidatePluginParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudFormationValidatePlugin{}

	_jsii_.Create(
		"aws-cdk-lib.CloudFormationValidatePlugin",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCloudFormationValidatePlugin_Override(c CloudFormationValidatePlugin, props *CloudFormationValidatePluginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CloudFormationValidatePlugin",
		[]interface{}{props},
		c,
	)
}

func CloudFormationValidatePlugin_PLUGIN_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.CloudFormationValidatePlugin",
		"PLUGIN_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudFormationValidatePlugin) Validate(context IPolicyValidationContext) *PolicyValidationPluginReport {
	if err := c.validateValidateParameters(context); err != nil {
		panic(err)
	}
	var returns *PolicyValidationPluginReport

	_jsii_.Invoke(
		c,
		"validate",
		[]interface{}{context},
		&returns,
	)

	return returns
}

