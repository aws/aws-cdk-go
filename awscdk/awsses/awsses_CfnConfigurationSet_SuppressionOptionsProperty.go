package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   suppressionOptionsProperty := &suppressionOptionsProperty{
//   	suppressedReasons: []*string{
//   		jsii.String("suppressedReasons"),
//   	},
//   }
//
type CfnConfigurationSet_SuppressionOptionsProperty struct {
	// `CfnConfigurationSet.SuppressionOptionsProperty.SuppressedReasons`.
	SuppressedReasons *[]*string `field:"optional" json:"suppressedReasons" yaml:"suppressedReasons"`
}

