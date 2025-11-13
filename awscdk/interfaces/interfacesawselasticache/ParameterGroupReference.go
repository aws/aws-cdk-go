package interfacesawselasticache


// A reference to a ParameterGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterGroupReference := &ParameterGroupReference{
//   	CacheParameterGroupName: jsii.String("cacheParameterGroupName"),
//   }
//
type ParameterGroupReference struct {
	// The CacheParameterGroupName of the ParameterGroup resource.
	CacheParameterGroupName *string `field:"required" json:"cacheParameterGroupName" yaml:"cacheParameterGroupName"`
}

