// The CDK Construct Library for AWS::ServiceCatalogAppRegistry
package awscdkservicecatalogappregistryalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a Service Catalog AppRegistry AutoApplication.
//
// Example:
//   app := awscdk.NewApp()
//   associatedApp := appreg.NewApplicationAssociator(app, jsii.String("AssociatedApplication"), &applicationAssociatorProps{
//   	applicationName: jsii.String("MyAssociatedApplication"),
//   	description: jsii.String("Testing associated application"),
//   	stackProps: &stackProps{
//   		stackName: jsii.String("MyAssociatedApplicationStack"),
//   		env: &environment{
//   			account: jsii.String("123456789012"),
//   			region: jsii.String("us-east-1"),
//   		},
//   	},
//   })
//
// Experimental.
type ApplicationAssociatorProps struct {
	// Stack properties.
	// Experimental.
	StackProps *awscdk.StackProps `field:"required" json:"stackProps" yaml:"stackProps"`
	// Enforces a particular application arn.
	// Experimental.
	ApplicationArnValue *string `field:"optional" json:"applicationArnValue" yaml:"applicationArnValue"`
	// Enforces a particular physical application name.
	// Experimental.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// Application description.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

