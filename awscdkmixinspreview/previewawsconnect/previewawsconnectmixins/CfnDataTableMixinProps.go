package previewawsconnectmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDataTablePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataTableMixinProps := &CfnDataTableMixinProps{
//   	Description: jsii.String("description"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeZone: jsii.String("timeZone"),
//   	ValueLockLevel: jsii.String("valueLockLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatable.html
//
type CfnDataTableMixinProps struct {
	// An optional description of the data table's purpose and contents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatable.html#cfn-connect-datatable-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatable.html#cfn-connect-datatable-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// The human-readable name of the data table.
	//
	// Must be unique within the instance and conform to Connect naming standards.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatable.html#cfn-connect-datatable-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The current status of the data table.
	//
	// One of PUBLISHED or SAVED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatable.html#cfn-connect-datatable-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Key-value pairs for attribute based access control (TBAC or ABAC) and organization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatable.html#cfn-connect-datatable-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The IANA timezone identifier used when resolving time based dynamic values.
	//
	// Required even if no time slices are specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatable.html#cfn-connect-datatable-timezone
	//
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
	// The data level that concurrent value edits are locked on.
	//
	// One of DATA_TABLE, PRIMARY_VALUE, ATTRIBUTE, VALUE, and NONE. Determines how concurrent edits are handled when multiple users attempt to modify values simultaneously.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatable.html#cfn-connect-datatable-valuelocklevel
	//
	ValueLockLevel *string `field:"optional" json:"valueLockLevel" yaml:"valueLockLevel"`
}

