package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestIni_Read(t *testing.T) {
	ini := Ini{}
	src, dst, err := ini.Read("testutil/source.ini", "testutil/destination.ini")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, `map[[Server "server2"]:[Name = "Server dua"]]`, fmt.Sprint(src))
	assert.Equal(t, `[# this is comment # baris kedua [Server "server1"] Name = "Server 1"  [Server "server2"] Name = "Server 2" ]`, fmt.Sprint(dst))
}

func TestIni_Override(t *testing.T) {
	src := Source{
		"[Server2]": []string{
			"URL = 'http://example.com'",
		},
	}
	dest := Destination{
		"[Server1]",
		"Name = 'Server1'",
		"URL = 'http://server1.com'",
		"[Server2]",
		"Name = 'Server2'",
		"URL = 'http://server2.com'",
	}

	ini := Ini{}
	dst, _ := ini.Override(src, dest)
	assert.Equal(t, Destination{
		"[Server1]",
		"Name = 'Server1'",
		"URL = 'http://server1.com'",
		"[Server2]",
		"Name = 'Server2'",
		"URL = 'http://example.com'",
	}, dst)
}

func TestIni_Write(t *testing.T) {
	ini := Ini{}
	dest := Destination{
		"[Server1]",
		"Name = 'Server1'",
		"URL = 'http://server1.com'",
		"[Server2]",
		"Name = 'Server2'",
		"URL = 'http://server2.com'",
	}
	err := ini.Write(dest, "testutil/new.ini")
	log.Println(err)
}
