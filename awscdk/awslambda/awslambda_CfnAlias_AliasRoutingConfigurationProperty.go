package awslambda


// The [traffic-shifting](https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html) configuration of a Lambda function alias.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aliasRoutingConfigurationProperty := &aliasRoutingConfigurationProperty{
//   	additionalVersionWeights: []interface{}{
//   		&versionWeightProperty{
//   			functionVersion: jsii.String("functionVersion"),
//   			functionWeight: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnAlias_AliasRoutingConfigurationProperty struct {
	// The second version, and the percentage of traffic that's routed to it.
	AdditionalVersionWeights interface{} `field:"required" json:"additionalVersionWeights" yaml:"additionalVersionWeights"`
}

