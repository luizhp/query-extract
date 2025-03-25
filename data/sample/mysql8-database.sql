-- Table creation with all MySQL 8 data types

CREATE TABLE data_types_example (
  id INT AUTO_INCREMENT PRIMARY KEY, -- Integer type
  tiny_int_example TINYINT, -- Tiny integer
  small_int_example SMALLINT, -- Small integer
  medium_int_example MEDIUMINT, -- Medium integer
  big_int_example BIGINT, -- Big integer
  decimal_example DECIMAL(10, 2), -- Fixed-point number
  float_example FLOAT, -- Floating-point number
  double_example DOUBLE, -- Double-precision floating-point number
  char_example CHAR(10), -- Fixed-length string
  varchar_example VARCHAR(255), -- Variable-length string
  text_example TEXT, -- Large text
  tinytext_example TINYTEXT, -- Small text
  mediumtext_example MEDIUMTEXT, -- Medium text
  longtext_example LONGTEXT, -- Large text
  blob_example BLOB, -- Binary large object
  tinyblob_example TINYBLOB, -- Small binary object
  mediumblob_example MEDIUMBLOB, -- Medium binary object
  longblob_example LONGBLOB, -- Large binary object
  enum_example ENUM('Option1', 'Option2', 'Option3'), -- Enumeration
  set_example SET('OptionA', 'OptionB', 'OptionC'), -- Set of values
  date_example DATE, -- Date
  datetime_example DATETIME, -- Date and time
  timestamp_example TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp
  time_example TIME, -- Time
  year_example YEAR, -- Year
  json_example JSON, -- JSON data
  boolean_example BOOLEAN, -- Boolean (alias for TINYINT(1))
  binary_example BINARY(16), -- Fixed-length binary data
  varbinary_example VARBINARY(255), -- Variable-length binary data
  geometry_example GEOMETRY, -- Spatial data
  point_example POINT, -- Point spatial data
  linestring_example LINESTRING, -- Line spatial data
  polygon_example POLYGON, -- Polygon spatial data
  multipoint_example MULTIPOINT, -- Multi-point spatial data
  multilinestring_example MULTILINESTRING, -- Multi-line spatial data
  multipolygon_example MULTIPOLYGON, -- Multi-polygon spatial data
  geometrycollection_example GEOMETRYCOLLECTION -- Collection of spatial data
);

-- Sample data insertion for the table
INSERT INTO data_types_example (
  tiny_int_example, small_int_example, medium_int_example, big_int_example,
  decimal_example, float_example, double_example,
  char_example, varchar_example, text_example,
  tinytext_example, mediumtext_example, longtext_example,
  blob_example, tinyblob_example, mediumblob_example, longblob_example,
  enum_example, set_example, date_example, datetime_example,
  timestamp_example, time_example, year_example, json_example,
  boolean_example, binary_example, varbinary_example,
  geometry_example, point_example, linestring_example, polygon_example,
  multipoint_example, multilinestring_example, multipolygon_example, geometrycollection_example
)
VALUES (
  127, 32767, 8388607, 9223372036854775807,
  12345.67, 12345.67, 12345.6789,
  'FixedText', 'VariableText', 'This is a text example.',
  'TinyText', 'MediumText', 'LongText',
  'BlobData', 'TinyBlobData', 'MediumBlobData', 'LongBlobData',
  'Option1', 'OptionA,OptionB', '2025-03-24', '2025-03-24 12:34:56',
  CURRENT_TIMESTAMP, '12:34:56', 2025, '{"key": "value"}',
  TRUE, UNHEX('00112233445566778899AABBCCDDEEFF'), UNHEX('AABBCCDDEEFF'),
  NULL, ST_GeomFromText('POINT(1 1)'), ST_GeomFromText('LINESTRING(0 0, 1 1, 2 2)'),
  ST_GeomFromText('POLYGON((0 0, 1 1, 1 0, 0 0))'),
  ST_GeomFromText('MULTIPOINT((0 0), (1 1))'),
  ST_GeomFromText('MULTILINESTRING((0 0, 1 1), (2 2, 3 3))'),
  ST_GeomFromText('MULTIPOLYGON(((0 0, 1 1, 1 0, 0 0)))'),
  ST_GeomFromText('GEOMETRYCOLLECTION(POINT(1 1), LINESTRING(0 0, 1 1))')
);