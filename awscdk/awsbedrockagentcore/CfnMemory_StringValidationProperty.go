package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringValidationProperty := &StringValidationProperty{
//   	AllowedValues: []*string{
//   		jsii.String("allowedValues"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-stringvalidation.html
//
type CfnMemory_StringValidationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-stringvalidation.html#cfn-bedrockagentcore-memory-stringvalidation-allowedvalues
	//
	AllowedValues *[]*string `field:"required" json:"allowedValues" yaml:"allowedValues"`
}

