package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringListValidationProperty := &StringListValidationProperty{
//   	AllowedValues: []*string{
//   		jsii.String("allowedValues"),
//   	},
//   	MaxItems: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-stringlistvalidation.html
//
type CfnMemory_StringListValidationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-stringlistvalidation.html#cfn-bedrockagentcore-memory-stringlistvalidation-allowedvalues
	//
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-stringlistvalidation.html#cfn-bedrockagentcore-memory-stringlistvalidation-maxitems
	//
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
}

