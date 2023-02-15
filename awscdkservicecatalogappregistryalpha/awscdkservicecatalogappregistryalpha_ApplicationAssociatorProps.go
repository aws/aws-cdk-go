// The CDK Construct Library for AWS::ServiceCatalogAppRegistry
package awscdkservicecatalogappregistryalpha


// Properties for Service Catalog AppRegistry Application Associator.
//
// Example:
//   app := awscdk.NewApp()
//   associatedApp := appreg.NewApplicationAssociator(app, jsii.String("AssociatedApplication"), &applicationAssociatorProps{
//   	applications: []targetApplication{
//   		appreg.*targetApplication.createApplicationStack(&createTargetApplicationOptions{
//   			applicationName: jsii.String("MyAssociatedApplication"),
//   			// 'Application containing stacks deployed via CDK.' is the default
//   			applicationDescription: jsii.String("Associated Application description"),
//   			stackName: jsii.String("MyAssociatedApplicationStack"),
//   			// AWS Account and Region that are implied by the current CLI configuration is the default
//   			env: &environment{
//   				account: jsii.String("123456789012"),
//   				region: jsii.String("us-east-1"),
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type ApplicationAssociatorProps struct {
	// Application associator properties.
	// Experimental.
	Applications *[]TargetApplication `field:"required" json:"applications" yaml:"applications"`
}

