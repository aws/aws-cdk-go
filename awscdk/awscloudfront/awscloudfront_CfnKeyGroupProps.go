package awscloudfront


// Properties for defining a `CfnKeyGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnKeyGroupProps := &cfnKeyGroupProps{
//   	keyGroupConfig: &keyGroupConfigProperty{
//   		items: []*string{
//   			jsii.String("items"),
//   		},
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		comment: jsii.String("comment"),
//   	},
//   }
//
type CfnKeyGroupProps struct {
	// The key group configuration.
	KeyGroupConfig interface{} `field:"required" json:"keyGroupConfig" yaml:"keyGroupConfig"`
}

