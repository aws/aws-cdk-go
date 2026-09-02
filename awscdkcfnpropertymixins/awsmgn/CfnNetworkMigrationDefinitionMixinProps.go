package awsmgn

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnNetworkMigrationDefinitionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnNetworkMigrationDefinitionMixinProps := &CfnNetworkMigrationDefinitionMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	ScopeTags: map[string]*string{
//   		"scopeTagsKey": jsii.String("scopeTags"),
//   	},
//   	SourceConfigurations: []interface{}{
//   		&SourceConfigurationProperty{
//   			SourceEnvironment: jsii.String("sourceEnvironment"),
//   			SourceS3Configuration: &SourceS3ConfigurationProperty{
//   				S3Bucket: jsii.String("s3Bucket"),
//   				S3BucketOwner: jsii.String("s3BucketOwner"),
//   				S3Key: jsii.String("s3Key"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetDeployment: jsii.String("targetDeployment"),
//   	TargetNetwork: &TargetNetworkProperty{
//   		InboundCidr: jsii.String("inboundCidr"),
//   		InspectionCidr: jsii.String("inspectionCidr"),
//   		OutboundCidr: jsii.String("outboundCidr"),
//   		Topology: jsii.String("topology"),
//   	},
//   	TargetS3Configuration: &TargetS3ConfigurationProperty{
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3BucketOwner: jsii.String("s3BucketOwner"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html
//
type CfnNetworkMigrationDefinitionMixinProps struct {
	// A description of the network migration definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html#cfn-mgn-networkmigrationdefinition-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the network migration definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html#cfn-mgn-networkmigrationdefinition-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Scope tags map for the network migration definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html#cfn-mgn-networkmigrationdefinition-scopetags
	//
	ScopeTags interface{} `field:"optional" json:"scopeTags" yaml:"scopeTags"`
	// A list of source configurations for the network migration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html#cfn-mgn-networkmigrationdefinition-sourceconfigurations
	//
	SourceConfigurations interface{} `field:"optional" json:"sourceConfigurations" yaml:"sourceConfigurations"`
	// Tags to assign to the network migration definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html#cfn-mgn-networkmigrationdefinition-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The target deployment configuration for the migrated network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html#cfn-mgn-networkmigrationdefinition-targetdeployment
	//
	TargetDeployment *string `field:"optional" json:"targetDeployment" yaml:"targetDeployment"`
	// Configuration for the target network topology and addressing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html#cfn-mgn-networkmigrationdefinition-targetnetwork
	//
	TargetNetwork interface{} `field:"optional" json:"targetNetwork" yaml:"targetNetwork"`
	// S3 configuration for storing target network artifacts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html#cfn-mgn-networkmigrationdefinition-targets3configuration
	//
	TargetS3Configuration interface{} `field:"optional" json:"targetS3Configuration" yaml:"targetS3Configuration"`
}

