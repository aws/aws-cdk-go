package awsresiliencehubv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnServicePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnServiceMixinProps := &CfnServiceMixinProps{
//   	Assertions: []interface{}{
//   		&AssertionDefinitionProperty{
//   			Text: jsii.String("text"),
//   		},
//   	},
//   	AssociatedSystems: []interface{}{
//   		&AssociatedSystemProperty{
//   			SystemArn: jsii.String("systemArn"),
//   			UserJourneyIds: []*string{
//   				jsii.String("userJourneyIds"),
//   			},
//   		},
//   	},
//   	DependencyDiscovery: jsii.String("dependencyDiscovery"),
//   	Description: jsii.String("description"),
//   	InputSources: []interface{}{
//   		&InputSourceDefinitionProperty{
//   			ResourceConfiguration: &ResourceConfigurationProperty{
//   				CfnStackArn: jsii.String("cfnStackArn"),
//   				DesignFileS3Url: jsii.String("designFileS3Url"),
//   				Eks: &EksSourceProperty{
//   					ClusterArn: jsii.String("clusterArn"),
//   					Namespaces: []*string{
//   						jsii.String("namespaces"),
//   					},
//   				},
//   				ResourceTags: []interface{}{
//   					&ResourceTagProperty{
//   						Key: jsii.String("key"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				TfStateFileUrl: jsii.String("tfStateFileUrl"),
//   			},
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Name: jsii.String("name"),
//   	PermissionModel: &PermissionModelProperty{
//   		CrossAccountRoleArns: []interface{}{
//   			&CrossAccountRoleConfigurationProperty{
//   				CrossAccountRoleArn: jsii.String("crossAccountRoleArn"),
//   				ExternalId: jsii.String("externalId"),
//   			},
//   		},
//   		InvokerRoleName: jsii.String("invokerRoleName"),
//   	},
//   	PolicyArn: jsii.String("policyArn"),
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   	ReportConfiguration: &ServiceReportConfigurationProperty{
//   		ReportOutput: []interface{}{
//   			&ReportOutputConfigurationProperty{
//   				S3: &S3ReportOutputConfigurationProperty{
//   					BucketOwner: jsii.String("bucketOwner"),
//   					BucketPath: jsii.String("bucketPath"),
//   				},
//   			},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-service.html
//
type CfnServiceMixinProps struct {
	// Assertions associated with this service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-service.html#cfn-resiliencehubv2-service-assertions
	//
	Assertions interface{} `field:"optional" json:"assertions" yaml:"assertions"`
	// Systems associated with this service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-service.html#cfn-resiliencehubv2-service-associatedsystems
	//
	AssociatedSystems interface{} `field:"optional" json:"associatedSystems" yaml:"associatedSystems"`
	// Dependency discovery state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-service.html#cfn-resiliencehubv2-service-dependencydiscovery
	//
	DependencyDiscovery *string `field:"optional" json:"dependencyDiscovery" yaml:"dependencyDiscovery"`
	// The description of the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-service.html#cfn-resiliencehubv2-service-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Input sources for this service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-service.html#cfn-resiliencehubv2-service-inputsources
	//
	InputSources interface{} `field:"optional" json:"inputSources" yaml:"inputSources"`
	// The KMS key ID for encrypting service data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-service.html#cfn-resiliencehubv2-service-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-service.html#cfn-resiliencehubv2-service-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-service.html#cfn-resiliencehubv2-service-permissionmodel
	//
	PermissionModel interface{} `field:"optional" json:"permissionModel" yaml:"permissionModel"`
	// The ARN of the resilience policy to associate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-service.html#cfn-resiliencehubv2-service-policyarn
	//
	PolicyArn *string `field:"optional" json:"policyArn" yaml:"policyArn"`
	// AWS regions for the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-service.html#cfn-resiliencehubv2-service-regions
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Configuration for automatic report generation on a Service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-service.html#cfn-resiliencehubv2-service-reportconfiguration
	//
	ReportConfiguration interface{} `field:"optional" json:"reportConfiguration" yaml:"reportConfiguration"`
	// Tags assigned to the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-service.html#cfn-resiliencehubv2-service-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

