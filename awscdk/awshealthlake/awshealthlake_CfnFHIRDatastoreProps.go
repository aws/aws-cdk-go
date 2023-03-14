package awshealthlake

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnFHIRDatastore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFHIRDatastoreProps := &cfnFHIRDatastoreProps{
//   	datastoreTypeVersion: jsii.String("datastoreTypeVersion"),
//
//   	// the properties below are optional
//   	datastoreName: jsii.String("datastoreName"),
//   	preloadDataConfig: &preloadDataConfigProperty{
//   		preloadDataType: jsii.String("preloadDataType"),
//   	},
//   	sseConfiguration: &sseConfigurationProperty{
//   		kmsEncryptionConfig: &kmsEncryptionConfigProperty{
//   			cmkType: jsii.String("cmkType"),
//
//   			// the properties below are optional
//   			kmsKeyId: jsii.String("kmsKeyId"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFHIRDatastoreProps struct {
	// The FHIR version of the Data Store.
	//
	// The only supported version is R4.
	DatastoreTypeVersion *string `field:"required" json:"datastoreTypeVersion" yaml:"datastoreTypeVersion"`
	// The user generated name for the Data Store.
	DatastoreName *string `field:"optional" json:"datastoreName" yaml:"datastoreName"`
	// The preloaded data configuration for the Data Store.
	//
	// Only data preloaded from Synthea is supported.
	PreloadDataConfig interface{} `field:"optional" json:"preloadDataConfig" yaml:"preloadDataConfig"`
	// The server-side encryption key configuration for a customer provided encryption key specified for creating a Data Store.
	SseConfiguration interface{} `field:"optional" json:"sseConfiguration" yaml:"sseConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

