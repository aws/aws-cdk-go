package awspcs


// Additional slurmdbd configuration settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slurmdbdCustomSettingProperty := &SlurmdbdCustomSettingProperty{
//   	ParameterName: jsii.String("parameterName"),
//   	ParameterValue: jsii.String("parameterValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-slurmdbdcustomsetting.html
//
type CfnCluster_SlurmdbdCustomSettingProperty struct {
	// The slurmdbd.conf parameter name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-slurmdbdcustomsetting.html#cfn-pcs-cluster-slurmdbdcustomsetting-parametername
	//
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// The value for the slurmdbd.conf parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-slurmdbdcustomsetting.html#cfn-pcs-cluster-slurmdbdcustomsetting-parametervalue
	//
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

