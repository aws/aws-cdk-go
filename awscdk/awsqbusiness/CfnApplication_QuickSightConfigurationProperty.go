package awsqbusiness


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   quickSightConfigurationProperty := &QuickSightConfigurationProperty{
//   	ClientNamespace: jsii.String("clientNamespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-application-quicksightconfiguration.html
//
type CfnApplication_QuickSightConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-application-quicksightconfiguration.html#cfn-qbusiness-application-quicksightconfiguration-clientnamespace
	//
	ClientNamespace *string `field:"required" json:"clientNamespace" yaml:"clientNamespace"`
}

