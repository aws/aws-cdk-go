package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCertificateProvider`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateProviderProps := &CfnCertificateProviderProps{
//   	AccountDefaultForOperations: []*string{
//   		jsii.String("accountDefaultForOperations"),
//   	},
//   	LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//
//   	// the properties below are optional
//   	CertificateProviderName: jsii.String("certificateProviderName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificateprovider.html
//
type CfnCertificateProviderProps struct {
	// A list of the operations that the certificate provider will use to generate certificates.
	//
	// Valid value: `CreateCertificateFromCsr` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificateprovider.html#cfn-iot-certificateprovider-accountdefaultforoperations
	//
	AccountDefaultForOperations *[]*string `field:"required" json:"accountDefaultForOperations" yaml:"accountDefaultForOperations"`
	// The ARN of the Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificateprovider.html#cfn-iot-certificateprovider-lambdafunctionarn
	//
	LambdaFunctionArn *string `field:"required" json:"lambdaFunctionArn" yaml:"lambdaFunctionArn"`
	// The name of the certificate provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificateprovider.html#cfn-iot-certificateprovider-certificateprovidername
	//
	CertificateProviderName *string `field:"optional" json:"certificateProviderName" yaml:"certificateProviderName"`
	// Metadata that can be used to manage the certificate provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificateprovider.html#cfn-iot-certificateprovider-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

