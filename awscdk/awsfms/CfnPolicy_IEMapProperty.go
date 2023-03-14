package awsfms


// Specifies the AWS account IDs and AWS Organizations organizational units (OUs) to include in or exclude from the policy.
//
// Specifying an OU is the equivalent of specifying all accounts in the OU and in any of its child OUs, including any child OUs and accounts that are added at a later time.
//
// This is used for the policy's `IncludeMap` and `ExcludeMap` .
//
// You can specify account IDs, OUs, or a combination:
//
// - Specify account IDs by setting the key to `ACCOUNT` . For example, the following is a valid map: `{“ACCOUNT” : [“accountID1”, “accountID2”]}` .
// - Specify OUs by setting the key to `ORGUNIT` . For example, the following is a valid map: `{“ORGUNIT” : [“ouid111”, “ouid112”]}` .
// - Specify accounts and OUs together in a single map, separated with a comma. For example, the following is a valid map: `{“ACCOUNT” : [“accountID1”, “accountID2”], “ORGUNIT” : [“ouid111”, “ouid112”]}` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iEMapProperty := map[string][]*string{
//   	"account": []*string{
//   		jsii.String("account"),
//   	},
//   	"orgunit": []*string{
//   		jsii.String("orgunit"),
//   	},
//   }
//
type CfnPolicy_IEMapProperty struct {
	// The account list for the map.
	Account *[]*string `field:"optional" json:"account" yaml:"account"`
	// The organizational unit list for the map.
	Orgunit *[]*string `field:"optional" json:"orgunit" yaml:"orgunit"`
}

