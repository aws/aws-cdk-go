package interfacesawscloudwatch


// A reference to a ResourceMetricsConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceMetricsConfigurationReference := &ResourceMetricsConfigurationReference{
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
type ResourceMetricsConfigurationReference struct {
	// The ResourceArn of the ResourceMetricsConfiguration resource.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

