package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchPathsProperty := &MatchPathsProperty{
//   	AnyOf: []*string{
//   		jsii.String("anyOf"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-matchpaths.html
//
type CfnGatewayRule_MatchPathsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-matchpaths.html#cfn-bedrockagentcore-gatewayrule-matchpaths-anyof
	//
	AnyOf *[]*string `field:"required" json:"anyOf" yaml:"anyOf"`
}

