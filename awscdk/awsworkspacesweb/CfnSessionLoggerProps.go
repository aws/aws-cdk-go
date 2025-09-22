package awsworkspacesweb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSessionLogger`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   cfnSessionLoggerProps := &CfnSessionLoggerProps{
//   	EventFilter: &EventFilterProperty{
//   		All: all,
//   		Include: []*string{
//   			jsii.String("include"),
//   		},
//   	},
//   	LogConfiguration: &LogConfigurationProperty{
//   		S3: &S3LogConfigurationProperty{
//   			Bucket: jsii.String("bucket"),
//   			FolderStructure: jsii.String("folderStructure"),
//   			LogFileFormat: jsii.String("logFileFormat"),
//
//   			// the properties below are optional
//   			BucketOwner: jsii.String("bucketOwner"),
//   			KeyPrefix: jsii.String("keyPrefix"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AdditionalEncryptionContext: map[string]*string{
//   		"additionalEncryptionContextKey": jsii.String("additionalEncryptionContext"),
//   	},
//   	CustomerManagedKey: jsii.String("customerManagedKey"),
//   	DisplayName: jsii.String("displayName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-sessionlogger.html
//
type CfnSessionLoggerProps struct {
	// The filter that specifies which events to monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-sessionlogger.html#cfn-workspacesweb-sessionlogger-eventfilter
	//
	EventFilter interface{} `field:"required" json:"eventFilter" yaml:"eventFilter"`
	// The configuration that specifies where logs are fowarded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-sessionlogger.html#cfn-workspacesweb-sessionlogger-logconfiguration
	//
	LogConfiguration interface{} `field:"required" json:"logConfiguration" yaml:"logConfiguration"`
	// The additional encryption context of the session logger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-sessionlogger.html#cfn-workspacesweb-sessionlogger-additionalencryptioncontext
	//
	AdditionalEncryptionContext interface{} `field:"optional" json:"additionalEncryptionContext" yaml:"additionalEncryptionContext"`
	// The custom managed key of the session logger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-sessionlogger.html#cfn-workspacesweb-sessionlogger-customermanagedkey
	//
	CustomerManagedKey *string `field:"optional" json:"customerManagedKey" yaml:"customerManagedKey"`
	// The human-readable display name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-sessionlogger.html#cfn-workspacesweb-sessionlogger-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-sessionlogger.html#cfn-workspacesweb-sessionlogger-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

