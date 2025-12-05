package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataTableProps := &CfnDataTableProps{
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
type CfnDataTableProps struct {
	// The description of the Data Table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatable.html#cfn-connect-datatable-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The identifier of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatable.html#cfn-connect-datatable-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// The name of the Data Table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatable.html#cfn-connect-datatable-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The status of the Data Table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatable.html#cfn-connect-datatable-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// One or more tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatable.html#cfn-connect-datatable-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The time zone of the Data Table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatable.html#cfn-connect-datatable-timezone
	//
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
	// The value lock level of the Data Table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatable.html#cfn-connect-datatable-valuelocklevel
	//
	ValueLockLevel *string `field:"optional" json:"valueLockLevel" yaml:"valueLockLevel"`
}

