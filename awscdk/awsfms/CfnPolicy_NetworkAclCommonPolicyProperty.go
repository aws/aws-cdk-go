package awsfms


// Defines a Firewall Manager network ACL policy.
//
// This is used in the `PolicyOption` of a `SecurityServicePolicyData` for a `Policy` , when the `SecurityServicePolicyData` type is set to `NETWORK_ACL_COMMON` .
//
// For information about network ACLs, see [Control traffic to subnets using network ACLs](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-network-acls.html) in the *Amazon Virtual Private Cloud User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkAclCommonPolicyProperty := &NetworkAclCommonPolicyProperty{
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclcommonpolicy.html
//
type CfnPolicy_NetworkAclCommonPolicyProperty struct {
}

