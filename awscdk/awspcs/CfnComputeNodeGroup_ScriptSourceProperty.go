package awspcs


// The external location of a lifecycle script.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scriptSourceProperty := &ScriptSourceProperty{
//   	ScriptLocation: jsii.String("scriptLocation"),
//
//   	// the properties below are optional
//   	Checksum: jsii.String("checksum"),
//   	S3VersionId: jsii.String("s3VersionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-scriptsource.html
//
type CfnComputeNodeGroup_ScriptSourceProperty struct {
	// The S3 URI or HTTPS URL where the script is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-scriptsource.html#cfn-pcs-computenodegroup-scriptsource-scriptlocation
	//
	ScriptLocation *string `field:"required" json:"scriptLocation" yaml:"scriptLocation"`
	// A 64-character hexadecimal SHA-256 digest used to verify script integrity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-scriptsource.html#cfn-pcs-computenodegroup-scriptsource-checksum
	//
	Checksum *string `field:"optional" json:"checksum" yaml:"checksum"`
	// The S3 object version ID of the script, when stored in a versioned bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-computenodegroup-scriptsource.html#cfn-pcs-computenodegroup-scriptsource-s3versionid
	//
	S3VersionId *string `field:"optional" json:"s3VersionId" yaml:"s3VersionId"`
}

