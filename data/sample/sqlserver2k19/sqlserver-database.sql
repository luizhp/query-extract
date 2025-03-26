CREATE TABLE DataTypesExample (
  -- Numeric types
  TinyIntColumn TINYINT,
  SmallIntColumn SMALLINT,
  IntColumn INT,
  BigIntColumn BIGINT,
  DecimalColumn DECIMAL(10, 2),
  NumericColumn NUMERIC(10, 2),
  FloatColumn FLOAT,
  RealColumn REAL,

  -- Date and time types
  DateColumn DATE,
  TimeColumn TIME,
  DateTimeColumn DATETIME,
  SmallDateTimeColumn SMALLDATETIME,
  DateTime2Column DATETIME2,
  DateTimeOffsetColumn DATETIMEOFFSET,

  -- Character types
  CharColumn CHAR(10),
  VarCharColumn VARCHAR(50),
  TextColumn TEXT,

  -- Unicode character types
  NCharColumn NCHAR(10),
  NVarCharColumn NVARCHAR(50),
  NTextColumn NTEXT,

  -- Binary types
  BinaryColumn BINARY(10),
  VarBinaryColumn VARBINARY(50),
  ImageColumn IMAGE,

  -- Other types
  BitColumn BIT,
  UniqueIdentifierColumn UNIQUEIDENTIFIER,
  XmlColumn XML,
  JsonColumn NVARCHAR(MAX), -- JSON is stored as NVARCHAR
  GeographyColumn GEOGRAPHY,
  GeometryColumn GEOMETRY,

  -- SQL_VARIANT and hierarchyid
  SqlVariantColumn SQL_VARIANT,
  HierarchyIdColumn HIERARCHYID
);

INSERT INTO DataTypesExample (
  TinyIntColumn, SmallIntColumn, IntColumn, BigIntColumn, DecimalColumn, NumericColumn, FloatColumn, RealColumn,
  DateColumn, TimeColumn, DateTimeColumn, SmallDateTimeColumn, DateTime2Column, DateTimeOffsetColumn,
  CharColumn, VarCharColumn, TextColumn,
  NCharColumn, NVarCharColumn, NTextColumn,
  BinaryColumn, VarBinaryColumn, ImageColumn,
  BitColumn, UniqueIdentifierColumn, XmlColumn, JsonColumn, GeographyColumn, GeometryColumn,
  SqlVariantColumn, HierarchyIdColumn
)
VALUES
  (
    1, 100, 1000, 100000, 1234.56, 7890.12, 3.14159, 2.71828,
    '2023-01-01', '12:34:56', '2023-01-01 12:34:56', '2023-01-01 12:00:00', '2023-01-01 12:34:56.789', '2023-01-01 12:34:56.789 +00:00',
    'FixedText', 'VariableText1', 'Sample text data',
    N'Unicode1', N'Unicode variable text 1', N'Unicode text data',
    0x1234567890, 0xABCDEF, NULL,
    1, NEWID(), '<root><element>XML Data</element></root>', N'{"key": "value"}', geography::STGeomFromText('POINT(-122.34900 47.65100)', 4326), geometry::STGeomFromText('POINT(1 1)', 0),
    CAST('VariantData1' AS SQL_VARIANT), hierarchyid::GetRoot()
  ),
  (
    255, 32767, 2147483647, 9223372036854775807, 9876.54, 6543.21, 1.61803, 0.57721,
    '2023-06-15', '23:59:59', '2023-06-15 23:59:59', '2023-06-15 23:00:00', '2023-06-15 23:59:59.999', '2023-06-15 23:59:59.999 +05:30',
    'AnotherTex', 'VariableText2', 'Another sample text',
    N'Unicode2', N'Unicode variable text 2', N'Another unicode text',
    0x9876543210, 0x123456, NULL,
    0, NEWID(), '<root><data>Another XML</data></root>', N'{"anotherKey": "anotherValue"}', geography::STGeomFromText('POINT(-77.03653 38.89768)', 4326), geometry::STGeomFromText('LINESTRING(0 0, 1 1)', 0),
    CAST(12345 AS SQL_VARIANT), hierarchyid::Parse('/1/')
  ),
  (
    0, -32768, -2147483648, -9223372036854775808, -1234.56, -7890.12, -3.14159, -2.71828,
    '2023-12-31', '00:00:00', '2023-12-31 00:00:00', '2023-12-31 00:00:00', '2023-12-31 00:00:00.000', '2023-12-31 00:00:00.000 -08:00',
    'FinalText', 'VariableText3', 'Final sample text',
    N'Unicode3', N'Unicode variable text 3', N'Final unicode text',
    0xABCDEF1234, 0x654321, NULL,
    1, NEWID(), '<root><final>Final XML</final></root>', N'{"finalKey": "finalValue"}', geography::STGeomFromText('POINT(0 0)', 4326), geometry::STGeomFromText('POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))', 0),
    CAST(GETDATE() AS SQL_VARIANT), hierarchyid::Parse('/1/2/')
  );