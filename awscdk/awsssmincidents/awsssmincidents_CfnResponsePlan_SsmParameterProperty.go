package awsssmincidents


// The key-value pair parameters to use when running the automation document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ssmParameterProperty := &ssmParameterProperty{
//   	key: jsii.String("key"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnResponsePlan_SsmParameterProperty struct {
	// The key parameter to use when running the automation document.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value parameter to use when running the automation document.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

