// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha


// A full specification of an alias that can be used to import it fluently into the CDK application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//
//   aliasAttributes := &AliasAttributes{
//   	AliasArn: jsii.String("aliasArn"),
//   	AliasId: jsii.String("aliasId"),
//   }
//
// Experimental.
type AliasAttributes struct {
	// The ARN of the alias.
	//
	// At least one of `aliasArn` and `aliasId` must be provided.
	// Experimental.
	AliasArn *string `field:"optional" json:"aliasArn" yaml:"aliasArn"`
	// The identifier of the alias.
	//
	// At least one of `aliasId` and `aliasArn`  must be provided.
	// Experimental.
	AliasId *string `field:"optional" json:"aliasId" yaml:"aliasId"`
}

