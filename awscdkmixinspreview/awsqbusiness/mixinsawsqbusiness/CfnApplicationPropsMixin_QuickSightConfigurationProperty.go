package mixinsawsqbusiness


// The Amazon Quick Suite configuration for an Amazon Q Business application that uses Quick Suite as the identity provider.
//
// For more information, see [Creating an Amazon Quick Suite integrated application](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/create-quicksight-integrated-application.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   quickSightConfigurationProperty := &QuickSightConfigurationProperty{
//   	ClientNamespace: jsii.String("clientNamespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-application-quicksightconfiguration.html
//
type CfnApplicationPropsMixin_QuickSightConfigurationProperty struct {
	// The Amazon Quick Suite namespace that is used as the identity provider.
	//
	// For more information about Quick Suite namespaces, see [Namespace operations](https://docs.aws.amazon.com/quicksight/latest/developerguide/namespace-operations.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-application-quicksightconfiguration.html#cfn-qbusiness-application-quicksightconfiguration-clientnamespace
	//
	ClientNamespace *string `field:"optional" json:"clientNamespace" yaml:"clientNamespace"`
}

