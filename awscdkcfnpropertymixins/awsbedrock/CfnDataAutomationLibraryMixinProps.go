package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDataAutomationLibraryPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDataAutomationLibraryMixinProps := &CfnDataAutomationLibraryMixinProps{
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsEncryptionContext: map[string]*string{
//   			"kmsEncryptionContextKey": jsii.String("kmsEncryptionContext"),
//   		},
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	LibraryDescription: jsii.String("libraryDescription"),
//   	LibraryName: jsii.String("libraryName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationlibrary.html
//
type CfnDataAutomationLibraryMixinProps struct {
	// KMS Encryption Configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationlibrary.html#cfn-bedrock-dataautomationlibrary-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Description of the DataAutomationLibrary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationlibrary.html#cfn-bedrock-dataautomationlibrary-librarydescription
	//
	LibraryDescription *string `field:"optional" json:"libraryDescription" yaml:"libraryDescription"`
	// Name of the DataAutomationLibrary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationlibrary.html#cfn-bedrock-dataautomationlibrary-libraryname
	//
	LibraryName *string `field:"optional" json:"libraryName" yaml:"libraryName"`
	// List of tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationlibrary.html#cfn-bedrock-dataautomationlibrary-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

