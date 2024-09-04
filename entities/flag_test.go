package entities

import (
	"testing"
)

func Test_MakeFeatureFlag_ShouldTransformSlugFromUpperCaseInLowerCase(t *testing.T) {
	value := true
	owners := []string{"owner_test"}
	services := []string{"service_test"}
	var user = User{GivenName: "Romane", FamilyName: "Cavey"}
	flag := MakeFeatureFlag(services, "LAbEL", &value, "service", owners, "description", user)
	if flag.Slug != "label" {
		t.Fatalf("wrong slug, expected : label")
	}

}

func Test__MakeFeatureFlag_ShouldTransformSlugFromAtSymbolIntoAtLetter(t *testing.T) {
	value := true
	owners := []string{"owner_test"}
	services := []string{"service_test"}
	var user = User{GivenName: "Romane", FamilyName: "Cavey"}
	flag := MakeFeatureFlag(services, "@label", &value, "service", owners, "description", user)
	if flag.Slug != "atlabel" {
		t.Fatalf("wrong slug, expected : atlabel")
	}

}

func Test__MakeFeatureFlag_ShouldTransformSlugFromSpaceIntoHyphen(t *testing.T) {
	value := true
	owners := []string{"owner_test"}
	services := []string{"service_test"}
	var user = User{GivenName: "Romane", FamilyName: "Cavey"}
	flag := MakeFeatureFlag(services, "citron is a great compagny", &value, "service", owners, "description", user)
	if flag.Slug != "citron-is-a-great-compagny" {
		t.Fatalf("wrong slug, expected: citron-is-a-great-compagny")
	}

}

func Test__MakeFeatureFlag_ShouldTransformSlugAndSymbolIntoAndLetter(t *testing.T) {
	value := true
	owners := []string{"owner_test"}
	services := []string{"service_test"}
	var user = User{GivenName: "Romane", FamilyName: "Cavey"}
	flag := MakeFeatureFlag(services, "test&compagnie", &value, "service", owners, "description", user)
	if flag.Slug != "testandcompagnie" {
		t.Fatalf("wrong slug, expected: testandcompagnie ")
	}

}

func Test__MakeFeatureFlag_ShouldsKeepLowerCase(t *testing.T) {
	value := true
	owners := []string{"owner_test"}
	services := []string{"service_test"}
	var user = User{GivenName: "Romane", FamilyName: "Cavey"}
	flag := MakeFeatureFlag(services, "label", &value, "service", owners, "description", user)
	if flag.Slug != "label" {
		t.Fatalf("wrong slug, expected: label")
	}

}
