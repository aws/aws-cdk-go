package awscloudfront


// Properties for defining a `CfnKeyGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnKeyGroupProps := &CfnKeyGroupProps{
//   	KeyGroupConfig: &KeyGroupConfigProperty{
//   		Items: []*string{
//   			jsii.String("items"),
//   		},
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Comment: jsii.String("comment"),
//   	},
//   }
//
type CfnKeyGroupProps struct {
	// The key group configuration.
	KeyGroupConfig interface{} `field:"required" json:"keyGroupConfig" yaml:"keyGroupConfig"`
}

