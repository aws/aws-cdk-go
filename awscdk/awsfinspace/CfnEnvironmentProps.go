package awsfinspace

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEnvironment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnvironmentProps := &CfnEnvironmentProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DataBundles: []*string{
//   		jsii.String("dataBundles"),
//   	},
//   	Description: jsii.String("description"),
//   	FederationMode: jsii.String("federationMode"),
//   	FederationParameters: &FederationParametersProperty{
//   		ApplicationCallBackUrl: jsii.String("applicationCallBackUrl"),
//   		AttributeMap: []interface{}{
//   			&AttributeMapItemsProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		FederationProviderName: jsii.String("federationProviderName"),
//   		FederationUrn: jsii.String("federationUrn"),
//   		SamlMetadataDocument: jsii.String("samlMetadataDocument"),
//   		SamlMetadataUrl: jsii.String("samlMetadataUrl"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	SuperuserParameters: &SuperuserParametersProperty{
//   		EmailAddress: jsii.String("emailAddress"),
//   		FirstName: jsii.String("firstName"),
//   		LastName: jsii.String("lastName"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-finspace-environment.html
//
type CfnEnvironmentProps struct {
	// The name of the FinSpace environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-finspace-environment.html#cfn-finspace-environment-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// ARNs of FinSpace Data Bundles to install.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-finspace-environment.html#cfn-finspace-environment-databundles
	//
	// Deprecated: this property has been deprecated.
	DataBundles *[]*string `field:"optional" json:"dataBundles" yaml:"dataBundles"`
	// The description of the FinSpace environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-finspace-environment.html#cfn-finspace-environment-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The authentication mode for the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-finspace-environment.html#cfn-finspace-environment-federationmode
	//
	FederationMode *string `field:"optional" json:"federationMode" yaml:"federationMode"`
	// Configuration information when authentication mode is FEDERATED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-finspace-environment.html#cfn-finspace-environment-federationparameters
	//
	FederationParameters interface{} `field:"optional" json:"federationParameters" yaml:"federationParameters"`
	// The KMS key id used to encrypt in the FinSpace environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-finspace-environment.html#cfn-finspace-environment-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Configuration information for the superuser.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-finspace-environment.html#cfn-finspace-environment-superuserparameters
	//
	SuperuserParameters interface{} `field:"optional" json:"superuserParameters" yaml:"superuserParameters"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-finspace-environment.html#cfn-finspace-environment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

