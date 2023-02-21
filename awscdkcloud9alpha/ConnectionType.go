// The CDK Construct Library for AWS::Cloud9
package awscdkcloud9alpha


// The connection type used for connecting to an Amazon EC2 environment.
// Experimental.
type ConnectionType string

const (
	// Connect through SSH.
	// Experimental.
	ConnectionType_CONNECT_SSH ConnectionType = "CONNECT_SSH"
	// Connect through AWS Systems Manager When using SSM, service role and instance profile aren't automatically created.
	//
	// See https://docs.aws.amazon.com/cloud9/latest/user-guide/ec2-ssm.html#service-role-ssm
	// Experimental.
	ConnectionType_CONNECT_SSM ConnectionType = "CONNECT_SSM"
)

