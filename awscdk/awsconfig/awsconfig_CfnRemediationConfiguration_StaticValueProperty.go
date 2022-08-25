package awsconfig


// The static value of the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   staticValueProperty := &staticValueProperty{
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnRemediationConfiguration_StaticValueProperty struct {
	// A list of values.
	//
	// For example, the ARN of the assumed role.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

