// The CDK Construct Library for AWS::ServiceCatalogAppRegistry
package awscdkservicecatalogappregistryalpha


// Properties for Service Catalog AppRegistry Application Associator.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//
//   app := awscdk.NewApp()
//
//   associatedApp := appreg.NewApplicationAssociator(app, jsii.String("AssociatedApplication"), &ApplicationAssociatorProps{
//   	Applications: []targetApplication{
//   		appreg.*targetApplication_CreateApplicationStack(&CreateTargetApplicationOptions{
//   			ApplicationName: jsii.String("MyAssociatedApplication"),
//   			// 'Application containing stacks deployed via CDK.' is the default
//   			ApplicationDescription: jsii.String("Associated Application description"),
//   			StackName: jsii.String("MyAssociatedApplicationStack"),
//   			// AWS Account and Region that are implied by the current CLI configuration is the default
//   			Env: &Environment{
//   				Account: jsii.String("123456789012"),
//   				Region: jsii.String("us-east-1"),
//   			},
//   		}),
//   	},
//   })
//
//   // Associate application to the attribute group.
//   associatedApp.appRegistryApplication.AddAttributeGroup(jsii.String("MyAttributeGroup"), &AttributeGroupAssociationProps{
//   	AttributeGroupName: jsii.String("MyAttributeGroupName"),
//   	Description: jsii.String("Test attribute group"),
//   	Attributes: map[string]interface{}{
//   	},
//   })
//
// Experimental.
type ApplicationAssociatorProps struct {
	// Application associator properties.
	// Experimental.
	Applications *[]TargetApplication `field:"required" json:"applications" yaml:"applications"`
}

