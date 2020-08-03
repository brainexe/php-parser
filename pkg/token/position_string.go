// Code generated by "stringer -type=Position -output ./position_string.go"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Start-0]
	_ = x[End-1]
	_ = x[SemiColon-2]
	_ = x[AltEnd-3]
	_ = x[Ampersand-4]
	_ = x[Name-5]
	_ = x[Key-6]
	_ = x[Var-7]
	_ = x[ReturnType-8]
	_ = x[CaseSeparator-9]
	_ = x[LexicalVars-10]
	_ = x[Params-11]
	_ = x[Ref-12]
	_ = x[Cast-13]
	_ = x[Expr-14]
	_ = x[InitExpr-15]
	_ = x[CondExpr-16]
	_ = x[IncExpr-17]
	_ = x[True-18]
	_ = x[Cond-19]
	_ = x[HaltCompiller-20]
	_ = x[Namespace-21]
	_ = x[Static-22]
	_ = x[Class-23]
	_ = x[Use-24]
	_ = x[While-25]
	_ = x[For-26]
	_ = x[Switch-27]
	_ = x[Foreach-28]
	_ = x[Declare-29]
	_ = x[Label-30]
	_ = x[Finally-31]
	_ = x[List-32]
	_ = x[Default-33]
	_ = x[If-34]
	_ = x[ElseIf-35]
	_ = x[Else-36]
	_ = x[Function-37]
	_ = x[Alias-38]
	_ = x[Equal-39]
	_ = x[Exit-40]
	_ = x[Array-41]
	_ = x[Isset-42]
	_ = x[Empty-43]
	_ = x[Eval-44]
	_ = x[Echo-45]
	_ = x[Try-46]
	_ = x[Catch-47]
	_ = x[Unset-48]
	_ = x[Stmts-49]
	_ = x[VarList-50]
	_ = x[ConstList-51]
	_ = x[NameList-52]
	_ = x[ParamList-53]
	_ = x[ModifierList-54]
	_ = x[ArrayPairList-55]
	_ = x[CaseListStart-56]
	_ = x[CaseListEnd-57]
	_ = x[ArgumentList-58]
	_ = x[PropertyList-59]
	_ = x[ParameterList-60]
	_ = x[AdaptationList-61]
	_ = x[LexicalVarList-62]
	_ = x[OpenParenthesisToken-63]
	_ = x[CloseParenthesisToken-64]
}

const _Position_name = "StartEndSemiColonAltEndAmpersandNameKeyVarReturnTypeCaseSeparatorLexicalVarsParamsRefCastExprInitExprCondExprIncExprTrueCondHaltCompillerNamespaceStaticClassUseWhileForSwitchForeachDeclareLabelFinallyListDefaultIfElseIfElseFunctionAliasEqualExitArrayIssetEmptyEvalEchoTryCatchUnsetStmtsVarListConstListNameListParamListModifierListArrayPairListCaseListStartCaseListEndArgumentListPropertyListParameterListAdaptationListLexicalVarListOpenParenthesisTokenCloseParenthesisToken"

var _Position_index = [...]uint16{0, 5, 8, 17, 23, 32, 36, 39, 42, 52, 65, 76, 82, 85, 89, 93, 101, 109, 116, 120, 124, 137, 146, 152, 157, 160, 165, 168, 174, 181, 188, 193, 200, 204, 211, 213, 219, 223, 231, 236, 241, 245, 250, 255, 260, 264, 268, 271, 276, 281, 286, 293, 302, 310, 319, 331, 344, 357, 368, 380, 392, 405, 419, 433, 453, 474}

func (i Position) String() string {
	if i < 0 || i >= Position(len(_Position_index)-1) {
		return "Position(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Position_name[_Position_index[i]:_Position_index[i+1]]
}