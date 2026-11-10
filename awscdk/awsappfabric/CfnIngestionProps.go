package awsappfabric

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIngestion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIngestionProps := &CfnIngestionProps{
//   	App: jsii.String("app"),
//   	AppBundleArn: jsii.String("appBundleArn"),
//   	IngestionType: jsii.String("ingestionType"),
//   	TenantId: jsii.String("tenantId"),
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-ingestion.html
//
type CfnIngestionProps struct {
	// The name of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-ingestion.html#cfn-appfabric-ingestion-app
	//
	App *string `field:"required" json:"app" yaml:"app"`
	// The Amazon Resource Name (ARN) of the app bundle for the ingestion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-ingestion.html#cfn-appfabric-ingestion-appbundlearn
	//
	AppBundleArn *string `field:"required" json:"appBundleArn" yaml:"appBundleArn"`
	// The type of the ingestion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-ingestion.html#cfn-appfabric-ingestion-ingestiontype
	//
	IngestionType *string `field:"required" json:"ingestionType" yaml:"ingestionType"`
	// The ID of the application tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-ingestion.html#cfn-appfabric-ingestion-tenantid
	//
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// A map of the key-value pairs of the tag or tags to assign to the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appfabric-ingestion.html#cfn-appfabric-ingestion-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

