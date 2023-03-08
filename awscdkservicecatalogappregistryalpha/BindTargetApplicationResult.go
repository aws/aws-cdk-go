// The CDK Construct Library for AWS::ServiceCatalogAppRegistry
package awscdkservicecatalogappregistryalpha


// Properties for Service Catalog AppRegistry Application Associator to work with.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import servicecatalogappregistry_alpha "github.com/aws/aws-cdk-go/awscdkservicecatalogappregistryalpha"
//
//   var application application
//
//   bindTargetApplicationResult := &BindTargetApplicationResult{
//   	Application: application,
//   }
//
// Experimental.
type BindTargetApplicationResult struct {
	// Created or imported application.
	// Experimental.
	Application IApplication `field:"required" json:"application" yaml:"application"`
}

