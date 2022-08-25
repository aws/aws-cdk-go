package awsiam


// Options for a service principal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var conditions interface{}
//
//   servicePrincipalOpts := &servicePrincipalOpts{
//   	conditions: map[string]interface{}{
//   		"conditionsKey": conditions,
//   	},
//   	region: jsii.String("region"),
//   }
//
type ServicePrincipalOpts struct {
	// Additional conditions to add to the Service Principal.
	Conditions *map[string]interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The region in which the service is operating.
	// Deprecated: You should not need to set this. The stack's region is always correct.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

