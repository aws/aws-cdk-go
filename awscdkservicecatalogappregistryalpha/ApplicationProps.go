package awscdkservicecatalogappregistryalpha


// Properties for a Service Catalog AppRegistry Application.
//
// Example:
//   application := appreg.NewApplication(this, jsii.String("MyFirstApplication"), &ApplicationProps{
//   	ApplicationName: jsii.String("MyFirstApplicationName"),
//   	Description: jsii.String("description for my application"),
//   })
//
// Experimental.
type ApplicationProps struct {
	// Enforces a particular physical application name.
	// Experimental.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// Description for application.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

