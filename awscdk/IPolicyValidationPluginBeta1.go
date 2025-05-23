package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a validation plugin that will be executed during synthesis.
//
// Example:
//   type myPlugin struct {
//   	name
//   }
//
//   func (this *myPlugin) validate(context iPolicyValidationContextBeta1) policyValidationPluginReportBeta1 {
//   	// First read the templates using context.templatePaths...
//
//   	// ...then perform the validation, and then compose and return the report.
//   	// Using hard-coded values here for better clarity:
//   	return &policyValidationPluginReportBeta1{
//   		Success: jsii.Boolean(false),
//   		Violations: []policyViolationBeta1{
//   			&policyViolationBeta1{
//   				RuleName: jsii.String("CKV_AWS_117"),
//   				Description: jsii.String("Ensure that AWS Lambda function is configured inside a VPC"),
//   				Fix: jsii.String("https://docs.bridgecrew.io/docs/ensure-that-aws-lambda-function-is-configured-inside-a-vpc-1"),
//   				ViolatingResources: []policyViolatingResourceBeta1{
//   					&policyViolatingResourceBeta1{
//   						ResourceLogicalId: jsii.String("MyFunction3BAA72D1"),
//   						TemplatePath: jsii.String("/home/johndoe/myapp/cdk.out/MyService.template.json"),
//   						Locations: []*string{
//   							jsii.String("Properties/VpcConfig"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	}
//   }
//
type IPolicyValidationPluginBeta1 interface {
	// The method that will be called by the CDK framework to perform validations.
	//
	// This is where the plugin will evaluate the CloudFormation
	// templates for compliance and report and violations.
	Validate(context IPolicyValidationContextBeta1) *PolicyValidationPluginReportBeta1
	// The name of the plugin that will be displayed in the validation report.
	Name() *string
	// The list of rule IDs that the plugin will evaluate.
	//
	// Used for analytics
	// purposes.
	// Default: - No rule is reported.
	//
	RuleIds() *[]*string
	// The version of the plugin, following the Semantic Versioning specification (see https://semver.org/). This version is used for analytics purposes, to measure the usage of different plugins and different versions. The value of this property should be kept in sync with the actual version of the software package. If the version is not provided or is not a valid semantic version, it will be reported as `0.0.0`.
	Version() *string
}

// The jsii proxy for IPolicyValidationPluginBeta1
type jsiiProxy_IPolicyValidationPluginBeta1 struct {
	_ byte // padding
}

func (i *jsiiProxy_IPolicyValidationPluginBeta1) Validate(context IPolicyValidationContextBeta1) *PolicyValidationPluginReportBeta1 {
	if err := i.validateValidateParameters(context); err != nil {
		panic(err)
	}
	var returns *PolicyValidationPluginReportBeta1

	_jsii_.Invoke(
		i,
		"validate",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IPolicyValidationPluginBeta1) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyValidationPluginBeta1) RuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ruleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyValidationPluginBeta1) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

