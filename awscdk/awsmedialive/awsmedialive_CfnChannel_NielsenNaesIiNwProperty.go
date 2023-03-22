package awsmedialive


// Complete these fields only if you want to insert watermarks of type Nielsen NAES II (N2) and Nielsen NAES VI (NW).
//
// The parent of this entity is NielsenWatermarksSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nielsenNaesIiNwProperty := &nielsenNaesIiNwProperty{
//   	checkDigitString: jsii.String("checkDigitString"),
//   	sid: jsii.Number(123),
//   }
//
type CfnChannel_NielsenNaesIiNwProperty struct {
	// Enter the check digit string for the watermark.
	CheckDigitString *string `field:"optional" json:"checkDigitString" yaml:"checkDigitString"`
	// Enter the Nielsen Source ID (SID) to include in the watermark.
	Sid *float64 `field:"optional" json:"sid" yaml:"sid"`
}

