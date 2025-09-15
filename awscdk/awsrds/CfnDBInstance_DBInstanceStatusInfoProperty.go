package awsrds


// Provides a list of status information for a DB instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBInstanceStatusInfoProperty := &DBInstanceStatusInfoProperty{
//   	Message: jsii.String("message"),
//   	Normal: jsii.Boolean(false),
//   	Status: jsii.String("status"),
//   	StatusType: jsii.String("statusType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-dbinstancestatusinfo.html
//
type CfnDBInstance_DBInstanceStatusInfoProperty struct {
	// Details of the error if there is an error for the instance.
	//
	// If the instance isn't in an error state, this value is blank.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-dbinstancestatusinfo.html#cfn-rds-dbinstance-dbinstancestatusinfo-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Indicates whether the instance is operating normally (TRUE) or is in an error state (FALSE).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-dbinstancestatusinfo.html#cfn-rds-dbinstance-dbinstancestatusinfo-normal
	//
	Normal interface{} `field:"optional" json:"normal" yaml:"normal"`
	// The status of the DB instance.
	//
	// For a StatusType of read replica, the values can be replicating, replication stop point set, replication stop point reached, error, stopped, or terminated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-dbinstancestatusinfo.html#cfn-rds-dbinstance-dbinstancestatusinfo-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// This value is currently "read replication.".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-dbinstancestatusinfo.html#cfn-rds-dbinstance-dbinstancestatusinfo-statustype
	//
	StatusType *string `field:"optional" json:"statusType" yaml:"statusType"`
}

