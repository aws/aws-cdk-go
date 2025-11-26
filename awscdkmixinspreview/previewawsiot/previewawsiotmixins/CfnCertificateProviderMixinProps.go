package previewawsiotmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCertificateProviderPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCertificateProviderMixinProps := &CfnCertificateProviderMixinProps{
//   	AccountDefaultForOperations: []*string{
//   		jsii.String("accountDefaultForOperations"),
//   	},
//   	CertificateProviderName: jsii.String("certificateProviderName"),
//   	LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
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
type CfnCertificateProviderMixinProps struct {
	// A list of the operations that the certificate provider will use to generate certificates.
	//
	// Valid value: `CreateCertificateFromCsr` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificateprovider.html#cfn-iot-certificateprovider-accountdefaultforoperations
	//
	AccountDefaultForOperations *[]*string `field:"optional" json:"accountDefaultForOperations" yaml:"accountDefaultForOperations"`
	// The name of the certificate provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificateprovider.html#cfn-iot-certificateprovider-certificateprovidername
	//
	CertificateProviderName *string `field:"optional" json:"certificateProviderName" yaml:"certificateProviderName"`
	// The ARN of the Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificateprovider.html#cfn-iot-certificateprovider-lambdafunctionarn
	//
	LambdaFunctionArn *string `field:"optional" json:"lambdaFunctionArn" yaml:"lambdaFunctionArn"`
	// Metadata that can be used to manage the certificate provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificateprovider.html#cfn-iot-certificateprovider-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

