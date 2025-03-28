# List

[
 { 0 TinyIntColumn          TINYINT          0x6dc0a0 0 0 0 true}
 { 1 SmallIntColumn         SMALLINT         0x6dc0a0 0 0 0 true}
 { 2 IntColumn              INT              0x6dc0a0 0 0 0 true}
 { 3 BigIntColumn           BIGINT           0x6dc0a0 0 0 0 true}
 { 4 DecimalColumn          DECIMAL          0x6d8340 0 10 2 true}
 { 5 NumericColumn          DECIMAL          0x6d8340 0 10 2 true}
 { 6 FloatColumn            FLOAT            0x6dc2a0 0 0 0 true}
 { 7 RealColumn             REAL             0x6dc2a0 0 0 0 true}
 { 8 DateColumn             DATE             0x737c80 0 0 0 true}
 { 9 TimeColumn             TIME             0x737c80 0 0 7 true}
 {10 DateTimeColumn         DATETIME         0x737c80 0 0 0 true}
 {11 SmallDateTimeColumn    SMALLDATETIME    0x737c80 0 0 0 true}
 {12 DateTime2Column        DATETIME2        0x737c80 0 0 7 true}
 {13 DateTimeOffsetColumn   DATETIMEOFFSET   0x737c80 0 0 7 true}
 {14 CharColumn             CHAR             0x6dbee0 10 0 0 true}
 {15 VarCharColumn          VARCHAR          0x6dbee0 50 0 0 true}
 {16 TextColumn             TEXT             0x6dbee0 2147483647 0 0 true}
 {17 NCharColumn            NCHAR            0x6dbee0 10 0 0 true}
 {18 NVarCharColumn         NVARCHAR         0x6dbee0 50 0 0 true}
 {19 NTextColumn            NTEXT            0x6dbee0 1073741823 0 0 true}
 {20 BinaryColumn           BINARY           0x6d8340 10 0 0 true}
 {21 VarBinaryColumn        VARBINARY        0x6d8340 50 0 0 true}
 {22 ImageColumn            IMAGE            0x6d8340 2147483647 0 0 true}
 {23 BitColumn              BIT              0x6dc2e0 0 0 0 true}
 {24 UniqueIdentifierColumn UNIQUEIDENTIFIER 0x6d8340 0 0 0 true}
 {25 XmlColumn              XML              0x6dbee0 1073741822 0 0 true}
 {26 JsonColumn             NVARCHAR         0x6dbee0 1073741822 0 0 true}
 {27 GeographyColumn        GEOGRAPHY        0x6d8340 0 0 0 true}
 {28 GeometryColumn         GEOMETRY         0x6d8340 2147483647 0 0 true}
 {29 SqlVariantColumn       SQL_VARIANT      nil      0 0 0 true}
 {30 HierarchyIdColumn      HIERARCHYID      0x6d8340 892 0 0 true}
]

# Detail
🔍 Column Name: TinyIntColumn
Column DatabaseTypeName: TINYINT
Column ScanType: int64
Column Nullable: true

🔍 Column Name: SmallIntColumn
Column DatabaseTypeName: SMALLINT
Column ScanType: int64
Column Nullable: true

🔍 Column Name: IntColumn
Column DatabaseTypeName: INT
Column ScanType: int64
Column Nullable: true

🔍 Column Name: BigIntColumn
Column DatabaseTypeName: BIGINT
Column ScanType: int64
Column Nullable: true

🔍 Column Name: DecimalColumn
Column DatabaseTypeName: DECIMAL
Column has precision 10 and scale 2
Column ScanType: []uint8
Column Nullable: true

🔍 Column Name: NumericColumn
Column DatabaseTypeName: DECIMAL
Column has precision 10 and scale 2
Column ScanType: []uint8
Column Nullable: true

🔍 Column Name: FloatColumn
Column DatabaseTypeName: FLOAT
Column ScanType: float64
Column Nullable: true

🔍 Column Name: RealColumn
Column DatabaseTypeName: REAL
Column ScanType: float64
Column Nullable: true

🔍 Column Name: DateColumn
Column DatabaseTypeName: DATE
Column ScanType: time.Time
Column Nullable: true

🔍 Column Name: TimeColumn
Column DatabaseTypeName: TIME
Column has precision 0 and scale 7
Column ScanType: time.Time
Column Nullable: true

🔍 Column Name: DateTimeColumn
Column DatabaseTypeName: DATETIME
Column ScanType: time.Time
Column Nullable: true

🔍 Column Name: SmallDateTimeColumn
Column DatabaseTypeName: SMALLDATETIME
Column ScanType: time.Time
Column Nullable: true

🔍 Column Name: DateTime2Column
Column DatabaseTypeName: DATETIME2
Column has precision 0 and scale 7
Column ScanType: time.Time
Column Nullable: true

🔍 Column Name: DateTimeOffsetColumn
Column DatabaseTypeName: DATETIMEOFFSET
Column has precision 0 and scale 7
Column ScanType: time.Time
Column Nullable: true

🔍 Column Name: CharColumn
Column DatabaseTypeName: CHAR
Column has length 10
Column ScanType: string
Column Nullable: true

🔍 Column Name: VarCharColumn
Column DatabaseTypeName: VARCHAR
Column has length 50
Column ScanType: string
Column Nullable: true

🔍 Column Name: TextColumn
Column DatabaseTypeName: TEXT
Column has length 2147483647
Column ScanType: string
Column Nullable: true

🔍 Column Name: NCharColumn
Column DatabaseTypeName: NCHAR
Column has length 10
Column ScanType: string
Column Nullable: true

🔍 Column Name: NVarCharColumn
Column DatabaseTypeName: NVARCHAR
Column has length 50
Column ScanType: string
Column Nullable: true

🔍 Column Name: NTextColumn
Column DatabaseTypeName: NTEXT
Column has length 1073741823
Column ScanType: string
Column Nullable: true

🔍 Column Name: BinaryColumn
Column DatabaseTypeName: BINARY
Column has length 10
Column ScanType: []uint8
Column Nullable: true

🔍 Column Name: VarBinaryColumn
Column DatabaseTypeName: VARBINARY
Column has length 50
Column ScanType: []uint8
Column Nullable: true

🔍 Column Name: ImageColumn
Column DatabaseTypeName: IMAGE
Column has length 2147483647
Column ScanType: []uint8
Column Nullable: true

🔍 Column Name: BitColumn
Column DatabaseTypeName: BIT
Column ScanType: bool
Column Nullable: true

🔍 Column Name: UniqueIdentifierColumn
Column DatabaseTypeName: UNIQUEIDENTIFIER
Column ScanType: []uint8
Column Nullable: true

🔍 Column Name: XmlColumn
Column DatabaseTypeName: XML
Column has length 1073741822
Column ScanType: string
Column Nullable: true

🔍 Column Name: JsonColumn
Column DatabaseTypeName: NVARCHAR
Column has length 1073741822
Column ScanType: string
Column Nullable: true

🔍 Column Name: GeographyColumn
Column DatabaseTypeName: GEOGRAPHY
Column ScanType: []uint8
Column Nullable: true

🔍 Column Name: GeometryColumn
Column DatabaseTypeName: GEOMETRY
Column has length 2147483647
Column ScanType: []uint8
Column Nullable: true

🔍 Column Name: SqlVariantColumn
Column DatabaseTypeName: SQL_VARIANT
Column ScanType: %!s(<nil>)
Column Nullable: true

🔍 Column Name: HierarchyIdColumn
Column DatabaseTypeName: HIERARCHYID
Column has length 892
Column ScanType: []uint8
Column Nullable: true


## notes


BIGINT
BINARY
BIT
CHAR
DATE
DATETIME
DATETIME2
DATETIMEOFFSET
DECIMAL
FLOAT
GEOGRAPHY
GEOMETRY
HIERARCHYID
IMAGE
INT
NCHAR
NTEXT
NVARCHAR
REAL
SMALLDATETIME
SMALLINT
SQL_VARIANT
TEXT
TIME
TINYINT
UNIQUEIDENTIFIER
VARBINARY
VARCHAR
XML


[]uint8
%!s(<nil>)
bool
float64
int64
string
time.Time



Column Name Column DatabaseTypeName Column ScanType Column Nullable

TinyIntColumn
@TINYINT
@@@int64
@true

SmallIntColumn
@SMALLINT
@@@int64
@true

IntColumn
@INT
@@@int64
@true

BigIntColumn
@BIGINT
@@@int64
@true

DecimalColumn
@DECIMAL
@10@2
@@@[]uint8
@true

NumericColumn
@DECIMAL
@10@2
@@@[]uint8
@true

FloatColumn
@FLOAT
@@@float64
@true

RealColumn
@REAL
@@@float64
@true

DateColumn
@DATE
@@@time.Time
@true

TimeColumn
@TIME
@0@7
@@@time.Time
@true

DateTimeColumn
@DATETIME
@@@time.Time
@true

SmallDateTimeColumn
@SMALLDATETIME
@@@time.Time
@true

DateTime2Column
@DATETIME2
@0@7
@@@time.Time
@true

DateTimeOffsetColumn
@DATETIMEOFFSET
@0@7
@@@time.Time
@true

CharColumn
@CHAR
@10
@@@string
@true

VarCharColumn
@VARCHAR
@50
@@@string
@true

TextColumn
@TEXT
@2147483647
@@@string
@true

NCharColumn
@NCHAR
@10
@@@string
@true

NVarCharColumn
@NVARCHAR
@50
@@@string
@true

NTextColumn
@NTEXT
@1073741823
@@@string
@true

BinaryColumn
@BINARY
@10
@@@[]uint8
@true

VarBinaryColumn
@VARBINARY
@50
@@@[]uint8
@true

ImageColumn
@IMAGE
@2147483647
@@@[]uint8
@true

BitColumn
@BIT
@@@bool
@true

UniqueIdentifierColumn
@UNIQUEIDENTIFIER
@@@[]uint8
@true

XmlColumn
@XML
@1073741822
@@@string
@true

JsonColumn
@NVARCHAR
@1073741822
@@@string
@true

GeographyColumn
@GEOGRAPHY
@@@[]uint8
@true

GeometryColumn
@GEOMETRY
@2147483647
@@@[]uint8
@true

SqlVariantColumn
@SQL_VARIANT
@@@%!s(<nil>)
@true

HierarchyIdColumn
@HIERARCHYID
@892
@@@[]uint8
@true

