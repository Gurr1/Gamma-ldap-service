package main

import (
	"fmt"
	"gopkg.in/ldap.v3"
	"log"
)

var cidList [] string

func GetITUsers() (Users, error) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "ldap.chalmers.se", 389))

	if err != nil {
		log.Fatal(err)
		return Users{}, err
	}
	defer l.Close()
	searchRequest := ldap.NewSearchRequest(
		"ou=groups,dc=chalmers,dc=se",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(cn=%s))", "s_passer_prog_tkite"),
		[]string{"memberUid"},
		nil,
	)
	result, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
		return Users{}, nil
	}
	for _, entry := range result.Entries {
		users := entry.GetAttributeValues("memberUid")
		for _, user := range users {
			cidList = append(cidList, user)
		}
	}
	return Users{Cids: cidList}, nil
}