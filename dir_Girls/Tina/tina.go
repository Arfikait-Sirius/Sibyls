package sibyls

import (
	"fmt"
)

type tina struct{}

var Tina tina

var LABEL_GIRLS_NAME string = "[GIRLS-NAME]: "
var LABEL_TARGET string = "[    TARGET]: "
var LABEL_JUDGE string = "[     JUDGE]: "
var LABEL_PRINT string = "[     PRINT]: "

func ( t tina ) FnSetGirlName( name string ) () {

	fmt.Printf( "%s%s\n", LABEL_GIRLS_NAME, name )

	return
}

func ( t tina ) FnSetSkillName( target string ) () {

	fmt.Printf( "%s%s()\n", LABEL_TARGET, target )

	return
}

func ( t tina ) FnJudge( judgement bool ) () {

	if judgement {
		fmt.Printf( "%sOK\n", LABEL_JUDGE )
	}else{
		fmt.Printf( "%sNG - Please check this source.\n", LABEL_JUDGE )
	}

	return
}
