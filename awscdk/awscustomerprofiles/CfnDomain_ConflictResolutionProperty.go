package awscustomerprofiles


// Determines how the auto-merging process should resolve conflicts between different profiles.
//
// For example, if Profile A and Profile B have the same `FirstName` and `LastName` , `ConflictResolution` specifies which `EmailAddress` should be used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conflictResolutionProperty := &ConflictResolutionProperty{
//   	ConflictResolvingModel: jsii.String("conflictResolvingModel"),
//
//   	// the properties below are optional
//   	SourceName: jsii.String("sourceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-conflictresolution.html
//
type CfnDomain_ConflictResolutionProperty struct {
	// How the auto-merging process should resolve conflicts between different profiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-conflictresolution.html#cfn-customerprofiles-domain-conflictresolution-conflictresolvingmodel
	//
	ConflictResolvingModel *string `field:"required" json:"conflictResolvingModel" yaml:"conflictResolvingModel"`
	// The `ObjectType` name that is used to resolve profile merging conflicts when choosing `SOURCE` as the `ConflictResolvingModel` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-conflictresolution.html#cfn-customerprofiles-domain-conflictresolution-sourcename
	//
	SourceName *string `field:"optional" json:"sourceName" yaml:"sourceName"`
}

