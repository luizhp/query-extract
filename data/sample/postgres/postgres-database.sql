-- PostgreSQL Table Creation Script (Revised - No Geography, Geometry, or Ltree)

CREATE TABLE DataTypesExample (
    -- Numeric types
    TinyIntColumn SMALLINT, -- TINYINT in SQL Server maps to SMALLINT in PostgreSQL (0-255)
    SmallIntColumn SMALLINT,
    IntColumn INTEGER,
    BigIntColumn BIGINT,
    DecimalColumn NUMERIC(10, 2),
    NumericColumn NUMERIC(10, 2),
    FloatColumn DOUBLE PRECISION,
    RealColumn REAL,

    -- Date and time types
    DateColumn DATE,
    TimeColumn TIME,
    DateTimeColumn TIMESTAMP,
    SmallDateTimeColumn TIMESTAMP, -- SMALLDATETIME in SQL Server maps to TIMESTAMP in PostgreSQL
    DateTime2Column TIMESTAMP,
    DateTimeOffsetColumn TIMESTAMPTZ,

    -- Character types
    CharColumn CHAR(10),
    VarCharColumn VARCHAR(50),
    TextColumn TEXT,

    -- Unicode character types (PostgreSQL handles Unicode natively)
    NCharColumn CHAR(10),
    NVarCharColumn VARCHAR(50),
    NTextColumn TEXT,

    -- Binary types
    BinaryColumn BYTEA,
    VarBinaryColumn BYTEA,
    ImageColumn BYTEA,

    -- Other types
    BitColumn BOOLEAN,
    UniqueIdentifierColumn UUID,
    XmlColumn XML,
    JsonColumn JSONB,

    -- SQL_VARIANT equivalent (using TEXT for simplicity, consider JSONB for structured data)
    SqlVariantColumn TEXT
);

-- PostgreSQL Data Insertion Script (10 rows - Revised)

INSERT INTO DataTypesExample (
    TinyIntColumn, SmallIntColumn, IntColumn, BigIntColumn, DecimalColumn, NumericColumn, FloatColumn, RealColumn,
    DateColumn, TimeColumn, DateTimeColumn, SmallDateTimeColumn, DateTime2Column, DateTimeOffsetColumn,
    CharColumn, VarCharColumn, TextColumn,
    NCharColumn, NVarCharColumn, NTextColumn,
    BinaryColumn, VarBinaryColumn, ImageColumn,
    BitColumn, UniqueIdentifierColumn, XmlColumn, JsonColumn,
    SqlVariantColumn
) VALUES
(
    1, 100, 1000, 100000, 1234.56, 7890.12, 3.14159, 2.71828,
    '2023-01-01', '12:34:56', '2023-01-01 12:34:56', '2023-01-01 12:00:00', '2023-01-01 12:34:56.789', '2023-01-01 12:34:56.789 +00:00',
    'FixedText', 'VariableText1', 'Sample text data',
    'Unicode1', 'Unicode variable text 1', 'Unicode text data',
    '\x1234567890', '\xABCDEF', NULL,
    true, 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', '<root><element>XML Data</element></root>', '{"key": "value"}',
    'VariantData1'
),
(
    255, 32767, 2147483647, 9223372036854775807, 9876.54, 6543.21, 1.61803, 0.57721,
    '2023-06-15', '23:59:59', '2023-06-15 23:59:59', '2023-06-15 23:00:00', '2023-06-15 23:59:59.999', '2023-06-15 23:59:59.999 +05:30',
    'AnotherTex', 'VariableText2', 'Another sample text',
    'Unicode2', 'Unicode variable text 2', 'Another unicode text',
    '\x9876543210', '\x123456', NULL,
    false, 'f47ac10b-58cc-4372-a567-0e02b2c3d479', '<root><data>Another XML</data></root>', '{"anotherKey": "anotherValue"}',
    '12345'
),
(
    0, -32768, -2147483648, -9223372036854775808, -1234.56, -7890.12, -3.14159, -2.71828,
    '2023-12-31', '00:00:00', '2023-12-31 00:00:00', '2023-12-31 00:00:00', '2023-12-31 00:00:00.000', '2023-12-31 00:00:00.000 -08:00',
    'FinalText', 'VariableText3', 'Final sample text',
    'Unicode3', 'Unicode variable text 3', 'Final unicode text',
    '\xABCDEF1234', '\x654321', NULL,
    true, '550e8400-e29b-41d4-a716-446655440000', '<root><final>Final XML</final></root>', '{"finalKey": "finalValue"}',
    '2024-01-20 10:00:00'
),
(
    128, 16384, 1073741824, 4611686018427387903, 5432.10, 1098.76, 2.71828, 1.61803,
    '2024-03-10', '10:15:30', '2024-03-10 10:15:30', '2024-03-10 10:00:00', '2024-03-10 10:15:30.123', '2024-03-10 10:15:30.123 +02:00',
    'MidText', 'VariableText4', 'Middle sample text',
    'Unicode4', 'Unicode variable text 4', 'Middle unicode text',
    '\x0011223344', '\x556677', NULL,
    false, '8f4b0a4e-0f9a-4a8c-b5a2-e7a9c1d2f3b4', '<root><middle>Middle XML</middle></root>', '{"middleKey": "middleValue"}',
    'AnotherVariant'
),
(
    200, -16384, -1073741824, -4611686018427387904, -5432.10, -1098.76, -2.71828, -1.61803,
    '2024-09-22', '18:45:00', '2024-09-22 18:45:00', '2024-09-22 18:00:00', '2024-09-22 18:45:00.456', '2024-09-22 18:45:00.456 -05:00',
    'LateText', 'VariableText5', 'Late sample text',
    'Unicode5', 'Unicode variable text 5', 'Late unicode text',
    '\x99AABBCCDD', '\xEEFF00', NULL,
    true, 'c7a9d3e1-b8f0-4c5a-9d2b-f6e8a1c0d3b4', '<root><late>Late XML</late></root>', '{"lateKey": "lateValue"}',
    'YetAnotherVariant'
),
(
    10, 500, 5000, 500000, 5555.55, 6666.66, 1.23456, 6.54321,
    '2023-05-15', '08:00:00', '2023-05-15 08:00:00', '2023-05-15 08:00:00', '2023-05-15 08:00:00.123', '2023-05-15 08:00:00.123 +03:00',
    'Test1', 'VariableTest1', 'Test text 1',
    'TestU1', 'TestUVar1', 'TestU Text 1',
    '\x1122334455', '\x667788', NULL,
    false, '123e4567-e89b-12d3-a456-426614174000', '<root><test>Test XML 1</test></root>', '{"testKey1": "testValue1"}',
    'TestVariant1'
),
(
    20, 1000, 10000, 1000000, 1111.11, 2222.22, 9.87654, 4.32109,
    '2023-07-20', '15:30:00', '2023-07-20 15:30:00', '2023-07-20 15:00:00', '2023-07-20 15:30:00.456', '2023-07-20 15:30:00.456 -02:00',
    'Test2', 'VariableTest2', 'Test text 2',
    'TestU2', 'TestUVar2', 'TestU Text 2',
    '\x99AABBCCDD', '\xEEFF00', NULL,
    true, '987f6543-21ba-dc09-87fe-426614174000', '<root><test>Test XML 2</test></root>', '{"testKey2": "testValue2"}',
    'TestVariant2'
),
(
    30, 1500, 15000, 1500000, 3333.33, 4444.44, 0.12345, 5.43210,
    '2023-09-25', '22:45:00', '2023-09-25 22:45:00', '2023-09-25 22:00:00', '2023-09-25 22:45:00.789', '2023-09-25 22:45:00.789 +01:00',
    'Test3', 'VariableTest3', 'Test text 3',
    'TestU3', 'TestUVar3', 'TestU Text 3',
    '\x1234567890', '\xABCDEF', NULL,
    false, 'abcdef12-3456-7890-abcd-ef1234567890', '<root><test>Test XML 3</test></root>', '{"testKey3": "testValue3"}',
    'TestVariant3'
),
(
    40, 2000, 20000, 2000000, 7777.77, 8888.88, 6.78901, 1.09876,
    '2023-11-30', '05:15:00', '2023-11-30 05:15:00', '2023-11-30 05:00:00', '2023-11-30 05:15:00.000', '2023-11-30 05:15:00.000 -04:00',
    'Test4', 'VariableTest4', 'Test text 4',
    'TestU4', 'TestUVar4', 'TestU Text 4',
    '\x0011223344', '\x556677', NULL,
    true, 'fedcba98-7654-3210-fedc-ba9876543210', '<root><test>Test XML 4</test></root>', '{"testKey4": "testValue4"}',
    'TestVariant4'
),
(
    50, 2500, 25000, 2500000, 9999.99, 10101.01, 3.45678, 8.76543,
    '2024-02-05', '12:00:00', '2024-02-05 12:00:00', '2024-02-05 12:00:00', '2024-02-05 12:00:00.321', '2024-02-05 12:00:00.321 +04:00',
    'Test5', 'VariableTest5', 'Test text 5',
    'TestU5', 'TestUVar5', 'TestU Text 5',
    '\x99AABBCCDD', '\xEEFF00', NULL,
    false, '1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d', '<root><test>Test XML 5</test></root>', '{"testKey5": "testValue5"}',
    'TestVariant5'
);
