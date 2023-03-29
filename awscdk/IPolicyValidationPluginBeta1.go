// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a validation plugin that will be executed during synthesis.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   myCustomValidatorPlugin struct {
//   name
//   }
//
//   func (this *myCustomValidatorPlugin) isReady() *bool {
//   	// check if the plugin tool is installed
//   	return jsii.Boolean(true)
//   }
//
//   func (this *myCustomValidatorPlugin) validate(context IValidationContext) interface{} {
//   	templatePaths := *context.templatePaths
//   	// perform validation on the template
//   	// if there are any failures report them
//   	return map[string]interface{}{
//   		"pluginName": this.name,
//   		"success": jsii.Boolean(false),
//   		"violations": []map[string]interface{}{
//   			map[string]interface{}{
//   				"ruleName": jsii.String("rule-name"),
//   				"description": jsii.String("description of the rule"),
//   				"violatingResources": []map[string]*string{
//   					map[string]*string{
//   						"resourceName": jsii.String("FailingResource"),
//   						"templatePath": jsii.String("/path/to/stack.template.json"),
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

func (j *jsiiProxy_IPolicyValidationPluginBeta1) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

