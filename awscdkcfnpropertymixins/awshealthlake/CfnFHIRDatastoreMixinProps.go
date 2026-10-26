package awshealthlake

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFHIRDatastorePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnFHIRDatastoreMixinProps := &CfnFHIRDatastoreMixinProps{
//   	AnalyticsConfiguration: &AnalyticsConfigurationProperty{
//   		Status: jsii.String("status"),
//   	},
//   	DatastoreName: jsii.String("datastoreName"),
//   	DatastoreTypeVersion: jsii.String("datastoreTypeVersion"),
//   	IdentityProviderConfiguration: &IdentityProviderConfigurationProperty{
//   		AuthorizationStrategy: jsii.String("authorizationStrategy"),
//   		FineGrainedAuthorizationEnabled: jsii.Boolean(false),
//   		IdpLambdaArn: jsii.String("idpLambdaArn"),
//   		Metadata: jsii.String("metadata"),
//   	},
//   	NlpConfiguration: &NlpConfigurationProperty{
//   		Status: jsii.String("status"),
//   	},
//   	PreloadDataConfig: &PreloadDataConfigProperty{
//   		PreloadDataType: jsii.String("preloadDataType"),
//   	},
//   	ProfileConfiguration: &ProfileConfigurationProperty{
//   		DefaultProfiles: []*string{
//   			jsii.String("defaultProfiles"),
//   		},
//   	},
//   	SseConfiguration: &SseConfigurationProperty{
//   		KmsEncryptionConfig: &KmsEncryptionConfigProperty{
//   			CmkType: jsii.String("cmkType"),
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html
//
type CfnFHIRDatastoreMixinProps struct {
	// The analytics configuration for the Data Store.
	//
	// Omitted when the Data Store can't carry analytics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-analyticsconfiguration
	//
	AnalyticsConfiguration interface{} `field:"optional" json:"analyticsConfiguration" yaml:"analyticsConfiguration"`
	// The data store name (user-generated).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-datastorename
	//
	DatastoreName *string `field:"optional" json:"datastoreName" yaml:"datastoreName"`
	// The FHIR release version supported by the data store.
	//
	// Current support is for version `R4` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-datastoretypeversion
	//
	DatastoreTypeVersion *string `field:"optional" json:"datastoreTypeVersion" yaml:"datastoreTypeVersion"`
	// The identity provider configuration selected when the data store was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration
	//
	IdentityProviderConfiguration interface{} `field:"optional" json:"identityProviderConfiguration" yaml:"identityProviderConfiguration"`
	// The natural language processing (NLP) configuration for the Data Store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-nlpconfiguration
	//
	NlpConfiguration interface{} `field:"optional" json:"nlpConfiguration" yaml:"nlpConfiguration"`
	// The preloaded Synthea data configuration for the data store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-preloaddataconfig
	//
	PreloadDataConfig interface{} `field:"optional" json:"preloadDataConfig" yaml:"preloadDataConfig"`
	// The profile configuration for the Data Store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-healthlake-fhirdatastore.html#cfn-healthlake-fhirdatastore-profileconfiguration
	//
	ProfileConfiguration interface{} `field:"optional" json:"profileConfiguration" yaml:"profileConfiguration"`
	// The server-side encryption key configuration for a customer-provided encryption key specified for creating a data store.
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

