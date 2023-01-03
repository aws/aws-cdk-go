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
	// The region in which you want to reference the service.
	//
	// This is only necessary for *cross-region* references to *opt-in* regions. In those
	// cases, the region name needs to be included to reference the correct service principal.
	// In all other cases, the global service principal name is sufficient.
	//
	// This field behaves differently depending on whether the `@aws-cdk/aws-iam:standardizedServicePrincipals`
	// flag is set or not:
	//
	// - If the flag is set, the input service principal is assumed to be of the form `SERVICE.amazonaws.com`.
	//    That value will always be returned, unless the given region is an opt-in region and the service
	//    principal is rendered in a stack in a different region, in which case `SERVICE.REGION.amazonaws.com`
	//    will be rendered. Under this regime, there is no downside to always specifying the region property:
	//    it will be rendered only if necessary.
	// - If the flag is not set, the service principal will resolve to a single principal
	//    whose name comes from the `@aws-cdk/region-info` package, using the region to override
	//    the stack region. If there is no entry for this service principal in the database,, the input
	//    service name is returned literally. This is legacy behavior and is not recommended.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

