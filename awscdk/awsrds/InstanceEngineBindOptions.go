package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// The options passed to `IInstanceEngine.bind`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var optionGroup optionGroup
//   var role role
//
//   instanceEngineBindOptions := &InstanceEngineBindOptions{
//   	Domain: jsii.String("domain"),
//   	OptionGroup: optionGroup,
//   	S3ExportRole: role,
//   	S3ImportRole: role,
//   	Timezone: jsii.String("timezone"),
//   }
//
type InstanceEngineBindOptions struct {
	// The Active Directory directory ID to create the DB instance in.
	// Default: - none (it's an optional field).
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The option group of the database.
	// Default: - none.
	//
	OptionGroup IOptionGroup `field:"optional" json:"optionGroup" yaml:"optionGroup"`
	// The role used for S3 exporting.
	// Default: - none.
	//
	S3ExportRole awsiam.IRole `field:"optional" json:"s3ExportRole" yaml:"s3ExportRole"`
	// The role used for S3 importing.
	// Default: - none.
	//
	S3ImportRole awsiam.IRole `field:"optional" json:"s3ImportRole" yaml:"s3ImportRole"`
	// The timezone of the database, set by the customer.
	// Default: - none (it's an optional field).
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

