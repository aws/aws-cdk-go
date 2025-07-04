package awsqbusiness


// The Amazon QuickSight configuration for an Amazon Q Business application that uses QuickSight as the identity provider.
//
// For more information, see [Creating an Amazon QuickSight integrated application](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/create-quicksight-integrated-application.html) .
//
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
	// The Amazon QuickSight namespace that is used as the identity provider.
	//
	// For more information about QuickSight namespaces, see [Namespace operations](https://docs.aws.amazon.com/quicksight/latest/developerguide/namespace-operations.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-application-quicksightconfiguration.html#cfn-qbusiness-application-quicksightconfiguration-clientnamespace
	//
	ClientNamespace *string `field:"required" json:"clientNamespace" yaml:"clientNamespace"`
}

