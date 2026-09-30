package awsglue


// Configuration that defines how BETWEEN range filter operations are translated into REST API request parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   betweenConfigurationProperty := &BetweenConfigurationProperty{
//   	HighBoundKey: jsii.String("highBoundKey"),
//   	LowBoundKey: jsii.String("lowBoundKey"),
//   	Template: jsii.String("template"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-betweenconfiguration.html
//
type CfnConnectionType_BetweenConfigurationProperty struct {
	// The parameter name used for the upper bound value in a BETWEEN filter operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-betweenconfiguration.html#cfn-glue-connectiontype-betweenconfiguration-highboundkey
	//
	HighBoundKey *string `field:"optional" json:"highBoundKey" yaml:"highBoundKey"`
	// The parameter name used for the lower bound value in a BETWEEN filter operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-betweenconfiguration.html#cfn-glue-connectiontype-betweenconfiguration-lowboundkey
	//
	LowBoundKey *string `field:"optional" json:"lowBoundKey" yaml:"lowBoundKey"`
	// A template string for constructing the BETWEEN filter expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-betweenconfiguration.html#cfn-glue-connectiontype-betweenconfiguration-template
	//
	Template *string `field:"optional" json:"template" yaml:"template"`
}

