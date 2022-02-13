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

var girlName string
var skillName string

func ( t tina ) FnSetGirlName( name string ) () {
	girlName = name

	fmt.Printf( "%s%s\n", LABEL_GIRLS_NAME, girlName )

	return
}

func ( t tina ) FnSetSkillName( target string ) () {
	skillName = target

	fmt.Printf( "%s%s()\n", LABEL_TARGET, skillName )

	return
}

func ( t tina ) FnJudge( judgement bool ) () {

	if judgement {
		fmt.Printf( "%sOK\n", LABEL_JUDGE )
	}else{
		fmt.Printf( "%sNG - Please check %s\n", LABEL_JUDGE, skillName )
	}

	return
}
