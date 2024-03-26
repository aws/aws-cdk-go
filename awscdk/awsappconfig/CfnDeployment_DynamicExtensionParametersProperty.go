package awsappconfig


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamicExtensionParametersProperty := &DynamicExtensionParametersProperty{
//   	ExtensionReference: jsii.String("extensionReference"),
//   	ParameterName: jsii.String("parameterName"),
//   	ParameterValue: jsii.String("parameterValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deployment-dynamicextensionparameters.html
//
type CfnDeployment_DynamicExtensionParametersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deployment-dynamicextensionparameters.html#cfn-appconfig-deployment-dynamicextensionparameters-extensionreference
	//
	ExtensionReference *string `field:"optional" json:"extensionReference" yaml:"extensionReference"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deployment-dynamicextensionparameters.html#cfn-appconfig-deployment-dynamicextensionparameters-parametername
	//
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deployment-dynamicextensionparameters.html#cfn-appconfig-deployment-dynamicextensionparameters-parametervalue
	//
	ParameterValue *string `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

