package mixinsawslambda


// The [traffic-shifting](https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html) configuration of a Lambda function alias.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aliasRoutingConfigurationProperty := &AliasRoutingConfigurationProperty{
//   	AdditionalVersionWeights: []interface{}{
//   		&VersionWeightProperty{
//   			FunctionVersion: jsii.String("functionVersion"),
//   			FunctionWeight: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-aliasroutingconfiguration.html
//
type CfnAliasPropsMixin_AliasRoutingConfigurationProperty struct {
	// The second version, and the percentage of traffic that's routed to it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-aliasroutingconfiguration.html#cfn-lambda-alias-aliasroutingconfiguration-additionalversionweights
	//
	AdditionalVersionWeights interface{} `field:"optional" json:"additionalVersionWeights" yaml:"additionalVersionWeights"`
}

