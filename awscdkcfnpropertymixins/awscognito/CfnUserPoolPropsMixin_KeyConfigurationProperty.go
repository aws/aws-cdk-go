package awscognito


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   keyConfigurationProperty := &KeyConfigurationProperty{
//   	KeyType: jsii.String("keyType"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-keyconfiguration.html
//
type CfnUserPoolPropsMixin_KeyConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-keyconfiguration.html#cfn-cognito-userpool-keyconfiguration-keytype
	//
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-keyconfiguration.html#cfn-cognito-userpool-keyconfiguration-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

