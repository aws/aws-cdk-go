package previewawsdataexchangeevents


// Type definition for Notification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   notification := &Notification{
//   	Comment: []*string{
//   		jsii.String("comment"),
//   	},
//   	Details: &Details{
//   		SchemaChange: &SchemaChange{
//   			Changes: []SchemaChangeItem{
//   				&SchemaChangeItem{
//   					Description: []*string{
//   						jsii.String("description"),
//   					},
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   					Type: []*string{
//   						jsii.String("type"),
//   					},
//   				},
//   			},
//   			SchemaChangeAt: []*string{
//   				jsii.String("schemaChangeAt"),
//   			},
//   		},
//   	},
//   	Scope: &Scope{
//   		LakeFormationTagPolicies: []LakeFormationTagPolicyDetails{
//   			&LakeFormationTagPolicyDetails{
//   				Database: []*string{
//   					jsii.String("database"),
//   				},
//   				Table: []*string{
//   					jsii.String("table"),
//   				},
//   			},
//   		},
//   		RedshiftDataShares: []RedshiftDataShareDetails{
//   			&RedshiftDataShareDetails{
//   				Arn: []*string{
//   					jsii.String("arn"),
//   				},
//   				Database: []*string{
//   					jsii.String("database"),
//   				},
//   				Function: []*string{
//   					jsii.String("function"),
//   				},
//   				Schema: []*string{
//   					jsii.String("schema"),
//   				},
//   				Table: []*string{
//   					jsii.String("table"),
//   				},
//   				View: []*string{
//   					jsii.String("view"),
//   				},
//   			},
//   		},
//   		S3DataAccesses: []S3DataAccessDetails{
//   			&S3DataAccessDetails{
//   				KeyPrefixes: []*string{
//   					jsii.String("keyPrefixes"),
//   				},
//   				Keys: []*string{
//   					jsii.String("keys"),
//   				},
//   			},
//   		},
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   }
//
// Experimental.
type SchemaChangePlannedForDataSet_Notification struct {
	// Comment property.
	//
	// Specify an array of string values to match this event if the actual value of Comment is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Comment *[]*string `field:"optional" json:"comment" yaml:"comment"`
	// Details property.
	//
	// Specify an array of string values to match this event if the actual value of Details is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Details *SchemaChangePlannedForDataSet_Details `field:"optional" json:"details" yaml:"details"`
	// Scope property.
	//
	// Specify an array of string values to match this event if the actual value of Scope is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Scope *SchemaChangePlannedForDataSet_Scope `field:"optional" json:"scope" yaml:"scope"`
	// Type property.
	//
	// Specify an array of string values to match this event if the actual value of Type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
}

