/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : logparser_test.go

* Purpose :

* Creation Date : 06-02-2015

* Last Modified : Wed 24 Jun 2015 01:19:59 AM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package logparser_test

import (
	"fmt"
	"github.com/kiyor/logparser"
	"testing"
)

func TestNewParser(t *testing.T) {
	line := `1433274147.666  1701 1.2.3.4 TCP_MISS/200 442 GET https://b.com/ - DIRECT/1.2.3.4 application/json "https://a.com/home/" "Mozilla/5.0 (Linux; Android 4.4.4; HW-HUAWEI Y635-CL00 Build/HuaweiY635-CL00) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Mobile Safari/537.36" unauth_redirect=%22https://a.com/home/%22;%20logged_in=true;%20session=a;%20utag_main=v_id:$_sn:7$_ss:0$_st:1433182713485$_22%7D;%20_ga=GA1.3.443936681.1432824651 <test1> (test2) {t(es)t3}`

	parser, err := logparser.NewParser('[', ']', '"', '"', '{', '}', '<', '>', '(', ')')
	if err != nil {
		t.Fatal(err.Error())
	}

	for k, v := range parser.Split(line) {
		fmt.Printf("%d|%v|\n", k, v)
	}
}
func ExampleNewParser(t *testing.T) {
	line := `1433274147.666  1701 1.2.3.4 TCP_MISS/200 442 GET https://b.com/ - DIRECT/1.2.3.4 application/json "https://a.com/home/" "Mozilla/5.0 (Linux; Android 4.4.4; HW-HUAWEI Y635-CL00 Build/HuaweiY635-CL00) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Mobile Safari/537.36" unauth_redirect=%22https://a.com/home/%22;%20logged_in=true;%20session=a;%20utag_main=v_id:$_sn:7$_ss:0$_st:1433182713485$_22%7D;%20_ga=GA1.3.443936681.1432824651 <test1> (test2) {t(es)t3}`

	parser, err := logparser.NewParser('[', ']', '"', '"', '{', '}', '<', '>', '(', ')')
	if err != nil {
		panic(err)
	}

	for k, v := range parser.Split(line) {
		fmt.Printf("%d|%v|\n", k, v)
	}

}
