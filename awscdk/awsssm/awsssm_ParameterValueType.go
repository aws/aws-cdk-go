package awsssm


// The type of CFN SSM Parameter.
//
// Using specific types can be helpful in catching invalid values
// at the start of creating or updating a stack. CloudFormation validates
// the values against existing values in the account.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   ssm.stringParameter.valueForTypedStringParameterV2(stack, jsii.String("/My/Public/Parameter"), ssm.parameterValueType_AWS_EC2_IMAGE_ID)
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html#aws-ssm-parameter-types
//
type ParameterValueType string

const (
	// String.
	ParameterValueType_STRING ParameterValueType = "STRING"
	// An Availability Zone, such as us-west-2a.
	ParameterValueType_AWS_EC2_AVAILABILITYZONE_NAME ParameterValueType = "AWS_EC2_AVAILABILITYZONE_NAME"
	// An Amazon EC2 image ID, such as ami-0ff8a91507f77f867.
	ParameterValueType_AWS_EC2_IMAGE_ID ParameterValueType = "AWS_EC2_IMAGE_ID"
	// An Amazon EC2 instance ID, such as i-1e731a32.
	ParameterValueType_AWS_EC2_INSTANCE_ID ParameterValueType = "AWS_EC2_INSTANCE_ID"
	// An Amazon EC2 key pair name.
	ParameterValueType_AWS_EC2_KEYPAIR_KEYNAME ParameterValueType = "AWS_EC2_KEYPAIR_KEYNAME"
	// An EC2-Classic or default VPC security group name, such as my-sg-abc.
	ParameterValueType_AWS_EC2_SECURITYGROUP_GROUPNAME ParameterValueType = "AWS_EC2_SECURITYGROUP_GROUPNAME"
	// A security group ID, such as sg-a123fd85.
	ParameterValueType_AWS_EC2_SECURITYGROUP_ID ParameterValueType = "AWS_EC2_SECURITYGROUP_ID"
	// A subnet ID, such as subnet-123a351e.
	ParameterValueType_AWS_EC2_SUBNET_ID ParameterValueType = "AWS_EC2_SUBNET_ID"
	// An Amazon EBS volume ID, such as vol-3cdd3f56.
	ParameterValueType_AWS_EC2_VOLUME_ID ParameterValueType = "AWS_EC2_VOLUME_ID"
	// A VPC ID, such as vpc-a123baa3.
	ParameterValueType_AWS_EC2_VPC_ID ParameterValueType = "AWS_EC2_VPC_ID"
	// An Amazon Route 53 hosted zone ID, such as Z23YXV4OVPL04A.
	ParameterValueType_AWS_ROUTE53_HOSTEDZONE_ID ParameterValueType = "AWS_ROUTE53_HOSTEDZONE_ID"
)

