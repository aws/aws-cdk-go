package awslogs


// Properties for defining a `CfnTransformer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransformerProps := &CfnTransformerProps{
//   	LogGroupIdentifier: jsii.String("logGroupIdentifier"),
//   	TransformerConfig: []interface{}{
//   		&ProcessorProperty{
//   			AddKeys: &AddKeysProperty{
//   				Entries: []interface{}{
//   					&AddKeyEntryProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//
//   						// the properties below are optional
//   						OverwriteIfExists: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			CopyValue: &CopyValueProperty{
//   				Entries: []interface{}{
//   					&CopyValueEntryProperty{
//   						Source: jsii.String("source"),
//   						Target: jsii.String("target"),
//
//   						// the properties below are optional
//   						OverwriteIfExists: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			Csv: &CsvProperty{
//   				Columns: []*string{
//   					jsii.String("columns"),
//   				},
//   				Delimiter: jsii.String("delimiter"),
//   				QuoteCharacter: jsii.String("quoteCharacter"),
//   				Source: jsii.String("source"),
//   			},
//   			DateTimeConverter: &DateTimeConverterProperty{
//   				MatchPatterns: []*string{
//   					jsii.String("matchPatterns"),
//   				},
//   				Source: jsii.String("source"),
//   				Target: jsii.String("target"),
//
//   				// the properties below are optional
//   				Locale: jsii.String("locale"),
//   				SourceTimezone: jsii.String("sourceTimezone"),
//   				TargetFormat: jsii.String("targetFormat"),
//   				TargetTimezone: jsii.String("targetTimezone"),
//   			},
//   			DeleteKeys: &DeleteKeysProperty{
//   				WithKeys: []*string{
//   					jsii.String("withKeys"),
//   				},
//   			},
//   			Grok: &GrokProperty{
//   				Match: jsii.String("match"),
//
//   				// the properties below are optional
//   				Source: jsii.String("source"),
//   			},
//   			ListToMap: &ListToMapProperty{
//   				Key: jsii.String("key"),
//   				Source: jsii.String("source"),
//
//   				// the properties below are optional
//   				Flatten: jsii.Boolean(false),
//   				FlattenedElement: jsii.String("flattenedElement"),
//   				Target: jsii.String("target"),
//   				ValueKey: jsii.String("valueKey"),
//   			},
//   			LowerCaseString: &LowerCaseStringProperty{
//   				WithKeys: []*string{
//   					jsii.String("withKeys"),
//   				},
//   			},
//   			MoveKeys: &MoveKeysProperty{
//   				Entries: []interface{}{
//   					&MoveKeyEntryProperty{
//   						Source: jsii.String("source"),
//   						Target: jsii.String("target"),
//
//   						// the properties below are optional
//   						OverwriteIfExists: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			ParseCloudfront: &ParseCloudfrontProperty{
//   				Source: jsii.String("source"),
//   			},
//   			ParseJson: &ParseJSONProperty{
//   				Destination: jsii.String("destination"),
//   				Source: jsii.String("source"),
//   			},
//   			ParseKeyValue: &ParseKeyValueProperty{
//   				Destination: jsii.String("destination"),
//   				FieldDelimiter: jsii.String("fieldDelimiter"),
//   				KeyPrefix: jsii.String("keyPrefix"),
//   				KeyValueDelimiter: jsii.String("keyValueDelimiter"),
//   				NonMatchValue: jsii.String("nonMatchValue"),
//   				OverwriteIfExists: jsii.Boolean(false),
//   				Source: jsii.String("source"),
//   			},
//   			ParsePostgres: &ParsePostgresProperty{
//   				Source: jsii.String("source"),
//   			},
//   			ParseRoute53: &ParseRoute53Property{
//   				Source: jsii.String("source"),
//   			},
//   			ParseVpc: &ParseVPCProperty{
//   				Source: jsii.String("source"),
//   			},
//   			ParseWaf: &ParseWAFProperty{
//   				Source: jsii.String("source"),
//   			},
//   			RenameKeys: &RenameKeysProperty{
//   				Entries: []interface{}{
//   					&RenameKeyEntryProperty{
//   						Key: jsii.String("key"),
//   						RenameTo: jsii.String("renameTo"),
//
//   						// the properties below are optional
//   						OverwriteIfExists: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			SplitString: &SplitStringProperty{
//   				Entries: []interface{}{
//   					&SplitStringEntryProperty{
//   						Delimiter: jsii.String("delimiter"),
//   						Source: jsii.String("source"),
//   					},
//   				},
//   			},
//   			SubstituteString: &SubstituteStringProperty{
//   				Entries: []interface{}{
//   					&SubstituteStringEntryProperty{
//   						From: jsii.String("from"),
//   						Source: jsii.String("source"),
//   						To: jsii.String("to"),
//   					},
//   				},
//   			},
//   			TrimString: &TrimStringProperty{
//   				WithKeys: []*string{
//   					jsii.String("withKeys"),
//   				},
//   			},
//   			TypeConverter: &TypeConverterProperty{
//   				Entries: []interface{}{
//   					&TypeConverterEntryProperty{
//   						Key: jsii.String("key"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			UpperCaseString: &UpperCaseStringProperty{
//   				WithKeys: []*string{
//   					jsii.String("withKeys"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-transformer.html
//
type CfnTransformerProps struct {
	// Specify either the name or ARN of the log group to create the transformer for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-transformer.html#cfn-logs-transformer-loggroupidentifier
	//
	LogGroupIdentifier *string `field:"required" json:"logGroupIdentifier" yaml:"logGroupIdentifier"`
	// This structure is an array that contains the configuration of this log transformer.
	//
	// A log transformer is an array of processors, where each processor applies one type of transformation to the log events that are ingested.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-transformer.html#cfn-logs-transformer-transformerconfig
	//
	TransformerConfig interface{} `field:"required" json:"transformerConfig" yaml:"transformerConfig"`
}

