package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var limits interface{}
//   var requests interface{}
//
//   resourcesProperty := &resourcesProperty{
//   	limits: limits,
//   	requests: requests,
//   }
//
type CfnJobDefinition_ResourcesProperty struct {
	// `CfnJobDefinition.ResourcesProperty.Limits`.
	Limits interface{} `field:"optional" json:"limits" yaml:"limits"`
	// `CfnJobDefinition.ResourcesProperty.Requests`.
	Requests interface{} `field:"optional" json:"requests" yaml:"requests"`
}

