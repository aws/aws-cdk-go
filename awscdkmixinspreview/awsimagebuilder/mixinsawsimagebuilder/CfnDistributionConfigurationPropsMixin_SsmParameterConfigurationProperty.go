package mixinsawsimagebuilder


// Configuration for a single Parameter in the AWS Systems Manager (SSM) Parameter Store in a given Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ssmParameterConfigurationProperty := &SsmParameterConfigurationProperty{
//   	AmiAccountId: jsii.String("amiAccountId"),
//   	DataType: jsii.String("dataType"),
//   	ParameterName: jsii.String("parameterName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-ssmparameterconfiguration.html
//
type CfnDistributionConfigurationPropsMixin_SsmParameterConfigurationProperty struct {
	// Specify the account that will own the Parameter in a given Region.
	//
	// During distribution, this account must be specified in distribution settings as a target account for the Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-ssmparameterconfiguration.html#cfn-imagebuilder-distributionconfiguration-ssmparameterconfiguration-amiaccountid
	//
	AmiAccountId *string `field:"optional" json:"amiAccountId" yaml:"amiAccountId"`
	// The data type specifies what type of value the Parameter contains.
	//
	// We recommend that you use data type `aws:ec2:image` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-ssmparameterconfiguration.html#cfn-imagebuilder-distributionconfiguration-ssmparameterconfiguration-datatype
	//
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// This is the name of the Parameter in the target Region or account.
	//
	// The image distribution creates the Parameter if it doesn't already exist. Otherwise, it updates the parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-ssmparameterconfiguration.html#cfn-imagebuilder-distributionconfiguration-ssmparameterconfiguration-parametername
	//
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
}

