package awslogs


// Properties for creating string mutator processors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringMutatorProps := &StringMutatorProps{
//   	Type: awscdk.Aws_logs.StringMutatorType_LOWER_CASE,
//
//   	// the properties below are optional
//   	LowerCaseKeys: []*string{
//   		jsii.String("lowerCaseKeys"),
//   	},
//   	SplitOptions: &SplitStringProperty{
//   		Entries: []splitStringEntryProperty{
//   			&splitStringEntryProperty{
//   				Delimiter: awscdk.*Aws_logs.DelimiterCharacter_COMMA,
//   				Source: jsii.String("source"),
//   			},
//   		},
//   	},
//   	SubstituteOptions: &SubstituteStringProperty{
//   		Entries: []substituteStringEntryProperty{
//   			&substituteStringEntryProperty{
//   				From: jsii.String("from"),
//   				Source: jsii.String("source"),
//   				To: jsii.String("to"),
//   			},
//   		},
//   	},
//   	TrimKeys: []*string{
//   		jsii.String("trimKeys"),
//   	},
//   	UpperCaseKeys: []*string{
//   		jsii.String("upperCaseKeys"),
//   	},
//   }
//
type StringMutatorProps struct {
	// The type of string mutation operation.
	Type StringMutatorType `field:"required" json:"type" yaml:"type"`
	// Keys for strings to convert to lowercase.
	//
	// Required when type is LOWER_CASE.
	// Default: - No lowercase processor is created if props not set.
	//
	LowerCaseKeys *[]*string `field:"optional" json:"lowerCaseKeys" yaml:"lowerCaseKeys"`
	// Options for string splitting.
	//
	// Required when type is SPLIT.
	// Default: - No string splitting processor is created if props not set.
	//
	SplitOptions *SplitStringProperty `field:"optional" json:"splitOptions" yaml:"splitOptions"`
	// Options for string substitution.
	//
	// Required when type is SUBSTITUTE.
	// Default: - No string substitution processor is created if props not set.
	//
	SubstituteOptions *SubstituteStringProperty `field:"optional" json:"substituteOptions" yaml:"substituteOptions"`
	// Keys for strings to trim.
	//
	// Required when type is TRIM.
	// Default: - No trim processor is created if props not set.
	//
	TrimKeys *[]*string `field:"optional" json:"trimKeys" yaml:"trimKeys"`
	// Keys for strings to convert to uppercase.
	//
	// Required when type is UPPER_CASE.
	// Default: - No uppercase processor is created if props not set.
	//
	UpperCaseKeys *[]*string `field:"optional" json:"upperCaseKeys" yaml:"upperCaseKeys"`
}

