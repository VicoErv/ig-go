package v1

import (
	"fmt"
	"math"
	"strconv"
)

// User as user
type User struct {
	username  string
	useragent string
	session   string
}

// def useragent_hash
// agent = [api + '/' + release, dpi + 'dpi',
// 		 resolution, info[:manufacturer],
// 		 info[:model], info[:device], @language]

// {
//   agent: agent.join('; '),
//   version: Constants::PRIVATE_KEY[:APP_VERSION]
// }
// end

// def useragent
// format('Instagram %s Android(%s)',
// useragent_hash[:version],
// useragent_hash[:agent].rstrip)
// end

// def md5int
// (md5.to_i(32) / 10e32).round
// end

var release = []string{
	"4.0.4", "4.3.1", "4.4.4", "5.1.1", "6.0.1",
}

var dpi = []string{
	"801", "577", "576", "538", "515", "424", "401", "373",
}

var resolution = []string{
	"3840x2160", "1440x2560", "2560x1440", "1440x2560",
	"2560x1440", "1080x1920", "1080x1920", "1080x1920",
}

func (u *User) userAgent() string {
	var i, _ = u.api()

	var release = release[i%5]
	var dpi = dpi[i%8]
	var resolution = resolution[i%8]

	fmt.Println(release, dpi, resolution)
	return ""
}

func (u *User) api() (int64, string) {
	var hash = CreateMD5(u.username)
	var n, _ = strconv.ParseInt(hash, 10, 32)

	var f = int64(math.Round(float64(n) / 10e32))

	return f, string(18 + (f % 5))
}

// Login create user instance
func Login(username string, password string) *User {
	user := &User{}

	user.username = username

	return user
}

// Using user session
func Using(session string) *User {
	user := &User{}

	user.session = session

	return user
}
