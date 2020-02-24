package main

import (
	"encoding/binary"
	"fmt"
	"strings"
)

func shortName(orig_name string, namespace string) string {

	// ---------------------------------------------------------------
	// try to create a short name for the pod, this is a bit tricky :)
	// ---------------------------------------------------------------
	short_name := ""
	long_name := ""

	// ---------------------------------------------------------------
	// strip off domain name
	// ---------------------------------------------------------------
	// example:
	//    given... kube-proxy-ip-10-7-32-158.ec2.internal
	//    want...  kube-proxy-ip-10-7-32-158
	pos := strings.IndexAny(orig_name, ".")
	if pos == -1 {
		long_name = orig_name
	} else {
		long_name = orig_name[0:pos]
	}

	// ---------------------------------------------------------------
	// create a short name from long_name
	// ---------------------------------------------------------------
	if namespace == "kube-system" {
		// -----------------------------------------------------------
		// kube-system pod names
		// -----------------------------------------------------------
		// examples:
		//   name       = kube-proxy-ip-10-7-32-158
		//   short_name = kube-proxy-ip-10-7-32-158
		//
		//   name       = kube-dns-57dd96bb49-bmcmw
		//   short_name = kube-dns, meaning strip last 17
		//
		pos = strings.Index(long_name, "-ip-")
		if pos == -1 {
			short_name = stripInstanceName(long_name, namespace)
		} else {
			short_name = long_name
		}
	} else {
		// -----------------------------------------------------------
		// other pod names
		// -----------------------------------------------------------
		short_name = stripInstanceName(long_name, namespace)
	}
	return short_name
}

func stripInstanceName(long_name string, namespace string) string {
	var ascii_dash byte
	ascii_dash = 45
	short_name := long_name
	long_name_b := []byte(long_name) // string -> []byte

	// example short names:
	//
	//   name       = cups-7f98769cb6-89mkr
	//   short_name = cups, meaning strip last 17
	//
	//   name       = vstream-daemonset-4cxfx
	//   short_name = vstream-damonset, meaning strip last 6
	//
	len_name := binary.Size(long_name_b)
	dash_pos := len_name - 17
	if long_name_b[dash_pos] == ascii_dash {
		short_name = long_name[0:dash_pos]
	} else {
		dash_pos = len_name - 6
		if long_name_b[dash_pos] == ascii_dash {
			short_name = long_name[0:dash_pos]
		}
	}
	return short_name
}
func main() {
	namespace := "default"
	name := "testing long name"
	short_name := "testing short name"

	name = "kube-dns-57dd96bb49-bmcmw"
	namespace = "default"
	short_name = shortName(name, namespace)
	fmt.Println("Name=", name, " Short name=", short_name)

	name = "cups-db-57dd96bb49-bmcmw"
	namespace = "default"
	short_name = shortName(name, namespace)
	fmt.Println("Name=", name, " Short name=", short_name)

	name = "kube-proxy-ip-10-7-32-15"
	namespace = "kube-system"
	short_name = shortName(name, namespace)
	fmt.Println("Name=", name, " Short name=", short_name)

	name = "kube-proxy-ip-10-7-32-158.ec2.internal"
	namespace = "kube-system"
	short_name = shortName(name, namespace)
	fmt.Println("Name=", name, " Short name=", short_name)

}
