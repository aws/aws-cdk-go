package previewawsdataexchangeevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.dataexchange@DataSetUpdateDelayed event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetUpdateDelayedProps := &DataSetUpdateDelayedProps{
//   	DataSet: &DataSet{
//   		AssetType: []*string{
//   			jsii.String("assetType"),
//   		},
//   		Id: []*string{
//   			jsii.String("id"),
//   		},
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
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
//   	Notification: &Notification{
//   		Comment: []*string{
//   			jsii.String("comment"),
//   		},
//   		Scope: &Scope{
//   			LakeFormationTagPolicies: []LakeFormationTagPolicyDetails{
//   				&LakeFormationTagPolicyDetails{
//   					Database: []*string{
//   						jsii.String("database"),
//   					},
//   					Table: []*string{
//   						jsii.String("table"),
//   					},
//   				},
//   			},
//   			RedshiftDataShares: []RedshiftDataShareDetails{
//   				&RedshiftDataShareDetails{
//   					Arn: []*string{
//   						jsii.String("arn"),
//   					},
//   					Database: []*string{
//   						jsii.String("database"),
//   					},
//   					Function: []*string{
//   						jsii.String("function"),
//   					},
//   					Table: []*string{
//   						jsii.String("table"),
//   					},
//   					View: []*string{
//   						jsii.String("view"),
//   					},
//   				},
//   			},
//   			S3DataAccesses: []S3DataAccessDetails{
//   				&S3DataAccessDetails{
//   					KeyPrefixes: []*string{
//   						jsii.String("keyPrefixes"),
//   					},
//   					Keys: []*string{
//   						jsii.String("keys"),
//   					},
//   				},
//   			},
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   	Product: &Product{
//   		Id: []*string{
//   			jsii.String("id"),
//   		},
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   		ProviderContact: []*string{
//   			jsii.String("providerContact"),
//   		},
//   	},
//   }
//
// Experimental.
type DataSetUpdateDelayed_DataSetUpdateDelayedProps struct {
	// DataSet property.
	//
	// Specify an array of string values to match this event if the actual value of DataSet is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DataSet *DataSetUpdateDelayed_DataSet `field:"optional" json:"dataSet" yaml:"dataSet"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// Notification property.
	//
	// Specify an array of string values to match this event if the actual value of Notification is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Notification *DataSetUpdateDelayed_Notification `field:"optional" json:"notification" yaml:"notification"`
	// Product property.
	//
	// Specify an array of string values to match this event if the actual value of Product is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Product *DataSetUpdateDelayed_Product `field:"optional" json:"product" yaml:"product"`
}

