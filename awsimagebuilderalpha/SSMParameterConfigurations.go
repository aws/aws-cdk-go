package awsimagebuilderalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

// The SSM parameters to create or update for the distributed AMIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stringParameter StringParameter
//
//   sSMParameterConfigurations := &SSMParameterConfigurations{
//   	Parameter: stringParameter,
//
//   	// the properties below are optional
//   	AmiAccount: jsii.String("amiAccount"),
//   	DataType: awscdk.Aws_ssm.ParameterDataType_TEXT,
//   }
//
// Experimental.
type SSMParameterConfigurations struct {
	// The SSM parameter to create or update.
	// Experimental.
	Parameter awsssm.IStringParameter `field:"required" json:"parameter" yaml:"parameter"`
	// The AWS account ID that will own the SSM parameter in the given region.
	//
	// This must be one of the target accounts
	// that was included in the list of AMI distribution target accounts.
	// Default: The current account is used.
	//
	// Experimental.
	AmiAccount *string `field:"optional" json:"amiAccount" yaml:"amiAccount"`
	// The data type of the SSM parameter.
	// Default: ssm.ParameterDataType.AWS_EC2_IMAGE
	//
	// Experimental.
	DataType awsssm.ParameterDataType `field:"optional" json:"dataType" yaml:"dataType"`
}

