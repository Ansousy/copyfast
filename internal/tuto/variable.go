package tuto

import "fmt"

//nil  pas de pointeur ni de valeurs
var leNomDeMaVariablePrivateNil string
var leNomDeMaVariablePublicNil string

var leNomDeMaVariableTypePrivateValue string = "VersionType"
var leNomDeMaVariableTypePublicValue string = "VersionType"

var leNomDeMaVariableNonTypePrivateValue string = "VersionNonType"
var leNomDeMaVariableNonTypePublicValue string = "VersionNonType"

const maConstantePriveTypePrivateValue string = "VersionType"
const maConstantePriveTypePublicValue string = "VersionType"

const maConstantePriveNonTypePrivateValue string = "VersionNonType"
const maConstantePriveNonTypePublicValue string = "VersionNonType"

func sampleFunction() {
	//nil  pas de pointeur ni de valeurs
	var leNomDeMaVariablePrivateNil string
	var leNomDeMaVariablePublicNil string
	fmt.Println(leNomDeMaVariablePrivateNil, leNomDeMaVariablePublicNil)

	var leNomDeMaVariableTypePrivateValue string = "VersionType"
	var leNomDeMaVariableTypePublicValue string = "VersionType"
	fmt.Println(leNomDeMaVariableTypePrivateValue, leNomDeMaVariableTypePublicValue)

	var leNomDeMaVariableNonTypePrivateValue string = "VersionNonType"
	var leNomDeMaVariableNonTypePublicValue string = "VersionNonType"
	fmt.Println(leNomDeMaVariableNonTypePrivateValue, leNomDeMaVariableNonTypePublicValue)

	const maConstantePriveTypePrivateValue string = "VersionType"
	const maConstantePriveTypePublicValue string = "VersionType"

	const maConstantePriveNonTypePrivateValue string = "VersionNonType"
	const maConstantePriveNonTypePublicValue string = "VersionNonType"

	variableLiteral := "Literal"
	fmt.Println(variableLiteral)
}
