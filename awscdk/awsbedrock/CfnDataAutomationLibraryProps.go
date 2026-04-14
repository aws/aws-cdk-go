package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataAutomationLibrary`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataAutomationLibraryProps := &CfnDataAutomationLibraryProps{
//   	LibraryName: jsii.String("libraryName"),
//
//   	// the properties below are optional
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//
//   		// the properties below are optional
//   		KmsEncryptionContext: map[string]*string{
//   			"kmsEncryptionContextKey": jsii.String("kmsEncryptionContext"),
//   		},
//   	},
//   	LibraryDescription: jsii.String("libraryDescription"),
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
type CfnDataAutomationLibraryProps struct {
	// Name of the DataAutomationLibrary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationlibrary.html#cfn-bedrock-dataautomationlibrary-libraryname
	//
	LibraryName *string `field:"required" json:"libraryName" yaml:"libraryName"`
	// KMS Encryption Configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationlibrary.html#cfn-bedrock-dataautomationlibrary-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Description of the DataAutomationLibrary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationlibrary.html#cfn-bedrock-dataautomationlibrary-librarydescription
	//
	LibraryDescription *string `field:"optional" json:"libraryDescription" yaml:"libraryDescription"`
	// List of tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-dataautomationlibrary.html#cfn-bedrock-dataautomationlibrary-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

