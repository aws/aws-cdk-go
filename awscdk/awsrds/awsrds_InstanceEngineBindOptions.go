package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// The options passed to {@link IInstanceEngine.bind}.
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
//   instanceEngineBindOptions := &instanceEngineBindOptions{
//   	domain: jsii.String("domain"),
//   	optionGroup: optionGroup,
//   	s3ExportRole: role,
//   	s3ImportRole: role,
//   	timezone: jsii.String("timezone"),
//   }
//
// Experimental.
type InstanceEngineBindOptions struct {
	// The Active Directory directory ID to create the DB instance in.
	// Experimental.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The option group of the database.
	// Experimental.
	OptionGroup IOptionGroup `field:"optional" json:"optionGroup" yaml:"optionGroup"`
	// The role used for S3 exporting.
	// Experimental.
	S3ExportRole awsiam.IRole `field:"optional" json:"s3ExportRole" yaml:"s3ExportRole"`
	// The role used for S3 importing.
	// Experimental.
	S3ImportRole awsiam.IRole `field:"optional" json:"s3ImportRole" yaml:"s3ImportRole"`
	// The timezone of the database, set by the customer.
	// Experimental.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

