package previewawssagemakerevents


// Type definition for VpcConfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vpcConfig := &VpcConfig{
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
// Experimental.
type SageMakerHyperParameterTuningJobStateChange_VpcConfig struct {
	// SecurityGroupIds property.
	//
	// Specify an array of string values to match this event if the actual value of SecurityGroupIds is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Subnets property.
	//
	// Specify an array of string values to match this event if the actual value of Subnets is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

