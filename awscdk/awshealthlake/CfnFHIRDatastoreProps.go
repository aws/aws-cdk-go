package awshealthlake

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFHIRDatastore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFHIRDatastoreProps := &CfnFHIRDatastoreProps{
//   	DatastoreTypeVersion: jsii.String("datastoreTypeVersion"),
//
//   	// the properties below are optional
//   	DatastoreName: jsii.String("datastoreName"),
//   	IdentityProviderConfiguration: &IdentityProviderConfigurationProperty{
//   		AuthorizationStrategy: jsii.String("authorizationStrategy"),
//
//   		// the properties below are optional
//   		FineGrainedAuthorizationEnabled: jsii.Boolean(false),
//   		IdpLambdaArn: jsii.String("idpLambdaArn"),
//   		Metadata: jsii.String("metadata"),
//   	},
//   	PreloadDataConfig: &PreloadDataConfigProperty{
//   		PreloadDataType: jsii.String("preloadDataType"),
//   	},
//   	SseConfiguration: &SseConfigurationProperty{
//   		KmsEncryptionConfig: &KmsEncryptionConfigProperty{
//   			CmkType: jsii.String("cmkType"),
//
//   			// the properties below are optional
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html
//
type CfnFHIRDatastoreProps struct {
	// The FHIR version of the Data Store.
	//
	// The only supported version is R4.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-datastoretypeversion
	//
	DatastoreTypeVersion *string `field:"required" json:"datastoreTypeVersion" yaml:"datastoreTypeVersion"`
	// The user generated name for the Data Store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-datastorename
	//
	DatastoreName *string `field:"optional" json:"datastoreName" yaml:"datastoreName"`
	// The identity provider configuration for the datastore.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration
	//
	IdentityProviderConfiguration interface{} `field:"optional" json:"identityProviderConfiguration" yaml:"identityProviderConfiguration"`
	// The preloaded data configuration for the Data Store.
	//
	// Only data preloaded from Synthea is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-preloaddataconfig
	//
	PreloadDataConfig interface{} `field:"optional" json:"preloadDataConfig" yaml:"preloadDataConfig"`
	// The server-side encryption key configuration for a customer provided encryption key specified for creating a Data Store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-sseconfiguration
	//
	SseConfiguration interface{} `field:"optional" json:"sseConfiguration" yaml:"sseConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

