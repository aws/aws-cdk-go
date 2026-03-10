package previewawssagemakerevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.sagemaker@SageMakerNotebookInstanceStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerNotebookInstanceStateChangeProps := &SageMakerNotebookInstanceStateChangeProps{
//   	CreationTime: []*string{
//   		jsii.String("creationTime"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	InstanceType: []*string{
//   		jsii.String("instanceType"),
//   	},
//   	LastModifiedTime: []*string{
//   		jsii.String("lastModifiedTime"),
//   	},
//   	NotebookInstanceArn: []*string{
//   		jsii.String("notebookInstanceArn"),
//   	},
//   	NotebookInstanceLifecycleConfigName: []*string{
//   		jsii.String("notebookInstanceLifecycleConfigName"),
//   	},
//   	NotebookInstanceName: []*string{
//   		jsii.String("notebookInstanceName"),
//   	},
//   	NotebookInstanceStatus: []*string{
//   		jsii.String("notebookInstanceStatus"),
//   	},
//   	RoleArn: []*string{
//   		jsii.String("roleArn"),
//   	},
//   	Tags: []Tags{
//   		&Tags{
//   			Key: []*string{
//   				jsii.String("key"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type SageMakerNotebookInstanceStateChange_SageMakerNotebookInstanceStateChangeProps struct {
	// CreationTime property.
	//
	// Specify an array of string values to match this event if the actual value of CreationTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreationTime *[]*string `field:"optional" json:"creationTime" yaml:"creationTime"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// InstanceType property.
	//
	// Specify an array of string values to match this event if the actual value of InstanceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceType *[]*string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// LastModifiedTime property.
	//
	// Specify an array of string values to match this event if the actual value of LastModifiedTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastModifiedTime *[]*string `field:"optional" json:"lastModifiedTime" yaml:"lastModifiedTime"`
	// NotebookInstanceArn property.
	//
	// Specify an array of string values to match this event if the actual value of NotebookInstanceArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NotebookInstanceArn *[]*string `field:"optional" json:"notebookInstanceArn" yaml:"notebookInstanceArn"`
	// NotebookInstanceLifecycleConfigName property.
	//
	// Specify an array of string values to match this event if the actual value of NotebookInstanceLifecycleConfigName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NotebookInstanceLifecycleConfigName *[]*string `field:"optional" json:"notebookInstanceLifecycleConfigName" yaml:"notebookInstanceLifecycleConfigName"`
	// NotebookInstanceName property.
	//
	// Specify an array of string values to match this event if the actual value of NotebookInstanceName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NotebookInstanceName *[]*string `field:"optional" json:"notebookInstanceName" yaml:"notebookInstanceName"`
	// NotebookInstanceStatus property.
	//
	// Specify an array of string values to match this event if the actual value of NotebookInstanceStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NotebookInstanceStatus *[]*string `field:"optional" json:"notebookInstanceStatus" yaml:"notebookInstanceStatus"`
	// RoleArn property.
	//
	// Specify an array of string values to match this event if the actual value of RoleArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RoleArn *[]*string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Tags property.
	//
	// Specify an array of string values to match this event if the actual value of Tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]*SageMakerNotebookInstanceStateChange_Tags `field:"optional" json:"tags" yaml:"tags"`
}

