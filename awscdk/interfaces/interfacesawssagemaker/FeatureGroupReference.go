package interfacesawssagemaker


// A reference to a FeatureGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   featureGroupReference := &FeatureGroupReference{
//   	FeatureGroupName: jsii.String("featureGroupName"),
//   }
//
type FeatureGroupReference struct {
	// The FeatureGroupName of the FeatureGroup resource.
	FeatureGroupName *string `field:"required" json:"featureGroupName" yaml:"featureGroupName"`
}

