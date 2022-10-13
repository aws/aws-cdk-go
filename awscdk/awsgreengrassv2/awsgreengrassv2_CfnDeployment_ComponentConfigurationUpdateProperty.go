package awsgreengrassv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentConfigurationUpdateProperty := &componentConfigurationUpdateProperty{
//   	merge: jsii.String("merge"),
//   	reset: []*string{
//   		jsii.String("reset"),
//   	},
//   }
//
type CfnDeployment_ComponentConfigurationUpdateProperty struct {
	// `CfnDeployment.ComponentConfigurationUpdateProperty.Merge`.
	Merge *string `field:"optional" json:"merge" yaml:"merge"`
	// `CfnDeployment.ComponentConfigurationUpdateProperty.Reset`.
	Reset *[]*string `field:"optional" json:"reset" yaml:"reset"`
}

