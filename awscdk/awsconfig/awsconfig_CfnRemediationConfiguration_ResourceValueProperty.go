package awsconfig


// The dynamic value of the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceValueProperty := &resourceValueProperty{
//   	value: jsii.String("value"),
//   }
//
type CfnRemediationConfiguration_ResourceValueProperty struct {
	// The value is a resource ID.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

