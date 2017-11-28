//line parser.y:2
package main

import __yyfmt__ "fmt"

//line parser.y:2
import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type node struct {
	name       string
	children   []node
	attributes map[string]string
}

func (n node) String() string {
	buf := new(bytes.Buffer)
	n.print(buf, " ")
	return buf.String()
}

func (n node) print(out io.Writer, indent string) {
	if len(n.attributes) > 0 {
		fmt.Fprintf(out, "\n%v%v %s", indent, n.name, n.attributes)
	} else {
		fmt.Fprintf(out, "\n%v%v", indent, n.name)
	}
	for _, nn := range n.children {
		nn.print(out, indent+"  ")
	}
}

func Node(name string) node {
	return node{name: name, attributes: make(map[string]string)}
}

func (n node) append(nn ...node) node {
	n.children = append(n.children, nn...)
	return n
}

func (n node) attribute(key string, value string) node {
	n.attributes[key] = value
	return n
}

//line parser.y:50
type yySymType struct {
	yys   int
	node  node
	token string
	value string
}

const T_INCLUDE = 57346
const T_INCLUDE_ONCE = 57347
const T_EVAL = 57348
const T_REQUIRE = 57349
const T_REQUIRE_ONCE = 57350
const T_LOGICAL_OR = 57351
const T_LOGICAL_XOR = 57352
const T_LOGICAL_AND = 57353
const T_PRINT = 57354
const T_YIELD = 57355
const T_DOUBLE_ARROW = 57356
const T_YIELD_FROM = 57357
const T_PLUS_EQUAL = 57358
const T_MINUS_EQUAL = 57359
const T_MUL_EQUAL = 57360
const T_DIV_EQUAL = 57361
const T_CONCAT_EQUAL = 57362
const T_MOD_EQUAL = 57363
const T_AND_EQUAL = 57364
const T_OR_EQUAL = 57365
const T_XOR_EQUAL = 57366
const T_SL_EQUAL = 57367
const T_SR_EQUAL = 57368
const T_POW_EQUAL = 57369
const T_COALESCE = 57370
const T_BOOLEAN_OR = 57371
const T_BOOLEAN_AND = 57372
const T_IS_EQUAL = 57373
const T_IS_NOT_EQUAL = 57374
const T_IS_IDENTICAL = 57375
const T_IS_NOT_IDENTICAL = 57376
const T_SPACESHIP = 57377
const T_IS_SMALLER_OR_EQUAL = 57378
const T_IS_GREATER_OR_EQUAL = 57379
const T_SL = 57380
const T_SR = 57381
const T_INSTANCEOF = 57382
const T_INC = 57383
const T_DEC = 57384
const T_INT_CAST = 57385
const T_DOUBLE_CAST = 57386
const T_STRING_CAST = 57387
const T_ARRAY_CAST = 57388
const T_OBJECT_CAST = 57389
const T_BOOL_CAST = 57390
const T_UNSET_CAST = 57391
const T_POW = 57392
const T_NEW = 57393
const T_CLONE = 57394
const T_ELSEIF = 57395
const T_ELSE = 57396
const T_ENDIF = 57397
const T_STATIC = 57398
const T_ABSTRACT = 57399
const T_FINAL = 57400
const T_PRIVATE = 57401
const T_PROTECTED = 57402
const T_PUBLIC = 57403
const T_EXIT = 57404
const T_IF = 57405
const T_LNUMBER = 57406
const T_DNUMBER = 57407
const T_STRING = 57408
const T_STRING_VARNAME = 57409
const T_VARIABLE = 57410
const T_NUM_STRING = 57411
const T_INLINE_HTML = 57412
const T_CHARACTER = 57413
const T_BAD_CHARACTER = 57414
const T_ENCAPSED_AND_WHITESPACE = 57415
const T_CONSTANT_ENCAPSED_STRING = 57416
const T_ECHO = 57417
const T_DO = 57418
const T_WHILE = 57419
const T_ENDWHILE = 57420
const T_FOR = 57421
const T_ENDFOR = 57422
const T_FOREACH = 57423
const T_ENDFOREACH = 57424
const T_DECLARE = 57425
const T_ENDDECLARE = 57426
const T_AS = 57427
const T_SWITCH = 57428
const T_ENDSWITCH = 57429
const T_CASE = 57430
const T_DEFAULT = 57431
const T_BREAK = 57432
const T_CONTINUE = 57433
const T_GOTO = 57434
const T_FUNCTION = 57435
const T_CONST = 57436
const T_RETURN = 57437
const T_TRY = 57438
const T_CATCH = 57439
const T_FINALLY = 57440
const T_THROW = 57441
const T_USE = 57442
const T_INSTEADOF = 57443
const T_GLOBAL = 57444
const T_VAR = 57445
const T_UNSET = 57446
const T_ISSET = 57447
const T_EMPTY = 57448
const T_HALT_COMPILER = 57449
const T_CLASS = 57450
const T_TRAIT = 57451
const T_INTERFACE = 57452
const T_EXTENDS = 57453
const T_IMPLEMENTS = 57454
const T_OBJECT_OPERATOR = 57455
const T_LIST = 57456
const T_ARRAY = 57457
const T_CALLABLE = 57458
const T_CLASS_C = 57459
const T_TRAIT_C = 57460
const T_METHOD_C = 57461
const T_FUNC_C = 57462
const T_LINE = 57463
const T_FILE = 57464
const T_COMMENT = 57465
const T_DOC_COMMENT = 57466
const T_OPEN_TAG = 57467
const T_OPEN_TAG_WITH_ECHO = 57468
const T_CLOSE_TAG = 57469
const T_WHITESPACE = 57470
const T_START_HEREDOC = 57471
const T_END_HEREDOC = 57472
const T_DOLLAR_OPEN_CURLY_BRACES = 57473
const T_CURLY_OPEN = 57474
const T_PAAMAYIM_NEKUDOTAYIM = 57475
const T_NAMESPACE = 57476
const T_NS_C = 57477
const T_DIR = 57478
const T_NS_SEPARATOR = 57479
const T_ELLIPSIS = 57480

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"T_INCLUDE",
	"T_INCLUDE_ONCE",
	"T_EVAL",
	"T_REQUIRE",
	"T_REQUIRE_ONCE",
	"','",
	"T_LOGICAL_OR",
	"T_LOGICAL_XOR",
	"T_LOGICAL_AND",
	"T_PRINT",
	"T_YIELD",
	"T_DOUBLE_ARROW",
	"T_YIELD_FROM",
	"'='",
	"T_PLUS_EQUAL",
	"T_MINUS_EQUAL",
	"T_MUL_EQUAL",
	"T_DIV_EQUAL",
	"T_CONCAT_EQUAL",
	"T_MOD_EQUAL",
	"T_AND_EQUAL",
	"T_OR_EQUAL",
	"T_XOR_EQUAL",
	"T_SL_EQUAL",
	"T_SR_EQUAL",
	"T_POW_EQUAL",
	"'?'",
	"':'",
	"T_COALESCE",
	"T_BOOLEAN_OR",
	"T_BOOLEAN_AND",
	"'|'",
	"'^'",
	"'&'",
	"T_IS_EQUAL",
	"T_IS_NOT_EQUAL",
	"T_IS_IDENTICAL",
	"T_IS_NOT_IDENTICAL",
	"T_SPACESHIP",
	"'<'",
	"T_IS_SMALLER_OR_EQUAL",
	"'>'",
	"T_IS_GREATER_OR_EQUAL",
	"T_SL",
	"T_SR",
	"'+'",
	"'-'",
	"'.'",
	"'*'",
	"'/'",
	"'%'",
	"'!'",
	"T_INSTANCEOF",
	"'~'",
	"T_INC",
	"T_DEC",
	"T_INT_CAST",
	"T_DOUBLE_CAST",
	"T_STRING_CAST",
	"T_ARRAY_CAST",
	"T_OBJECT_CAST",
	"T_BOOL_CAST",
	"T_UNSET_CAST",
	"'@'",
	"T_POW",
	"'['",
	"T_NEW",
	"T_CLONE",
	"T_ELSEIF",
	"T_ELSE",
	"T_ENDIF",
	"T_STATIC",
	"T_ABSTRACT",
	"T_FINAL",
	"T_PRIVATE",
	"T_PROTECTED",
	"T_PUBLIC",
	"T_EXIT",
	"T_IF",
	"T_LNUMBER",
	"T_DNUMBER",
	"T_STRING",
	"T_STRING_VARNAME",
	"T_VARIABLE",
	"T_NUM_STRING",
	"T_INLINE_HTML",
	"T_CHARACTER",
	"T_BAD_CHARACTER",
	"T_ENCAPSED_AND_WHITESPACE",
	"T_CONSTANT_ENCAPSED_STRING",
	"T_ECHO",
	"T_DO",
	"T_WHILE",
	"T_ENDWHILE",
	"T_FOR",
	"T_ENDFOR",
	"T_FOREACH",
	"T_ENDFOREACH",
	"T_DECLARE",
	"T_ENDDECLARE",
	"T_AS",
	"T_SWITCH",
	"T_ENDSWITCH",
	"T_CASE",
	"T_DEFAULT",
	"T_BREAK",
	"T_CONTINUE",
	"T_GOTO",
	"T_FUNCTION",
	"T_CONST",
	"T_RETURN",
	"T_TRY",
	"T_CATCH",
	"T_FINALLY",
	"T_THROW",
	"T_USE",
	"T_INSTEADOF",
	"T_GLOBAL",
	"T_VAR",
	"T_UNSET",
	"T_ISSET",
	"T_EMPTY",
	"T_HALT_COMPILER",
	"T_CLASS",
	"T_TRAIT",
	"T_INTERFACE",
	"T_EXTENDS",
	"T_IMPLEMENTS",
	"T_OBJECT_OPERATOR",
	"T_LIST",
	"T_ARRAY",
	"T_CALLABLE",
	"T_CLASS_C",
	"T_TRAIT_C",
	"T_METHOD_C",
	"T_FUNC_C",
	"T_LINE",
	"T_FILE",
	"T_COMMENT",
	"T_DOC_COMMENT",
	"T_OPEN_TAG",
	"T_OPEN_TAG_WITH_ECHO",
	"T_CLOSE_TAG",
	"T_WHITESPACE",
	"T_START_HEREDOC",
	"T_END_HEREDOC",
	"T_DOLLAR_OPEN_CURLY_BRACES",
	"T_CURLY_OPEN",
	"T_PAAMAYIM_NEKUDOTAYIM",
	"T_NAMESPACE",
	"T_NS_C",
	"T_DIR",
	"T_NS_SEPARATOR",
	"T_ELLIPSIS",
	"';'",
	"'{'",
	"'}'",
	"'('",
	"')'",
	"'$'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:461
const src = `<?
namespace foo\bar\test;

+$b++;

do {
    function test(string $a, \bar\baz $b = $t) { 
        yield $a => $b;
    }
} while($a = $b = $c);

`

func main() {
	yyDebug = 0
	yyErrorVerbose = true
	l := newLexer(bytes.NewBufferString(src), os.Stdout, "file.name")
	yyParse(l)
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 67,
	158, 46,
	-2, 197,
	-1, 233,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 169,
	-1, 234,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 170,
	-1, 235,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 171,
	-1, 236,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 172,
	-1, 237,
	38, 0,
	39, 0,
	40, 0,
	41, 0,
	42, 0,
	-2, 173,
	-1, 238,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 174,
	-1, 239,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 175,
	-1, 240,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 176,
	-1, 241,
	43, 0,
	44, 0,
	45, 0,
	46, 0,
	-2, 177,
	-1, 271,
	162, 107,
	-2, 112,
}

const yyPrivate = 57344

const yyLast = 1543

var yyAct = [...]int{

	11, 129, 285, 291, 283, 286, 296, 49, 271, 268,
	191, 126, 189, 294, 273, 310, 272, 247, 303, 46,
	181, 208, 46, 184, 185, 186, 187, 188, 207, 190,
	308, 192, 193, 194, 195, 196, 197, 198, 199, 200,
	201, 202, 203, 204, 302, 209, 215, 128, 312, 135,
	137, 136, 43, 44, 128, 267, 246, 245, 287, 165,
	181, 145, 49, 190, 45, 192, 193, 202, 203, 159,
	299, 160, 133, 134, 138, 140, 139, 152, 153, 150,
	151, 154, 155, 156, 157, 158, 148, 149, 142, 143,
	141, 144, 146, 147, 205, 47, 289, 290, 47, 162,
	144, 146, 147, 163, 305, 15, 314, 145, 17, 297,
	51, 52, 206, 128, 212, 293, 145, 4, 292, 1,
	213, 214, 164, 5, 8, 130, 182, 183, 18, 16,
	282, 281, 49, 304, 216, 217, 218, 219, 220, 221,
	222, 223, 224, 225, 226, 227, 228, 229, 230, 231,
	232, 233, 234, 235, 236, 237, 238, 239, 240, 241,
	242, 244, 289, 290, 284, 13, 210, 248, 250, 251,
	252, 253, 254, 255, 256, 257, 258, 259, 260, 261,
	2, 293, 288, 127, 292, 3, 48, 161, 307, 298,
	263, 0, 264, 0, 0, 0, 0, 0, 135, 137,
	136, 295, 0, 0, 0, 0, 266, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 159, 0,
	160, 133, 134, 138, 140, 139, 152, 153, 150, 151,
	154, 155, 156, 157, 158, 148, 149, 142, 143, 141,
	144, 146, 147, 0, 270, 0, 0, 0, 0, 0,
	274, 0, 0, 0, 0, 0, 145, 142, 143, 141,
	144, 146, 147, 0, 0, 0, 277, 0, 0, 279,
	280, 61, 62, 63, 64, 65, 145, 68, 69, 70,
	66, 67, 0, 42, 148, 149, 142, 143, 141, 144,
	146, 147, 0, 300, 0, 0, 301, 0, 0, 0,
	0, 0, 306, 0, 0, 145, 309, 0, 311, 0,
	0, 0, 313, 0, 0, 316, 22, 23, 0, 0,
	0, 0, 24, 71, 25, 20, 21, 32, 33, 34,
	35, 36, 37, 38, 39, 0, 0, 72, 59, 75,
	76, 77, 53, 54, 55, 56, 57, 58, 73, 74,
	276, 0, 50, 0, 46, 0, 0, 0, 0, 0,
	0, 78, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 105, 106, 107, 108, 109, 99, 100, 101,
	102, 103, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 60, 0, 117, 115, 116, 112, 113, 0,
	104, 110, 111, 118, 119, 121, 120, 122, 123, 0,
	0, 0, 0, 0, 131, 28, 29, 30, 31, 0,
	114, 125, 124, 40, 41, 0, 42, 0, 26, 0,
	47, 0, 134, 138, 140, 139, 152, 153, 150, 151,
	154, 155, 156, 157, 158, 148, 149, 142, 143, 141,
	144, 146, 147, 0, 0, 0, 0, 0, 0, 22,
	23, 0, 0, 0, 0, 24, 145, 25, 20, 21,
	32, 33, 34, 35, 36, 37, 38, 39, 0, 0,
	0, 19, 0, 0, 0, 0, 43, 44, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 46, 0, 0,
	0, 0, 0, 0, 0, 10, 0, 0, 0, 155,
	156, 157, 158, 148, 149, 142, 143, 141, 144, 146,
	147, 0, 12, 135, 137, 136, 0, 0, 0, 131,
	28, 29, 30, 31, 145, 27, 0, 14, 40, 41,
	0, 42, 0, 159, 0, 160, 133, 134, 138, 140,
	139, 152, 153, 150, 151, 154, 155, 156, 157, 158,
	148, 149, 142, 143, 141, 144, 146, 147, 0, 9,
	315, 26, 0, 47, 22, 23, 0, 0, 0, 0,
	24, 145, 25, 20, 21, 32, 33, 34, 35, 36,
	37, 38, 39, 0, 0, 0, 19, 0, 0, 0,
	0, 43, 44, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 46, 0, 0, 0, 0, 0, 0, 0,
	10, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 12, 0, 0,
	0, 0, 0, 0, 6, 28, 29, 30, 31, 0,
	27, 0, 14, 40, 41, 0, 42, 152, 153, 150,
	151, 154, 155, 156, 157, 158, 148, 149, 142, 143,
	141, 144, 146, 147, 0, 275, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 211, 26, 145, 47, 22,
	23, 0, 0, 0, 0, 24, 0, 25, 20, 21,
	32, 33, 34, 35, 36, 37, 38, 39, 0, 0,
	0, 19, 0, 0, 0, 0, 43, 44, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 46, 131, 28,
	29, 30, 31, 0, 0, 10, 0, 40, 41, 0,
	42, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 12, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 27, 0, 14, 0, 0,
	0, 0, 0, 22, 23, 0, 0, 0, 0, 24,
	0, 25, 20, 21, 32, 33, 34, 35, 36, 37,
	38, 39, 0, 7, 0, 19, 0, 0, 0, 9,
	0, 26, 0, 47, 135, 137, 136, 0, 0, 0,
	0, 46, 0, 0, 0, 0, 0, 0, 0, 10,
	0, 0, 0, 0, 159, 0, 160, 133, 134, 138,
	140, 139, 152, 153, 150, 151, 154, 155, 156, 157,
	158, 148, 149, 142, 143, 141, 144, 146, 147, 27,
	0, 131, 28, 29, 30, 31, 0, 0, 0, 0,
	40, 41, 145, 42, 139, 152, 153, 150, 151, 154,
	155, 156, 157, 158, 148, 149, 142, 143, 141, 144,
	146, 147, 0, 9, 249, 26, 0, 47, 0, 0,
	0, 0, 0, 0, 0, 145, 22, 23, 0, 0,
	0, 0, 24, 0, 25, 20, 21, 32, 33, 34,
	35, 36, 37, 38, 39, 0, 0, 0, 19, 0,
	131, 28, 29, 30, 31, 0, 0, 0, 0, 40,
	41, 0, 42, 0, 46, 0, 0, 0, 0, 131,
	28, 29, 30, 31, 0, 0, 0, 243, 40, 41,
	0, 42, 0, 0, 0, 0, 262, 0, 0, 0,
	0, 0, 0, 0, 0, 22, 23, 0, 0, 0,
	0, 24, 27, 25, 20, 21, 32, 33, 34, 35,
	36, 37, 38, 39, 22, 23, 0, 19, 0, 0,
	24, 0, 25, 20, 21, 32, 33, 34, 35, 36,
	37, 38, 39, 46, 0, 0, 19, 0, 26, 0,
	47, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 46, 166, 167, 168, 169, 171, 172, 173,
	174, 175, 176, 177, 178, 170, 0, 0, 0, 0,
	0, 27, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	27, 0, 0, 0, 179, 180, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 26, 0, 47,
	135, 137, 136, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 26, 0, 47, 0,
	159, 0, 160, 133, 134, 138, 140, 139, 152, 153,
	150, 151, 154, 155, 156, 157, 158, 148, 149, 142,
	143, 141, 144, 146, 147, 135, 137, 136, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 145, 0,
	0, 0, 0, 0, 0, 159, 0, 160, 133, 134,
	138, 140, 139, 152, 153, 150, 151, 154, 155, 156,
	157, 158, 148, 149, 142, 143, 141, 144, 146, 147,
	135, 137, 136, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 145, 0, 0, 0, 0, 0, 0,
	159, 269, 160, 133, 134, 138, 140, 139, 152, 153,
	150, 151, 154, 155, 156, 157, 158, 148, 149, 142,
	143, 141, 144, 146, 147, 0, 135, 137, 136, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 145, 0,
	278, 0, 0, 0, 0, 0, 159, 0, 160, 133,
	134, 138, 140, 139, 152, 153, 150, 151, 154, 155,
	156, 157, 158, 148, 149, 142, 143, 141, 144, 146,
	147, 137, 136, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 132, 145, 0, 0, 0, 0, 0,
	159, 0, 160, 133, 134, 138, 140, 139, 152, 153,
	150, 151, 154, 155, 156, 157, 158, 148, 149, 142,
	143, 141, 144, 146, 147, 136, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 145, 0,
	0, 0, 0, 159, 0, 160, 133, 134, 138, 140,
	139, 152, 153, 150, 151, 154, 155, 156, 157, 158,
	148, 149, 142, 143, 141, 144, 146, 147, 265, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 145, 0, 159, 0, 160, 133, 134, 138, 140,
	139, 152, 153, 150, 151, 154, 155, 156, 157, 158,
	148, 149, 142, 143, 141, 144, 146, 147, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	159, 145, 160, 133, 134, 138, 140, 139, 152, 153,
	150, 151, 154, 155, 156, 157, 158, 148, 149, 142,
	143, 141, 144, 146, 147, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 145, 160,
	133, 134, 138, 140, 139, 152, 153, 150, 151, 154,
	155, 156, 157, 158, 148, 149, 142, 143, 141, 144,
	146, 147, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 145, 138, 140, 139, 152,
	153, 150, 151, 154, 155, 156, 157, 158, 148, 149,
	142, 143, 141, 144, 146, 147, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 145,
	140, 139, 152, 153, 150, 151, 154, 155, 156, 157,
	158, 148, 149, 142, 143, 141, 144, 146, 147, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 145,
}
var yyPact = [...]int{

	-1000, -1000, 640, -1000, -1000, -1000, 267, -31, -1000, -1000,
	724, 1115, 62, -24, -26, 1006, -1000, -1000, -1000, 935,
	-68, -68, 935, 935, 935, 935, 935, -149, 935, -151,
	935, 935, 935, 935, 935, 935, 935, 935, 935, 935,
	935, 935, 935, -1000, -1000, -1000, -1000, -65, -130, 1206,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 935,
	-149, 935, 935, -151, 935, 935, 935, 935, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -137, -111, -1000, 525,
	-50, 935, -1000, 935, 935, 935, 935, 935, 935, 935,
	935, 935, 935, 935, 935, 935, 935, 935, 935, 935,
	935, 935, 935, 935, 935, 935, 935, 935, 935, 916,
	935, -28, -1000, -29, -1000, -142, 847, 935, 935, 935,
	935, 935, 935, 935, 935, 935, 935, 935, 935, -1000,
	-1000, -1000, -1000, -1000, -7, -7, -7, -7, 794, 935,
	1206, 935, 1206, 1206, -7, -7, -7, -7, -7, -7,
	-7, -7, 1370, 1333, 1370, 935, -1000, -1000, -1000, -30,
	-1000, -1000, -1000, -1000, -1000, -152, 398, 1441, 1250, 1370,
	1293, 1474, 619, 827, 48, 48, 48, -7, -7, -7,
	-7, 208, 208, 466, 466, 466, 466, 466, 237, 237,
	237, 237, 1160, 935, 1407, -153, -143, -146, 1370, 935,
	1370, 1370, 1370, 1370, 1370, 1370, 1370, 1370, 1370, 1370,
	1370, 1370, -1000, 513, 188, 935, 1070, -1000, 935, 935,
	1407, 28, -147, -1000, 619, -1000, -1000, 1370, -1000, 39,
	1407, -156, 100, -1000, 33, -1000, -1000, -38, -1000, -1000,
	-1000, -1000, -31, -112, -1000, -140, 73, 28, -127, -1000,
	-1000, -1000, -31, -1000, -144, 28, -1000, -39, -1000, -1000,
	-1000, -1000, 89, 410, 935, -1000, 1206,
}
var yyPgo = [...]int{

	0, 108, 189, 188, 187, 186, 185, 3, 183, 182,
	180, 114, 166, 1, 165, 121, 120, 164, 133, 2,
	5, 131, 130, 4, 0, 129, 128, 105, 64, 119,
	111, 110,
}
var yyR1 = [...]int{

	0, 29, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 31,
	31, 31, 31, 31, 31, 31, 5, 5, 8, 8,
	7, 9, 9, 9, 10, 10, 6, 6, 6, 6,
	6, 13, 13, 12, 12, 12, 11, 11, 11, 15,
	15, 14, 14, 1, 1, 16, 21, 21, 22, 22,
	23, 23, 17, 17, 4, 4, 2, 2, 3, 3,
	19, 19, 20, 20, 20, 18, 18, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 24, 24, 26, 27, 28, 28, 28,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	1, 1, 2, 3, 2, 0, 1, 1, 3, 3,
	1, 2, 0, 1, 1, 1, 3, 7, 2, 5,
	4, 1, 2, 1, 1, 10, 1, 0, 1, 3,
	4, 6, 0, 1, 0, 1, 0, 1, 0, 1,
	1, 2, 1, 1, 1, 0, 2, 3, 4, 4,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 2, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 2, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	4, 3, 4, 2, 2, 4, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 1, 2, 4,
	2, 1, 1, 1, 1, 1, 4, 2,
}
var yyChk = [...]int{

	-1000, -29, -10, -6, -11, -16, 4, 153, -15, 159,
	95, -24, 112, -14, 127, -27, -25, -1, -26, 71,
	58, 59, 49, 50, 55, 57, 161, 125, 5, 6,
	7, 8, 60, 61, 62, 63, 64, 65, 66, 67,
	13, 14, 16, 76, 77, -28, 87, 163, -5, -24,
	85, -31, -30, 75, 76, 77, 78, 79, 80, 71,
	125, 4, 5, 6, 7, 8, 13, 14, 10, 11,
	12, 56, 70, 81, 82, 72, 73, 74, 94, 95,
	96, 97, 98, 99, 100, 101, 102, 103, 104, 115,
	116, 117, 118, 119, 120, 121, 122, 123, 124, 110,
	111, 112, 113, 114, 133, 105, 106, 107, 108, 109,
	134, 135, 130, 131, 153, 128, 129, 127, 136, 137,
	139, 138, 140, 141, 155, 154, -7, -8, 85, -13,
	-11, 4, 158, 33, 34, 10, 12, 11, 35, 37,
	36, 51, 49, 50, 52, 68, 53, 54, 47, 48,
	40, 41, 38, 39, 42, 43, 44, 45, 46, 30,
	32, -4, 37, 127, -1, 85, 17, 18, 19, 20,
	29, 21, 22, 23, 24, 25, 26, 27, 28, 58,
	59, -24, -27, -27, -24, -24, -24, -24, -24, 161,
	-24, 161, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, 159, -28, 158, 158, 156,
	-12, 160, -11, -16, -15, 96, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -24, 31, -24, 85, 85, 159, -24, 37,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, 162, -24, -24, 15, -24, 85, 161, 31,
	-24, 161, 159, 160, -24, 162, 162, -24, 160, -24,
	-24, -21, -22, -23, -17, -19, -20, 30, -9, 134,
	135, -7, 156, 153, 160, 162, 162, 9, -2, 37,
	-20, -7, 156, 158, -18, 31, -23, -3, 157, -7,
	159, -19, 87, -13, 17, 160, -24,
}
var yyDef = [...]int{

	85, -2, 1, 84, 86, 87, 0, 0, 90, 92,
	0, 0, 114, 0, 0, 201, 202, 101, 204, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 197, 0, 103, 104, 203, 205, 0, 0, 183,
	76, 77, 69, 70, 71, 72, 73, 74, 75, 12,
	39, 2, 3, 4, 5, 6, 45, -2, 7, 8,
	9, 10, 11, 13, 14, 15, 16, 17, 18, 19,
	20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 40,
	41, 42, 43, 44, 47, 48, 49, 50, 51, 52,
	53, 54, 55, 56, 57, 58, 59, 60, 61, 62,
	63, 64, 65, 66, 67, 68, 0, 80, 78, 0,
	0, 0, 98, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 115, 0, 102, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 144,
	146, 130, 145, 147, 165, 166, 167, 168, 0, 0,
	184, 0, 186, 187, 188, 189, 190, 191, 192, 193,
	194, 195, 196, 198, 200, 0, 207, 88, 89, 0,
	91, 96, 93, 94, 95, 0, 148, 149, 150, 151,
	152, 153, 154, 155, 156, 157, 158, 159, 160, 161,
	162, 163, 164, -2, -2, -2, -2, -2, -2, -2,
	-2, -2, 0, 0, 181, 0, 0, 0, 127, 0,
	131, 132, 133, 134, 135, 136, 137, 138, 140, 141,
	142, 143, 178, 0, 0, 0, 0, 79, 0, 0,
	180, -2, 0, 100, 128, 182, 185, 199, 206, 0,
	179, 0, 106, 108, 116, 113, 120, 0, 122, 123,
	124, 81, 0, 0, 99, 0, 125, 112, 118, 117,
	121, 82, 0, 97, 0, 0, 109, 0, 119, 83,
	92, 126, 110, 0, 0, 105, 111,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 55, 3, 3, 163, 54, 37, 3,
	161, 162, 52, 49, 9, 50, 51, 53, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 31, 158,
	43, 17, 45, 30, 67, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 69, 3, 3, 36, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 159, 35, 160, 57,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 10, 11, 12,
	13, 14, 15, 16, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 28, 29, 32, 33, 34, 38,
	39, 40, 41, 42, 44, 46, 47, 48, 56, 58,
	59, 60, 61, 62, 63, 64, 65, 66, 68, 70,
	71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
	81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
	101, 102, 103, 104, 105, 106, 107, 108, 109, 110,
	111, 112, 113, 114, 115, 116, 117, 118, 119, 120,
	121, 122, 123, 124, 125, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
	141, 142, 143, 144, 145, 146, 147, 148, 149, 150,
	151, 152, 153, 154, 155, 156, 157,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:206
		{
			fmt.Println(yyDollar[1].node)
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:225
		{
			yyVAL.node = Node("identifier")
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:226
		{
			yyVAL.node = Node("reserved")
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:230
		{
			yyVAL.node = Node("NamespaceParts").append(Node(yyDollar[1].token))
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:231
		{
			yyVAL.node = yyDollar[1].node.append(Node(yyDollar[3].token))
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:235
		{
			yyVAL.node = yyDollar[1].node
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:239
		{
			yyVAL.node = Node("Name").append(yyDollar[1].node)
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:240
		{
			yyVAL.node = Node("Name").append(yyDollar[2].node).attribute("FullyQualified", "true")
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:241
		{
			yyVAL.node = Node("Name").append(yyDollar[3].node).attribute("Relative", "true")
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:245
		{
			yyVAL.node = yyDollar[1].node.append(yyDollar[2].node)
		}
	case 85:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:246
		{
			yyVAL.node = Node("Statements")
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:250
		{
			yyVAL.node = yyDollar[1].node
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:251
		{
			yyVAL.node = yyDollar[1].node
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:252
		{
			yyVAL.node = yyDollar[2].node /*TODO: identifier stub, refactor it*/
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:253
		{
			yyVAL.node = Node("Namespace").append(yyDollar[2].node)
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:254
		{
			yyVAL.node = yyDollar[1].node
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:258
		{
			yyVAL.node = yyDollar[1].node.append(yyDollar[2].node)
		}
	case 92:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:259
		{
			yyVAL.node = Node("statement_list")
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:263
		{
			yyVAL.node = yyDollar[1].node
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:264
		{
			yyVAL.node = yyDollar[1].node
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:265
		{
			yyVAL.node = yyDollar[1].node
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:268
		{
			yyVAL.node = yyDollar[2].node
		}
	case 97:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:269
		{
			yyVAL.node = Node("Do While").append(yyDollar[2].node).append(yyDollar[5].node)
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:270
		{
			yyVAL.node = yyDollar[1].node
		}
	case 99:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:273
		{
			yyVAL.node = yyDollar[1].node.attribute("name", yyDollar[3].token)
		}
	case 100:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:274
		{
			yyVAL.node = Node("Class").attribute("name", yyDollar[2].token)
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:277
		{
			yyVAL.node = Node("Class").attribute(yyDollar[1].value, "true")
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:278
		{
			yyVAL.node = yyDollar[1].node.attribute(yyDollar[2].value, "true")
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:282
		{
			yyVAL.value = "abstract"
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:283
		{
			yyVAL.value = "final"
		}
	case 105:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line parser.y:288
		{
			yyVAL.node = Node("Function").
				attribute("name", yyDollar[3].token).
				attribute("returns_ref", yyDollar[2].value).
				append(yyDollar[5].node).
				append(yyDollar[7].node).
				append(yyDollar[9].node)
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:299
		{
			yyVAL.node = yyDollar[1].node
		}
	case 107:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:300
		{
			yyVAL.node = Node("Parameter list")
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:303
		{
			yyVAL.node = Node("Parameter list").append(yyDollar[1].node)
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:304
		{
			yyVAL.node = yyDollar[1].node.append(yyDollar[3].node)
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:308
		{
			yyVAL.node = Node("Parameter").
				append(yyDollar[1].node).
				attribute("is_reference", yyDollar[2].value).
				attribute("is_variadic", yyDollar[3].value).
				attribute("var", yyDollar[4].token)
		}
	case 111:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:316
		{
			yyVAL.node = Node("Parameter").
				append(yyDollar[1].node).
				attribute("is_reference", yyDollar[2].value).
				attribute("is_variadic", yyDollar[3].value).
				attribute("var", yyDollar[4].token).
				append(yyDollar[6].node)
		}
	case 112:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:327
		{
			yyVAL.node = Node("No type")
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:328
		{
			yyVAL.node = yyDollar[1].node
		}
	case 114:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:332
		{
			yyVAL.value = "false"
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:333
		{
			yyVAL.value = "true"
		}
	case 116:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:337
		{
			yyVAL.value = "false"
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:338
		{
			yyVAL.value = "true"
		}
	case 118:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:342
		{
			yyVAL.value = "false"
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:343
		{
			yyVAL.value = "true"
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:347
		{
			yyVAL.node = yyDollar[1].node
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:348
		{
			yyVAL.node = yyDollar[2].node
			yyVAL.node.attribute("nullable", "true")
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:352
		{
			yyVAL.node = yyDollar[1].node
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:353
		{
			yyVAL.node = Node("array type")
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:354
		{
			yyVAL.node = Node("callable type")
		}
	case 125:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:358
		{
			yyVAL.node = Node("void")
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:359
		{
			yyVAL.node = yyDollar[2].node
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:363
		{
			yyVAL.node = Node("Assign").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 128:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:364
		{
			yyVAL.node = Node("AssignRef").append(yyDollar[1].node).append(yyDollar[4].node)
		}
	case 129:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:365
		{
			yyVAL.node = Node("AssignRef").append(yyDollar[1].node).append(yyDollar[4].node)
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:366
		{
			yyVAL.node = Node("Clone").append(yyDollar[2].node)
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:367
		{
			yyVAL.node = Node("AssignAdd").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:368
		{
			yyVAL.node = Node("AssignSub").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:369
		{
			yyVAL.node = Node("AssignMul").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:370
		{
			yyVAL.node = Node("AssignPow").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:371
		{
			yyVAL.node = Node("AssignDiv").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:372
		{
			yyVAL.node = Node("AssignConcat").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:373
		{
			yyVAL.node = Node("AssignMod").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:374
		{
			yyVAL.node = Node("AssignAnd").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:375
		{
			yyVAL.node = Node("AssignAnd").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:376
		{
			yyVAL.node = Node("AssignOr").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:377
		{
			yyVAL.node = Node("AssignXor").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:378
		{
			yyVAL.node = Node("AssignShiftLeft").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:379
		{
			yyVAL.node = Node("AssignShiftRight").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:380
		{
			yyVAL.node = Node("PostIncrement").append(yyDollar[1].node)
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:381
		{
			yyVAL.node = Node("PreIncrement").append(yyDollar[2].node)
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:382
		{
			yyVAL.node = Node("PostDecrement").append(yyDollar[1].node)
		}
	case 147:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:383
		{
			yyVAL.node = Node("PreDecrement").append(yyDollar[2].node)
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:384
		{
			yyVAL.node = Node("Or").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:385
		{
			yyVAL.node = Node("And").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:386
		{
			yyVAL.node = Node("Or").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:387
		{
			yyVAL.node = Node("And").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:388
		{
			yyVAL.node = Node("Xor").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:389
		{
			yyVAL.node = Node("BitwiseOr").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:390
		{
			yyVAL.node = Node("BitwiseAnd").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:391
		{
			yyVAL.node = Node("BitwiseXor").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:392
		{
			yyVAL.node = Node("Concat").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:393
		{
			yyVAL.node = Node("Add").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:394
		{
			yyVAL.node = Node("Sub").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:395
		{
			yyVAL.node = Node("Mul").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:396
		{
			yyVAL.node = Node("Pow").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:397
		{
			yyVAL.node = Node("Div").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:398
		{
			yyVAL.node = Node("Mod").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:399
		{
			yyVAL.node = Node("ShiftLeft").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:400
		{
			yyVAL.node = Node("ShiftRight").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 165:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:401
		{
			yyVAL.node = Node("UnaryPlus").append(yyDollar[2].node)
		}
	case 166:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:402
		{
			yyVAL.node = Node("UnaryMinus").append(yyDollar[2].node)
		}
	case 167:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:403
		{
			yyVAL.node = Node("BooleanNot").append(yyDollar[2].node)
		}
	case 168:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:404
		{
			yyVAL.node = Node("BitwiseNot").append(yyDollar[2].node)
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:405
		{
			yyVAL.node = Node("Identical").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:406
		{
			yyVAL.node = Node("NotIdentical").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:407
		{
			yyVAL.node = Node("Equal").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:408
		{
			yyVAL.node = Node("NotEqual").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:409
		{
			yyVAL.node = Node("Spaceship").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:410
		{
			yyVAL.node = Node("Smaller").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:411
		{
			yyVAL.node = Node("SmallerOrEqual").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:412
		{
			yyVAL.node = Node("Greater").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:413
		{
			yyVAL.node = Node("GreaterOrEqual").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:414
		{
			yyVAL.node = yyDollar[2].node
		}
	case 179:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:415
		{
			yyVAL.node = Node("Ternary").append(yyDollar[1].node).append(yyDollar[3].node).append(yyDollar[5].node)
		}
	case 180:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:416
		{
			yyVAL.node = Node("Ternary").append(yyDollar[1].node).append(yyDollar[4].node)
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:417
		{
			yyVAL.node = Node("Coalesce").append(yyDollar[1].node).append(yyDollar[3].node)
		}
	case 182:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:418
		{
			yyVAL.node = Node("Empty").append(yyDollar[3].node)
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:419
		{
			yyVAL.node = Node("Include").append(yyDollar[2].node)
		}
	case 184:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:420
		{
			yyVAL.node = Node("IncludeOnce").append(yyDollar[2].node)
		}
	case 185:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:421
		{
			yyVAL.node = Node("Eval").append(yyDollar[3].node)
		}
	case 186:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:422
		{
			yyVAL.node = Node("Require").append(yyDollar[2].node)
		}
	case 187:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:423
		{
			yyVAL.node = Node("RequireOnce").append(yyDollar[2].node)
		}
	case 188:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:424
		{
			yyVAL.node = Node("CastInt").append(yyDollar[2].node)
		}
	case 189:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:425
		{
			yyVAL.node = Node("CastDouble").append(yyDollar[2].node)
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:426
		{
			yyVAL.node = Node("CastString").append(yyDollar[2].node)
		}
	case 191:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:427
		{
			yyVAL.node = Node("CastArray").append(yyDollar[2].node)
		}
	case 192:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:428
		{
			yyVAL.node = Node("CastObject").append(yyDollar[2].node)
		}
	case 193:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:429
		{
			yyVAL.node = Node("CastBool").append(yyDollar[2].node)
		}
	case 194:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:430
		{
			yyVAL.node = Node("CastUnset").append(yyDollar[2].node)
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:431
		{
			yyVAL.node = Node("Silence").append(yyDollar[2].node)
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:433
		{
			yyVAL.node = Node("Print").append(yyDollar[2].node)
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:434
		{
			yyVAL.node = Node("Yield")
		}
	case 198:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:435
		{
			yyVAL.node = Node("Yield").append(yyDollar[2].node)
		}
	case 199:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:436
		{
			yyVAL.node = Node("Yield").append(yyDollar[2].node).append(yyDollar[4].node)
		}
	case 200:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:437
		{
			yyVAL.node = Node("YieldFrom").append(yyDollar[2].node)
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:441
		{
			yyVAL.node = yyDollar[1].node
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:442
		{
			yyVAL.node = yyDollar[1].node
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:446
		{
			yyVAL.node = yyDollar[1].node
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:450
		{
			yyVAL.node = yyDollar[1].node
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:454
		{
			yyVAL.node = Node("Variable").attribute("name", yyDollar[1].token)
		}
	case 206:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:455
		{
			yyVAL.node = yyDollar[3].node
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:456
		{
			yyVAL.node = Node("Variable").append(yyDollar[2].node)
		}
	}
	goto yystack /* stack new state and value */
}
