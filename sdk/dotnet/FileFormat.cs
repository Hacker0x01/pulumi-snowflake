// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Snowflake
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Snowflake = Pulumi.Snowflake;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleFileFormat = new Snowflake.FileFormat("exampleFileFormat", new Snowflake.FileFormatArgs
    ///         {
    ///             Database = "EXAMPLE_DB",
    ///             FormatType = "CSV",
    ///             Schema = "EXAMPLE_SCHEMA",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// # format is database name | schema name | file format name
    /// 
    /// ```sh
    ///  $ pulumi import snowflake:index/fileFormat:FileFormat example 'dbName|schemaName|fileFormatName'
    /// ```
    /// </summary>
    [SnowflakeResourceType("snowflake:index/fileFormat:FileFormat")]
    public partial class FileFormat : Pulumi.CustomResource
    {
        /// <summary>
        /// Boolean that specifies to allow duplicate object field names (only the last one will be preserved).
        /// </summary>
        [Output("allowDuplicate")]
        public Output<bool?> AllowDuplicate { get; private set; } = null!;

        /// <summary>
        /// Boolean that specifies whether to interpret columns with no defined logical data type as UTF-8 text.
        /// </summary>
        [Output("binaryAsText")]
        public Output<bool?> BinaryAsText { get; private set; } = null!;

        /// <summary>
        /// Defines the encoding format for binary input or output.
        /// </summary>
        [Output("binaryFormat")]
        public Output<string?> BinaryFormat { get; private set; } = null!;

        /// <summary>
        /// Specifies a comment for the file format.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Specifies the current compression algorithm for the data file.
        /// </summary>
        [Output("compression")]
        public Output<string?> Compression { get; private set; } = null!;

        /// <summary>
        /// The database in which to create the file format.
        /// </summary>
        [Output("database")]
        public Output<string> Database { get; private set; } = null!;

        /// <summary>
        /// Defines the format of date values in the data files (data loading) or table (data unloading).
        /// </summary>
        [Output("dateFormat")]
        public Output<string?> DateFormat { get; private set; } = null!;

        /// <summary>
        /// Boolean that specifies whether the XML parser disables automatic conversion of numeric and Boolean values from text to native representation.
        /// </summary>
        [Output("disableAutoConvert")]
        public Output<bool?> DisableAutoConvert { get; private set; } = null!;

        /// <summary>
        /// Boolean that specifies whether the XML parser disables recognition of Snowflake semi-structured data tags.
        /// </summary>
        [Output("disableSnowflakeData")]
        public Output<bool?> DisableSnowflakeData { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to insert SQL NULL for empty fields in an input file, which are represented by two successive delimiters.
        /// </summary>
        [Output("emptyFieldAsNull")]
        public Output<bool?> EmptyFieldAsNull { get; private set; } = null!;

        /// <summary>
        /// Boolean that enables parsing of octal numbers.
        /// </summary>
        [Output("enableOctal")]
        public Output<bool?> EnableOctal { get; private set; } = null!;

        /// <summary>
        /// String (constant) that specifies the character set of the source data when loading data into a table.
        /// </summary>
        [Output("encoding")]
        public Output<string?> Encoding { get; private set; } = null!;

        /// <summary>
        /// Boolean that specifies whether to generate a parsing error if the number of delimited columns (i.e. fields) in an input file does not match the number of columns in the corresponding table.
        /// </summary>
        [Output("errorOnColumnCountMismatch")]
        public Output<bool?> ErrorOnColumnCountMismatch { get; private set; } = null!;

        /// <summary>
        /// Single character string used as the escape character for field values.
        /// </summary>
        [Output("escape")]
        public Output<string?> Escape { get; private set; } = null!;

        /// <summary>
        /// Single character string used as the escape character for unenclosed field values only.
        /// </summary>
        [Output("escapeUnenclosedField")]
        public Output<string?> EscapeUnenclosedField { get; private set; } = null!;

        /// <summary>
        /// Specifies one or more singlebyte or multibyte characters that separate fields in an input file (data loading) or unloaded file (data unloading).
        /// </summary>
        [Output("fieldDelimiter")]
        public Output<string?> FieldDelimiter { get; private set; } = null!;

        /// <summary>
        /// Character used to enclose strings.
        /// </summary>
        [Output("fieldOptionallyEnclosedBy")]
        public Output<string?> FieldOptionallyEnclosedBy { get; private set; } = null!;

        /// <summary>
        /// Specifies the extension for files unloaded to a stage.
        /// </summary>
        [Output("fileExtension")]
        public Output<string?> FileExtension { get; private set; } = null!;

        /// <summary>
        /// Specifies the format of the input files (for data loading) or output files (for data unloading).
        /// </summary>
        [Output("formatType")]
        public Output<string> FormatType { get; private set; } = null!;

        /// <summary>
        /// Boolean that specifies whether UTF-8 encoding errors produce error conditions.
        /// </summary>
        [Output("ignoreUtf8Errors")]
        public Output<bool?> IgnoreUtf8Errors { get; private set; } = null!;

        /// <summary>
        /// Specifies the identifier for the file format; must be unique for the database and schema in which the file format is created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// String used to convert to and from SQL NULL.
        /// </summary>
        [Output("nullIfs")]
        public Output<ImmutableArray<string>> NullIfs { get; private set; } = null!;

        /// <summary>
        /// Boolean that specifies whether the XML parser preserves leading and trailing spaces in element content.
        /// </summary>
        [Output("preserveSpace")]
        public Output<bool?> PreserveSpace { get; private set; } = null!;

        /// <summary>
        /// Specifies one or more singlebyte or multibyte characters that separate records in an input file (data loading) or unloaded file (data unloading).
        /// </summary>
        [Output("recordDelimiter")]
        public Output<string?> RecordDelimiter { get; private set; } = null!;

        /// <summary>
        /// Boolean that specifies whether to replace invalid UTF-8 characters with the Unicode replacement character (�).
        /// </summary>
        [Output("replaceInvalidCharacters")]
        public Output<bool?> ReplaceInvalidCharacters { get; private set; } = null!;

        /// <summary>
        /// The schema in which to create the file format.
        /// </summary>
        [Output("schema")]
        public Output<string> Schema { get; private set; } = null!;

        /// <summary>
        /// Boolean that specifies to skip any blank lines encountered in the data files.
        /// </summary>
        [Output("skipBlankLines")]
        public Output<bool?> SkipBlankLines { get; private set; } = null!;

        /// <summary>
        /// Boolean that specifies whether to skip the BOM (byte order mark), if present in a data file.
        /// </summary>
        [Output("skipByteOrderMark")]
        public Output<bool?> SkipByteOrderMark { get; private set; } = null!;

        /// <summary>
        /// Number of lines at the start of the file to skip.
        /// </summary>
        [Output("skipHeader")]
        public Output<int?> SkipHeader { get; private set; } = null!;

        /// <summary>
        /// Boolean that instructs the JSON parser to remove object fields or array elements containing null values.
        /// </summary>
        [Output("stripNullValues")]
        public Output<bool?> StripNullValues { get; private set; } = null!;

        /// <summary>
        /// Boolean that instructs the JSON parser to remove outer brackets.
        /// </summary>
        [Output("stripOuterArray")]
        public Output<bool?> StripOuterArray { get; private set; } = null!;

        /// <summary>
        /// Boolean that specifies whether the XML parser strips out the outer XML element, exposing 2nd level elements as separate documents.
        /// </summary>
        [Output("stripOuterElement")]
        public Output<bool?> StripOuterElement { get; private set; } = null!;

        /// <summary>
        /// Defines the format of time values in the data files (data loading) or table (data unloading).
        /// </summary>
        [Output("timeFormat")]
        public Output<string?> TimeFormat { get; private set; } = null!;

        /// <summary>
        /// Defines the format of timestamp values in the data files (data loading) or table (data unloading).
        /// </summary>
        [Output("timestampFormat")]
        public Output<string?> TimestampFormat { get; private set; } = null!;

        /// <summary>
        /// Boolean that specifies whether to remove white space from fields.
        /// </summary>
        [Output("trimSpace")]
        public Output<bool?> TrimSpace { get; private set; } = null!;

        /// <summary>
        /// Boolean that specifies whether to validate UTF-8 character encoding in string column data.
        /// </summary>
        [Output("validateUtf8")]
        public Output<bool?> ValidateUtf8 { get; private set; } = null!;


        /// <summary>
        /// Create a FileFormat resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FileFormat(string name, FileFormatArgs args, CustomResourceOptions? options = null)
            : base("snowflake:index/fileFormat:FileFormat", name, args ?? new FileFormatArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FileFormat(string name, Input<string> id, FileFormatState? state = null, CustomResourceOptions? options = null)
            : base("snowflake:index/fileFormat:FileFormat", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FileFormat resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FileFormat Get(string name, Input<string> id, FileFormatState? state = null, CustomResourceOptions? options = null)
        {
            return new FileFormat(name, id, state, options);
        }
    }

    public sealed class FileFormatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean that specifies to allow duplicate object field names (only the last one will be preserved).
        /// </summary>
        [Input("allowDuplicate")]
        public Input<bool>? AllowDuplicate { get; set; }

        /// <summary>
        /// Boolean that specifies whether to interpret columns with no defined logical data type as UTF-8 text.
        /// </summary>
        [Input("binaryAsText")]
        public Input<bool>? BinaryAsText { get; set; }

        /// <summary>
        /// Defines the encoding format for binary input or output.
        /// </summary>
        [Input("binaryFormat")]
        public Input<string>? BinaryFormat { get; set; }

        /// <summary>
        /// Specifies a comment for the file format.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Specifies the current compression algorithm for the data file.
        /// </summary>
        [Input("compression")]
        public Input<string>? Compression { get; set; }

        /// <summary>
        /// The database in which to create the file format.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// Defines the format of date values in the data files (data loading) or table (data unloading).
        /// </summary>
        [Input("dateFormat")]
        public Input<string>? DateFormat { get; set; }

        /// <summary>
        /// Boolean that specifies whether the XML parser disables automatic conversion of numeric and Boolean values from text to native representation.
        /// </summary>
        [Input("disableAutoConvert")]
        public Input<bool>? DisableAutoConvert { get; set; }

        /// <summary>
        /// Boolean that specifies whether the XML parser disables recognition of Snowflake semi-structured data tags.
        /// </summary>
        [Input("disableSnowflakeData")]
        public Input<bool>? DisableSnowflakeData { get; set; }

        /// <summary>
        /// Specifies whether to insert SQL NULL for empty fields in an input file, which are represented by two successive delimiters.
        /// </summary>
        [Input("emptyFieldAsNull")]
        public Input<bool>? EmptyFieldAsNull { get; set; }

        /// <summary>
        /// Boolean that enables parsing of octal numbers.
        /// </summary>
        [Input("enableOctal")]
        public Input<bool>? EnableOctal { get; set; }

        /// <summary>
        /// String (constant) that specifies the character set of the source data when loading data into a table.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// Boolean that specifies whether to generate a parsing error if the number of delimited columns (i.e. fields) in an input file does not match the number of columns in the corresponding table.
        /// </summary>
        [Input("errorOnColumnCountMismatch")]
        public Input<bool>? ErrorOnColumnCountMismatch { get; set; }

        /// <summary>
        /// Single character string used as the escape character for field values.
        /// </summary>
        [Input("escape")]
        public Input<string>? Escape { get; set; }

        /// <summary>
        /// Single character string used as the escape character for unenclosed field values only.
        /// </summary>
        [Input("escapeUnenclosedField")]
        public Input<string>? EscapeUnenclosedField { get; set; }

        /// <summary>
        /// Specifies one or more singlebyte or multibyte characters that separate fields in an input file (data loading) or unloaded file (data unloading).
        /// </summary>
        [Input("fieldDelimiter")]
        public Input<string>? FieldDelimiter { get; set; }

        /// <summary>
        /// Character used to enclose strings.
        /// </summary>
        [Input("fieldOptionallyEnclosedBy")]
        public Input<string>? FieldOptionallyEnclosedBy { get; set; }

        /// <summary>
        /// Specifies the extension for files unloaded to a stage.
        /// </summary>
        [Input("fileExtension")]
        public Input<string>? FileExtension { get; set; }

        /// <summary>
        /// Specifies the format of the input files (for data loading) or output files (for data unloading).
        /// </summary>
        [Input("formatType", required: true)]
        public Input<string> FormatType { get; set; } = null!;

        /// <summary>
        /// Boolean that specifies whether UTF-8 encoding errors produce error conditions.
        /// </summary>
        [Input("ignoreUtf8Errors")]
        public Input<bool>? IgnoreUtf8Errors { get; set; }

        /// <summary>
        /// Specifies the identifier for the file format; must be unique for the database and schema in which the file format is created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nullIfs")]
        private InputList<string>? _nullIfs;

        /// <summary>
        /// String used to convert to and from SQL NULL.
        /// </summary>
        public InputList<string> NullIfs
        {
            get => _nullIfs ?? (_nullIfs = new InputList<string>());
            set => _nullIfs = value;
        }

        /// <summary>
        /// Boolean that specifies whether the XML parser preserves leading and trailing spaces in element content.
        /// </summary>
        [Input("preserveSpace")]
        public Input<bool>? PreserveSpace { get; set; }

        /// <summary>
        /// Specifies one or more singlebyte or multibyte characters that separate records in an input file (data loading) or unloaded file (data unloading).
        /// </summary>
        [Input("recordDelimiter")]
        public Input<string>? RecordDelimiter { get; set; }

        /// <summary>
        /// Boolean that specifies whether to replace invalid UTF-8 characters with the Unicode replacement character (�).
        /// </summary>
        [Input("replaceInvalidCharacters")]
        public Input<bool>? ReplaceInvalidCharacters { get; set; }

        /// <summary>
        /// The schema in which to create the file format.
        /// </summary>
        [Input("schema", required: true)]
        public Input<string> Schema { get; set; } = null!;

        /// <summary>
        /// Boolean that specifies to skip any blank lines encountered in the data files.
        /// </summary>
        [Input("skipBlankLines")]
        public Input<bool>? SkipBlankLines { get; set; }

        /// <summary>
        /// Boolean that specifies whether to skip the BOM (byte order mark), if present in a data file.
        /// </summary>
        [Input("skipByteOrderMark")]
        public Input<bool>? SkipByteOrderMark { get; set; }

        /// <summary>
        /// Number of lines at the start of the file to skip.
        /// </summary>
        [Input("skipHeader")]
        public Input<int>? SkipHeader { get; set; }

        /// <summary>
        /// Boolean that instructs the JSON parser to remove object fields or array elements containing null values.
        /// </summary>
        [Input("stripNullValues")]
        public Input<bool>? StripNullValues { get; set; }

        /// <summary>
        /// Boolean that instructs the JSON parser to remove outer brackets.
        /// </summary>
        [Input("stripOuterArray")]
        public Input<bool>? StripOuterArray { get; set; }

        /// <summary>
        /// Boolean that specifies whether the XML parser strips out the outer XML element, exposing 2nd level elements as separate documents.
        /// </summary>
        [Input("stripOuterElement")]
        public Input<bool>? StripOuterElement { get; set; }

        /// <summary>
        /// Defines the format of time values in the data files (data loading) or table (data unloading).
        /// </summary>
        [Input("timeFormat")]
        public Input<string>? TimeFormat { get; set; }

        /// <summary>
        /// Defines the format of timestamp values in the data files (data loading) or table (data unloading).
        /// </summary>
        [Input("timestampFormat")]
        public Input<string>? TimestampFormat { get; set; }

        /// <summary>
        /// Boolean that specifies whether to remove white space from fields.
        /// </summary>
        [Input("trimSpace")]
        public Input<bool>? TrimSpace { get; set; }

        /// <summary>
        /// Boolean that specifies whether to validate UTF-8 character encoding in string column data.
        /// </summary>
        [Input("validateUtf8")]
        public Input<bool>? ValidateUtf8 { get; set; }

        public FileFormatArgs()
        {
        }
    }

    public sealed class FileFormatState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean that specifies to allow duplicate object field names (only the last one will be preserved).
        /// </summary>
        [Input("allowDuplicate")]
        public Input<bool>? AllowDuplicate { get; set; }

        /// <summary>
        /// Boolean that specifies whether to interpret columns with no defined logical data type as UTF-8 text.
        /// </summary>
        [Input("binaryAsText")]
        public Input<bool>? BinaryAsText { get; set; }

        /// <summary>
        /// Defines the encoding format for binary input or output.
        /// </summary>
        [Input("binaryFormat")]
        public Input<string>? BinaryFormat { get; set; }

        /// <summary>
        /// Specifies a comment for the file format.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Specifies the current compression algorithm for the data file.
        /// </summary>
        [Input("compression")]
        public Input<string>? Compression { get; set; }

        /// <summary>
        /// The database in which to create the file format.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// Defines the format of date values in the data files (data loading) or table (data unloading).
        /// </summary>
        [Input("dateFormat")]
        public Input<string>? DateFormat { get; set; }

        /// <summary>
        /// Boolean that specifies whether the XML parser disables automatic conversion of numeric and Boolean values from text to native representation.
        /// </summary>
        [Input("disableAutoConvert")]
        public Input<bool>? DisableAutoConvert { get; set; }

        /// <summary>
        /// Boolean that specifies whether the XML parser disables recognition of Snowflake semi-structured data tags.
        /// </summary>
        [Input("disableSnowflakeData")]
        public Input<bool>? DisableSnowflakeData { get; set; }

        /// <summary>
        /// Specifies whether to insert SQL NULL for empty fields in an input file, which are represented by two successive delimiters.
        /// </summary>
        [Input("emptyFieldAsNull")]
        public Input<bool>? EmptyFieldAsNull { get; set; }

        /// <summary>
        /// Boolean that enables parsing of octal numbers.
        /// </summary>
        [Input("enableOctal")]
        public Input<bool>? EnableOctal { get; set; }

        /// <summary>
        /// String (constant) that specifies the character set of the source data when loading data into a table.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// Boolean that specifies whether to generate a parsing error if the number of delimited columns (i.e. fields) in an input file does not match the number of columns in the corresponding table.
        /// </summary>
        [Input("errorOnColumnCountMismatch")]
        public Input<bool>? ErrorOnColumnCountMismatch { get; set; }

        /// <summary>
        /// Single character string used as the escape character for field values.
        /// </summary>
        [Input("escape")]
        public Input<string>? Escape { get; set; }

        /// <summary>
        /// Single character string used as the escape character for unenclosed field values only.
        /// </summary>
        [Input("escapeUnenclosedField")]
        public Input<string>? EscapeUnenclosedField { get; set; }

        /// <summary>
        /// Specifies one or more singlebyte or multibyte characters that separate fields in an input file (data loading) or unloaded file (data unloading).
        /// </summary>
        [Input("fieldDelimiter")]
        public Input<string>? FieldDelimiter { get; set; }

        /// <summary>
        /// Character used to enclose strings.
        /// </summary>
        [Input("fieldOptionallyEnclosedBy")]
        public Input<string>? FieldOptionallyEnclosedBy { get; set; }

        /// <summary>
        /// Specifies the extension for files unloaded to a stage.
        /// </summary>
        [Input("fileExtension")]
        public Input<string>? FileExtension { get; set; }

        /// <summary>
        /// Specifies the format of the input files (for data loading) or output files (for data unloading).
        /// </summary>
        [Input("formatType")]
        public Input<string>? FormatType { get; set; }

        /// <summary>
        /// Boolean that specifies whether UTF-8 encoding errors produce error conditions.
        /// </summary>
        [Input("ignoreUtf8Errors")]
        public Input<bool>? IgnoreUtf8Errors { get; set; }

        /// <summary>
        /// Specifies the identifier for the file format; must be unique for the database and schema in which the file format is created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nullIfs")]
        private InputList<string>? _nullIfs;

        /// <summary>
        /// String used to convert to and from SQL NULL.
        /// </summary>
        public InputList<string> NullIfs
        {
            get => _nullIfs ?? (_nullIfs = new InputList<string>());
            set => _nullIfs = value;
        }

        /// <summary>
        /// Boolean that specifies whether the XML parser preserves leading and trailing spaces in element content.
        /// </summary>
        [Input("preserveSpace")]
        public Input<bool>? PreserveSpace { get; set; }

        /// <summary>
        /// Specifies one or more singlebyte or multibyte characters that separate records in an input file (data loading) or unloaded file (data unloading).
        /// </summary>
        [Input("recordDelimiter")]
        public Input<string>? RecordDelimiter { get; set; }

        /// <summary>
        /// Boolean that specifies whether to replace invalid UTF-8 characters with the Unicode replacement character (�).
        /// </summary>
        [Input("replaceInvalidCharacters")]
        public Input<bool>? ReplaceInvalidCharacters { get; set; }

        /// <summary>
        /// The schema in which to create the file format.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// Boolean that specifies to skip any blank lines encountered in the data files.
        /// </summary>
        [Input("skipBlankLines")]
        public Input<bool>? SkipBlankLines { get; set; }

        /// <summary>
        /// Boolean that specifies whether to skip the BOM (byte order mark), if present in a data file.
        /// </summary>
        [Input("skipByteOrderMark")]
        public Input<bool>? SkipByteOrderMark { get; set; }

        /// <summary>
        /// Number of lines at the start of the file to skip.
        /// </summary>
        [Input("skipHeader")]
        public Input<int>? SkipHeader { get; set; }

        /// <summary>
        /// Boolean that instructs the JSON parser to remove object fields or array elements containing null values.
        /// </summary>
        [Input("stripNullValues")]
        public Input<bool>? StripNullValues { get; set; }

        /// <summary>
        /// Boolean that instructs the JSON parser to remove outer brackets.
        /// </summary>
        [Input("stripOuterArray")]
        public Input<bool>? StripOuterArray { get; set; }

        /// <summary>
        /// Boolean that specifies whether the XML parser strips out the outer XML element, exposing 2nd level elements as separate documents.
        /// </summary>
        [Input("stripOuterElement")]
        public Input<bool>? StripOuterElement { get; set; }

        /// <summary>
        /// Defines the format of time values in the data files (data loading) or table (data unloading).
        /// </summary>
        [Input("timeFormat")]
        public Input<string>? TimeFormat { get; set; }

        /// <summary>
        /// Defines the format of timestamp values in the data files (data loading) or table (data unloading).
        /// </summary>
        [Input("timestampFormat")]
        public Input<string>? TimestampFormat { get; set; }

        /// <summary>
        /// Boolean that specifies whether to remove white space from fields.
        /// </summary>
        [Input("trimSpace")]
        public Input<bool>? TrimSpace { get; set; }

        /// <summary>
        /// Boolean that specifies whether to validate UTF-8 character encoding in string column data.
        /// </summary>
        [Input("validateUtf8")]
        public Input<bool>? ValidateUtf8 { get; set; }

        public FileFormatState()
        {
        }
    }
}
